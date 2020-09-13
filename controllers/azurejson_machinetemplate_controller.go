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
	"sigs.k8s.io/cluster-api/util/annotations"
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

// AzureJSONTemplateReconciler reconciles azure json secrets for AzureMachineTemplate objects
type AzureJSONTemplateReconciler struct {
	client.Client
	Log              logr.Logger
	Recorder         record.EventRecorder
	ReconcileTimeout time.Duration
}

// SetupWithManager initializes this controller with a manager
func (r *AzureJSONTemplateReconciler) SetupWithManager(mgr ctrl.Manager, options controller.Options) error {
	return ctrl.NewControllerManagedBy(mgr).
		WithOptions(options).
		For(&infrav1.AzureMachineTemplate{}).
		Owns(&corev1.Secret{}).
		Complete(r)
}

// Reconcile reconciles azure json secrets for azure machine templates
func (r *AzureJSONTemplateReconciler) Reconcile(req ctrl.Request) (_ ctrl.Result, reterr error) {
	ctx, cancel := context.WithTimeout(context.Background(), reconciler.DefaultedLoopTimeout(r.ReconcileTimeout))
	defer cancel()
	log := r.Log.WithValues("namespace", req.Namespace, "AzureMachineTemplate", req.Name)

	// Fetch the AzureMachineTemplate instance
	azureMachineTemplate := &infrav1.AzureMachineTemplate{}
	err := r.Get(ctx, req.NamespacedName, azureMachineTemplate)
	if err != nil {
		if apierrors.IsNotFound(err) {
			log.Info("object was not found")
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}

	// Fetch the Cluster.
	cluster, err := util.GetOwnerCluster(ctx, r.Client, azureMachineTemplate.ObjectMeta)
	if err != nil {
		return reconcile.Result{}, err
	}
	if cluster == nil {
		log.Info("Cluster Controller has not yet set OwnerRef")
		return reconcile.Result{}, nil
	}

	log = log.WithValues("cluster", cluster.Name)

	// Return early if the object or Cluster is paused.
	if annotations.IsPaused(cluster, azureMachineTemplate) {
		log.Info("AzureMachineTemplate or linked Cluster is marked as paused. Won't reconcile")
		return ctrl.Result{}, nil
	}

	// TODO(ace): refactor the next ~50-100 lines
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
		r.Recorder.Eventf(azureMachineTemplate, corev1.EventTypeNormal, msg, msg)
		log.Info(msg)
		return reconcile.Result{}, nil
	}
	// TODO(ace): end refactor

	log = log.WithValues(clusterDescriberRef.Kind, clusterDescriber.GetName())

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

	apiVersion, kind := infrav1.GroupVersion.WithKind("AzureMachineTemplate").ToAPIVersionAndKind()
	owner := metav1.OwnerReference{
		APIVersion: apiVersion,
		Kind:       kind,
		Name:       azureMachineTemplate.GetName(),
		UID:        azureMachineTemplate.GetUID(),
	}

	// Construct secret for this machine template
	userAssignedIdentityIfExists := ""
	if len(azureMachineTemplate.Spec.Template.Spec.UserAssignedIdentities) > 0 {
		userAssignedIdentityIfExists = azureMachineTemplate.Spec.Template.Spec.UserAssignedIdentities[0].ProviderID
	}

	newSecret, err := GetCloudProviderSecret(
		clusterScope,
		azureMachineTemplate.Namespace,
		azureMachineTemplate.Name,
		owner,
		azureMachineTemplate.Spec.Template.Spec.Identity,
		userAssignedIdentityIfExists,
	)

	if err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "failed to create cloud provider config")
	}

	if err := reconcileAzureSecret(ctx, log, r.Client, owner, newSecret, clusterScope.ClusterName()); err != nil {
		r.Recorder.Eventf(azureMachineTemplate, corev1.EventTypeWarning, "Error reconciling cloud provider secret for AzureMachineTemplate", err.Error())
		return ctrl.Result{}, errors.Wrap(err, "failed to reconcile azure secret")
	}

	return ctrl.Result{}, nil
}
