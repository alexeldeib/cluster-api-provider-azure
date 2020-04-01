/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"os"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/record"
	infrav1 "sigs.k8s.io/cluster-api-provider-azure/api/v1alpha3"
	expv1 "sigs.k8s.io/cluster-api/exp/api/v1alpha3"

	"sigs.k8s.io/cluster-api/util"
	"sigs.k8s.io/cluster-api/util/patch"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// AzureManagedMachinePoolReconciler reconciles a AzureManagedMachinePool object
type AzureManagedMachinePoolReconciler struct {
	client.Client
	Log      logr.Logger
	Recorder record.EventRecorder
}

func (r *AzureManagedMachinePoolReconciler) SetupWithManager(mgr ctrl.Manager, options controller.Options) error {
	return ctrl.NewControllerManagedBy(mgr).
		WithOptions(options).
		For(&infrav1.AzureManagedMachinePool{}).
		Complete(r)
}

// +kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=azuremanagedmachinepools,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=azuremanagedmachinepools/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=cluster.x-k8s.io,resources=clusters;clusters/status,verbs=get;list;watch;patch
// +kubebuilder:rbac:groups=exp.cluster.x-k8s.io,resources=machinepools;machinepools/status,verbs=get;list;watch

func (r *AzureManagedMachinePoolReconciler) Reconcile(req ctrl.Request) (_ ctrl.Result, reterr error) {
	ctx := context.TODO()
	log := r.Log.WithValues("namespace", req.Namespace, "infraPool", req.Name)

	// Fetch the AzureManagedMachinePool instance
	infraPool := &infrav1.AzureManagedMachinePool{}
	err := r.Get(ctx, req.NamespacedName, infraPool)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}

	// Fetch the owning MachinePool.
	ownerPool, err := getOwnerMachinePool(ctx, r.Client, infraPool.ObjectMeta)
	if err != nil {
		return reconcile.Result{}, err
	}
	if ownerPool == nil {
		log.Info("MachinePool Controller has not yet set OwnerRef")
		return reconcile.Result{}, nil
	}

	// Fetch the Cluster.
	ownerCluster, err := util.GetOwnerCluster(ctx, r.Client, ownerPool.ObjectMeta)
	if err != nil {
		return reconcile.Result{}, err
	}
	if ownerCluster == nil {
		log.Info("Cluster Controller has not yet set OwnerRef")
		return reconcile.Result{}, nil
	}

	log = log.WithValues("ownerCluster", ownerCluster.Name)

	// Fetch the managed cluster
	infraCluster := &infrav1.AzureManagedCluster{}
	infraClusterName := client.ObjectKey{
		Namespace: infraPool.Namespace,
		Name:      ownerCluster.Spec.InfrastructureRef.Name,
	}
	if err := r.Client.Get(ctx, infraClusterName, infraCluster); err != nil {
		return reconcile.Result{}, err
	}

	// initialize patch helper
	patchhelper, err := patch.NewHelper(infraPool, r.Client)
	if err != nil {
		return reconcile.Result{}, errors.Wrap(err, "failed to init patchhelper")
	}

	// Always close the scope when exiting this function so we can persist any changes.
	defer func() {
		if err := patchhelper.Patch(ctx, infraPool); err != nil && reterr == nil {
			reterr = err
		}
	}()

	// extract subscription ID from environment
	// TODO(ace): don't do this here
	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	if subscriptionID == "" {
		return reconcile.Result{}, errors.New("error creating azure services. Environment variable AZURE_SUBSCRIPTION_ID is not set")
	}

	// fetch azure authorizer
	// TODO(ace): probably use a secret ref/object ref instead?
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		return reconcile.Result{}, errors.Wrap(err, "failed to get authorizer from environment")
	}

	// ...
	scope := &ManagedClusterContext{
		log:          log,
		patchhelper:  patchhelper,
		aksCluster:   infraCluster,
		infraPool:    infraPool,
		ownerCluster: ownerCluster,
		ownerPool:    ownerPool,
	}

	// Handle deleted clusters
	if !infraPool.DeletionTimestamp.IsZero() {
		return r.reconcileDelete(ctx, scope, subscriptionID, authorizer)
	}

	// Handle non-deleted clusters
	return r.reconcileNormal(ctx, scope, subscriptionID, authorizer)
}

func (r *AzureManagedMachinePoolReconciler) reconcileNormal(ctx context.Context, scope *ManagedClusterContext, subscriptionID string, authorizer autorest.Authorizer) (reconcile.Result, error) {
	scope.log.Info("Reconciling AzureManagedMachinePool")

	// If the AzureManagedMachinePool doesn't have our finalizer, add it.
	controllerutil.AddFinalizer(scope.infraPool, infrav1.ClusterFinalizer)
	// Register the finalizer immediately to avoid orphaning Azure resources on delete
	if err := scope.patchhelper.Patch(ctx, scope.infraPool); err != nil {
		return reconcile.Result{}, err
	}

	if err := newAzureManagedMachinePoolReconciler(subscriptionID, authorizer, r.Client).Reconcile(ctx, scope); err != nil {
		return reconcile.Result{}, errors.Wrapf(err, "error creating AzureManagedMachinePool %s/%s", scope.infraPool.Namespace, scope.infraPool.Name)
	}

	// No errors, so mark us ready so the Cluster API Cluster Controller can pull it
	scope.infraPool.Status.Ready = true

	if err := scope.patchhelper.Patch(ctx, scope.infraPool); err != nil {
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}

func (r *AzureManagedMachinePoolReconciler) reconcileDelete(ctx context.Context, scope *ManagedClusterContext, subscriptionID string, authorizer autorest.Authorizer) (reconcile.Result, error) {
	scope.log.Info("Reconciling AzureManagedMachinePool delete")

	if err := newAzureManagedMachinePoolReconciler(subscriptionID, authorizer, r.Client).Delete(ctx, scope); err != nil {
		return reconcile.Result{}, errors.Wrapf(err, "error deleting AzureManagedMachinePool %s/%s", scope.infraPool.Namespace, scope.infraPool.Name)
	}

	// Cluster is deleted so remove the finalizer.
	controllerutil.RemoveFinalizer(scope.infraPool, infrav1.ClusterFinalizer)

	return reconcile.Result{}, nil
}

// getOwnerMachinePool returns the MachinePool object owning the current resource.
func getOwnerMachinePool(ctx context.Context, c client.Client, obj metav1.ObjectMeta) (*expv1.MachinePool, error) {
	for _, ref := range obj.OwnerReferences {
		if ref.Kind == "MachinePool" && ref.APIVersion == expv1.GroupVersion.String() {
			return getMachinePoolByName(ctx, c, obj.Namespace, ref.Name)
		}
	}
	return nil, nil
}

// getMachinePoolByName finds and return a MachinePool object using the specified params.
func getMachinePoolByName(ctx context.Context, c client.Client, namespace, name string) (*expv1.MachinePool, error) {
	m := &expv1.MachinePool{}
	key := client.ObjectKey{Name: name, Namespace: namespace}
	if err := c.Get(ctx, key, m); err != nil {
		return nil, err
	}
	return m, nil
}
