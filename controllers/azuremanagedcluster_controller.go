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
	"k8s.io/client-go/tools/record"
	infrav1 "sigs.k8s.io/cluster-api-provider-azure/api/v1alpha3"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha3"
	expv1 "sigs.k8s.io/cluster-api/exp/api/v1alpha3"

	"sigs.k8s.io/cluster-api/util"
	"sigs.k8s.io/cluster-api/util/patch"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type ManagedClusterContext struct {
	log          logr.Logger
	patchhelper  *patch.Helper
	aksCluster   *infrav1.AzureManagedCluster
	ownerCluster *clusterv1.Cluster
	infraPool    *infrav1.AzureManagedMachinePool
	ownerPool    *expv1.MachinePool
}

// AzureManagedClusterReconciler reconciles a AzureManagedCluster object
type AzureManagedClusterReconciler struct {
	client.Client
	Log      logr.Logger
	Recorder record.EventRecorder
}

func (r *AzureManagedClusterReconciler) SetupWithManager(mgr ctrl.Manager, options controller.Options) error {
	return ctrl.NewControllerManagedBy(mgr).
		WithOptions(options).
		For(&infrav1.AzureManagedCluster{}).
		Complete(r)
}

// +kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=azuremanagedclusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=azuremanagedclusters/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=cluster.x-k8s.io,resources=clusters;clusters/status,verbs=get;list;watch

func (r *AzureManagedClusterReconciler) Reconcile(req ctrl.Request) (_ ctrl.Result, reterr error) {
	ctx := context.TODO()
	log := r.Log.WithValues("namespace", req.Namespace, "aksCluster", req.Name)

	// Fetch the AzureManagedCluster instance
	aksCluster := &infrav1.AzureManagedCluster{}
	err := r.Get(ctx, req.NamespacedName, aksCluster)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}

	// Fetch the Cluster.
	cluster, err := util.GetOwnerCluster(ctx, r.Client, aksCluster.ObjectMeta)
	if err != nil {
		return reconcile.Result{}, err
	}
	if cluster == nil {
		log.Info("Cluster Controller has not yet set OwnerRef")
		return reconcile.Result{}, nil
	}

	log = log.WithValues("cluster", cluster.Name)

	// initialize patch helper
	patchhelper, err := patch.NewHelper(aksCluster, r.Client)
	if err != nil {
		return reconcile.Result{}, errors.Wrap(err, "failed to init patchhelper")
	}

	// Always close the scope when exiting this function so we can persist any changes.
	defer func() {
		if err := patchhelper.Patch(ctx, aksCluster); err != nil && reterr == nil {
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

	// fetch default pool
	defaultPoolKey := client.ObjectKey{
		Name:      aksCluster.Spec.DefaultPoolRef.Name,
		Namespace: aksCluster.ObjectMeta.Namespace,
	}
	defaultPool := &infrav1.AzureManagedMachinePool{}
	if err := r.Client.Get(ctx, defaultPoolKey, defaultPool); err != nil {
		return reconcile.Result{}, errors.Wrapf(err, "failed to fetch default pool reference")
	}

	// fetch owner of default pool
	ownerPool := &expv1.MachinePool{}
	for _, ref := range defaultPool.OwnerReferences {
		if ref.Kind == "MachinePool" && ref.APIVersion == clusterv1.GroupVersion.String() {
			key := client.ObjectKey{
				Namespace: ownerPool.Namespace,
				Name:      ref.Name,
			}
			if err := r.Client.Get(ctx, key, ownerPool); err != nil {
				return reconcile.Result{}, errors.Wrapf(err, "failed to fetch owner pool for default pool in managed cluster")
			}
		}
	}

	// ...
	scope := &ManagedClusterContext{
		log:          log,
		patchhelper:  patchhelper,
		aksCluster:   aksCluster,
		ownerCluster: cluster,
		infraPool:    defaultPool,
		ownerPool:    ownerPool,
	}

	// Handle deleted clusters
	if !aksCluster.DeletionTimestamp.IsZero() {
		return r.reconcileDelete(ctx, scope, subscriptionID, authorizer)
	}

	// Handle non-deleted clusters
	return r.reconcileNormal(ctx, scope, subscriptionID, authorizer)
}

func (r *AzureManagedClusterReconciler) reconcileNormal(ctx context.Context, scope *ManagedClusterContext, subscriptionID string, authorizer autorest.Authorizer) (reconcile.Result, error) {
	scope.log.Info("Reconciling AzureManagedCluster")

	// If the AzureManagedCluster doesn't have our finalizer, add it.
	controllerutil.AddFinalizer(scope.aksCluster, infrav1.ClusterFinalizer)
	// Register the finalizer immediately to avoid orphaning Azure resources on delete
	if err := scope.patchhelper.Patch(ctx, scope.aksCluster); err != nil {
		return reconcile.Result{}, err
	}

	if err := newAzureManagedClusterReconciler(subscriptionID, authorizer, r.Client).Reconcile(ctx, scope); err != nil {
		return reconcile.Result{}, errors.Wrapf(err, "error creating AzureManagedCluster %s/%s", scope.aksCluster.Namespace, scope.aksCluster.Name)
	}

	// No errors, so mark us ready so the Cluster API Cluster Controller can pull it
	scope.aksCluster.Status.Ready = true
	scope.ownerCluster.Status.ControlPlaneInitialized = true
	scope.ownerCluster.Status.ControlPlaneReady = true

	if err := scope.patchhelper.Patch(ctx, scope.aksCluster); err != nil {
		return reconcile.Result{}, err
	}

	if err := scope.patchhelper.Patch(ctx, scope.ownerCluster); err != nil {
		return reconcile.Result{}, err
	}

	scope.log.Info("Successfully reconciled")

	return reconcile.Result{}, nil
}

func (r *AzureManagedClusterReconciler) reconcileDelete(ctx context.Context, scope *ManagedClusterContext, subscriptionID string, authorizer autorest.Authorizer) (reconcile.Result, error) {
	scope.log.Info("Reconciling AzureManagedCluster delete")

	if err := newAzureManagedClusterReconciler(subscriptionID, authorizer, r.Client).Delete(ctx, scope); err != nil {
		return reconcile.Result{}, errors.Wrapf(err, "error deleting AzureManagedCluster %s/%s", scope.aksCluster.Namespace, scope.aksCluster.Name)
	}

	// Cluster is deleted so remove the finalizer.
	controllerutil.RemoveFinalizer(scope.aksCluster, infrav1.ClusterFinalizer)

	return reconcile.Result{}, nil
}
