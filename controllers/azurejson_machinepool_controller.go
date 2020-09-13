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
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/cluster-api/util"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	infrav1 "sigs.k8s.io/cluster-api-provider-azure/api/v1alpha3"
	azure "sigs.k8s.io/cluster-api-provider-azure/cloud"
	"sigs.k8s.io/cluster-api-provider-azure/cloud/scope"
	expv1 "sigs.k8s.io/cluster-api-provider-azure/exp/api/v1alpha3"
	"sigs.k8s.io/cluster-api-provider-azure/util/reconciler"
)

// AzureJSONMachinePoolReconciler reconciles azure json secrets for AzureMachinePool objects
type AzureJSONMachinePoolReconciler struct {
	client.Client
	Log              logr.Logger
	Recorder         record.EventRecorder
	ReconcileTimeout time.Duration
}

// SetupWithManager initializes this controller with a manager
func (r *AzureJSONMachinePoolReconciler) SetupWithManager(mgr ctrl.Manager, options controller.Options) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&expv1.AzureMachinePool{}).
		Owns(&corev1.Secret{}).
		Complete(r)
}

// Reconcile reconciles the azure json for AzureMachinePool objects
func (r *AzureJSONMachinePoolReconciler) Reconcile(req ctrl.Request) (_ ctrl.Result, reterr error) {
	ctx, cancel := context.WithTimeout(context.Background(), reconciler.DefaultedLoopTimeout(r.ReconcileTimeout))
	defer cancel()
	log := r.Log.WithValues("namespace", req.Namespace, "AzureMachinePool", req.Name)

	// Fetch the AzureMachine instance
	azureMachinePool := &expv1.AzureMachinePool{}
	err := r.Get(ctx, req.NamespacedName, azureMachinePool)
	if err != nil {
		if apierrors.IsNotFound(err) {
			log.Info("object was not found")
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}

	// Fetch the CAPI MachinePool.
	machinePool, err := GetOwnerMachinePool(ctx, r.Client, azureMachinePool.ObjectMeta)
	if err != nil {
		return reconcile.Result{}, err
	}
	if machinePool == nil {
		log.Info("MachinePool Controller has not yet set OwnerRef")
		return reconcile.Result{}, nil
	}

	log = log.WithValues("machinePool", machinePool.Name)

	// Fetch the Cluster.
	cluster, err := util.GetClusterFromMetadata(ctx, r.Client, machinePool.ObjectMeta)
	if err != nil {
		log.Info("MachinePool is missing cluster label or cluster does not exist")
		return reconcile.Result{}, nil
	}

	log = log.WithValues("cluster", cluster.Name)

	// fetch the corresponding azure cluster
	controlPlaneRefKind := cluster.Spec.ControlPlaneRef.Kind
	controlPlaneRefGV, err := schema.ParseGroupVersion(cluster.Spec.ControlPlaneRef.APIVersion)
	if err != nil {
		return ctrl.Result{}, errors.Wrap(err, "failed to parse gv from control plane apiversion")
	}

	infraRefKind := cluster.Spec.InfrastructureRef.Kind
	infraRefGV, err := schema.ParseGroupVersion(cluster.Spec.InfrastructureRef.APIVersion)
	if err != nil {
		return ctrl.Result{}, errors.Wrap(err, "failed to parse gv from infra ref apiversion")
	}

	// TODO(ace): refactor the next ~50-100 lines
	// AzureManagedControlPlane only works with AzureManagedCluster
	var isManaged bool
	var clusterDescriberRef *corev1.ObjectReference

	if controlPlaneRefGV.Group == expv1.GroupVersion.Group && controlPlaneRefKind == "AzureManagedControlPlane" {
		isManaged = true
		clusterDescriberRef = cluster.Spec.ControlPlaneRef
		if infraRefGV.Group != expv1.GroupVersion.Group || infraRefKind != "AzureManagedCluster" {
			log.Info(fmt.Sprintf(
				"AzureManagedControlPlane only works with AzureManagedCluster, but infraRef on Cluster is of Group: '%s', Kind: '%s'",
				infraRefGV.Group,
				infraRefKind,
			))
			// do not attempt requeue, this is a terminal error until user action
			return ctrl.Result{}, nil
		}
	} else if infraRefGV.Group == infrav1.GroupVersion.Group && infraRefKind == "AzureCluster" {
		clusterDescriberRef = cluster.Spec.InfrastructureRef
	} else {
		log.Info(fmt.Sprintf(
			"unable to generate azure cloud provider configuration for infraRef of Group: '%s', Kind: '%s', and controlPlaneRef of Group: '%s', Kind: '%s'",
			infraRefGV.Group,
			infraRefKind,
			controlPlaneRefGV.Group,
			controlPlaneRefKind,
		))
		// do not attempt requeue, this is a terminal error until user action
		return ctrl.Result{}, nil
	}

	var clusterDescriber azure.ClusterDescriber
	if isManaged {
		clusterDescriber = new(expv1.AzureManagedControlPlane)
	} else {
		clusterDescriber = new(infrav1.AzureCluster)
	}

	clusterDescriberName := client.ObjectKey{
		Namespace: clusterDescriberRef.Namespace,
		Name:      clusterDescriberRef.Name,
	}
	if err := r.Client.Get(ctx, clusterDescriberName, clusterDescriber); err != nil {
		msg := fmt.Sprintf("%s unavailable", clusterDescriberRef.Kind)
		r.Recorder.Eventf(azureMachinePool, corev1.EventTypeNormal, msg, msg)
		log.Info(msg)
		return reconcile.Result{}, nil
	}
	// TODO(ace): end refactor

	// Create the scope.
	clusterScope, err := scope.NewClusterScope(scope.ClusterScopeParams{
		Client:           r.Client,
		Logger:           log,
		Cluster:          cluster,
		ClusterDescriber: clusterDescriber,
	})
	if err != nil {
		return reconcile.Result{}, errors.Errorf("failed to create scope: %+v", err)
	}

	apiVersion, kind := infrav1.GroupVersion.WithKind("AzureMachinePool").ToAPIVersionAndKind()
	owner := metav1.OwnerReference{
		APIVersion: apiVersion,
		Kind:       kind,
		Name:       azureMachinePool.GetName(),
		UID:        azureMachinePool.GetUID(),
	}

	newSecret, err := GetCloudProviderSecret(
		clusterScope,
		azureMachinePool.Namespace,
		azureMachinePool.Name,
		owner,
		infrav1.VMIdentityNone,
		"",
	)

	if err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "failed to create cloud provider config")
	}

	if err := reconcileAzureSecret(ctx, log, r.Client, owner, newSecret, clusterScope.ClusterName()); err != nil {
		r.Recorder.Eventf(azureMachinePool, corev1.EventTypeWarning, "Error reconciling cloud provider secret for AzureMachinePool", err.Error())
		return ctrl.Result{}, errors.Wrap(err, "failed to reconcile azure secret")
	}

	return ctrl.Result{}, nil
}
