/*
Copyright 2019 The Kubernetes Authors.

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

	"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2020-02-01/containerservice"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	azure "sigs.k8s.io/cluster-api-provider-azure/cloud"
	"sigs.k8s.io/cluster-api-provider-azure/cloud/scope"
	"sigs.k8s.io/cluster-api-provider-azure/cloud/services/managedclusters"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha3"
	"sigs.k8s.io/cluster-api/util/secret"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// azureManagedControlPlaneReconciler are list of services required by cluster controller
type azureManagedControlPlaneReconciler struct {
	kubeclient         client.Client
	managedClustersSvc azure.CredentialGetter
}

// newAzureManagedControlPlaneReconciler populates all the services based on input scope
func newAzureManagedControlPlaneReconciler(scope *scope.ManagedControlPlaneScope) *azureManagedControlPlaneReconciler {
	return &azureManagedControlPlaneReconciler{
		kubeclient:         scope.Client,
		managedClustersSvc: managedclusters.NewService(scope.AzureClients.Authorizer, scope.AzureClients.SubscriptionID),
	}
}

// Reconcile reconciles all the services in pre determined order
func (r *azureManagedControlPlaneReconciler) Reconcile(ctx context.Context, scope *scope.ManagedControlPlaneScope) error {
	scope.Logger.Info("reconciling cluster")
	managedClusterSpec := &managedclusters.Spec{
		Name:            scope.ControlPlane.Name,
		ResourceGroup:   scope.ControlPlane.Spec.ResourceGroup,
		Location:        scope.ControlPlane.Spec.Location,
		Tags:            scope.ControlPlane.Spec.AdditionalTags,
		Version:         scope.ControlPlane.Spec.Version,
		LoadBalancerSKU: scope.ControlPlane.Spec.LoadBalancerSKU,
		NetworkPlugin:   scope.ControlPlane.Spec.NetworkPlugin,
		NetworkPolicy:   scope.ControlPlane.Spec.NetworkPolicy,
		SSHPublicKey:    scope.ControlPlane.Spec.SSHPublicKey,
	}

	if net := scope.Cluster.Spec.ClusterNetwork; net != nil {
		if net.Services != nil {
			// A user may provide zero or one CIDR blocks. If they provide an empty array,
			// we ignore it and use the default. AKS doesn't support > 1 Service/Pod CIDR.
			if len(net.Services.CIDRBlocks) > 1 {
				return errors.New("managed control planes only allow one service cidr")
			}
			if len(net.Services.CIDRBlocks) == 1 {
				managedClusterSpec.ServiceCIDR = net.Services.CIDRBlocks[0]
			}
		}
		if net.Pods != nil {
			// A user may provide zero or one CIDR blocks. If they provide an empty array,
			// we ignore it and use the default. AKS doesn't support > 1 Service/Pod CIDR.
			if len(net.Pods.CIDRBlocks) > 1 {
				return errors.New("managed control planes only allow one service cidr")
			}
			if len(net.Pods.CIDRBlocks) == 1 {
				managedClusterSpec.PodCIDR = net.Pods.CIDRBlocks[0]
			}
		}
	}

	_, err := r.managedClustersSvc.Get(ctx, managedClusterSpec)
	if err != nil && !azure.ResourceNotFound(err) {
		return errors.Wrapf(err, "failed to reconcile managed cluster %s", scope.ControlPlane.Name)
	}

	if azure.ResourceNotFound(err) {
		// We are creating this cluster for the first time.
		// Configure the default pool, rest will be handled by machinepool controller
		defaultPoolSpec := managedclusters.PoolSpec{
			Name:         scope.InfraMachinePool.Name,
			SKU:          scope.InfraMachinePool.Spec.SKU,
			Replicas:     1,
			OSDiskSizeGB: 0,
		}

		// Set optional values
		if scope.InfraMachinePool.Spec.OSDiskSizeGB != nil {
			defaultPoolSpec.OSDiskSizeGB = *scope.InfraMachinePool.Spec.OSDiskSizeGB
		}
		if scope.MachinePool.Spec.Replicas != nil {
			defaultPoolSpec.Replicas = *scope.MachinePool.Spec.Replicas
		}

		// Add to cluster spec
		managedClusterSpec.AgentPools = []managedclusters.PoolSpec{defaultPoolSpec}
	}

	// Send to Azure for updates
	if err := r.managedClustersSvc.Reconcile(ctx, managedClusterSpec); err != nil {
		return errors.Wrapf(err, "failed to reconcile managed cluster %s", scope.ControlPlane.Name)
	}

	// Fetch newly updated cluster
	managedClusterResult, err := r.managedClustersSvc.Get(ctx, managedClusterSpec)
	if err != nil {
		return err
	}

	managedCluster, ok := managedClusterResult.(containerservice.ManagedCluster)
	if !ok {
		return fmt.Errorf("expected containerservice ManagedCluster object")
	}

	old := scope.ControlPlane.DeepCopyObject()

	scope.ControlPlane.Spec.ControlPlaneEndpoint = clusterv1.APIEndpoint{
		Host: *managedCluster.ManagedClusterProperties.Fqdn,
		Port: 443,
	}

	if err := r.kubeclient.Patch(ctx, scope.ControlPlane, client.MergeFrom(old)); err != nil {
		return errors.Wrapf(err, "failed to set control plane endpoint")
	}

	// Always fetch credentials in case of rotation
	data, err := r.managedClustersSvc.GetCredentials(ctx, managedClusterSpec)
	if err != nil {
		return errors.Wrapf(err, "failed to get credentials for managed cluster")
	}

	// Construct and store secret
	kubeconfig := makeKubeconfig(scope.Cluster)
	_, err = controllerutil.CreateOrUpdate(ctx, r.kubeclient, kubeconfig, func() error {
		kubeconfig.Data = map[string][]byte{
			secret.KubeconfigDataName: data,
		}
		return nil
	})

	return err
}

// Delete reconciles all the services in pre determined order
func (r *azureManagedControlPlaneReconciler) Delete(ctx context.Context, scope *scope.ManagedControlPlaneScope) error {
	if err := r.managedClustersSvc.Delete(ctx, nil); err != nil {
		return errors.Wrapf(err, "failed to delete managed cluster %s", scope.ControlPlane.Name)
	}

	return nil
}

func makeKubeconfig(cluster *clusterv1.Cluster) *corev1.Secret {
	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secret.Name(cluster.Name, secret.Kubeconfig),
			Namespace: cluster.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				{
					APIVersion: clusterv1.GroupVersion.String(),
					Kind:       "Cluster",
					Name:       cluster.Name,
					UID:        cluster.UID,
				},
			},
		},
	}
}
