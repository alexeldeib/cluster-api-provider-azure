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
	"github.com/Azure/go-autorest/autorest"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	azure "sigs.k8s.io/cluster-api-provider-azure/cloud"
	"sigs.k8s.io/cluster-api-provider-azure/cloud/services/managedclusters"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha3"
	"sigs.k8s.io/cluster-api/util/secret"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// azureManagedClusterReconciler are list of services required by cluster controller
type azureManagedClusterReconciler struct {
	kubeclient         client.Client
	managedClustersSvc azure.CredentialGetter
}

// newAzureManagedClusterReconciler populates all the services based on input scope
func newAzureManagedClusterReconciler(subscriptionID string, authorizer autorest.Authorizer, kubeclient client.Client) *azureManagedClusterReconciler {
	return &azureManagedClusterReconciler{
		kubeclient:         kubeclient,
		managedClustersSvc: managedclusters.NewService(authorizer, subscriptionID),
	}
}

// Reconcile reconciles all the services in pre determined order
func (r *azureManagedClusterReconciler) Reconcile(ctx context.Context, scope *ManagedClusterContext) error {
	scope.log.Info("reconciling cluster")
	managedClusterSpec := &managedclusters.Spec{
		Name:            scope.aksCluster.Name,
		ResourceGroup:   scope.aksCluster.Spec.ResourceGroup,
		Location:        scope.aksCluster.Spec.Location,
		Tags:            scope.aksCluster.Spec.AdditionalTags,
		Version:         scope.aksCluster.Spec.Version,
		LoadBalancerSKU: scope.aksCluster.Spec.LoadBalancerSKU,
		NetworkPlugin:   scope.aksCluster.Spec.NetworkPlugin,
		NetworkPolicy:   scope.aksCluster.Spec.NetworkPolicy,
		SSHPublicKey:    scope.aksCluster.Spec.SSHPublicKey,
	}

	_, err := r.managedClustersSvc.Get(ctx, managedClusterSpec)
	if err != nil && !azure.ResourceNotFound(err) {
		return errors.Wrapf(err, "failed to reconcile managed cluster %s", scope.aksCluster.Name)
	}

	// Configure the default pool, rest will be handled by machinepool controller
	defaultPoolSpec := managedclusters.PoolSpec{
		Name:         scope.infraPool.Name,
		SKU:          scope.infraPool.Spec.SKU,
		Replicas:     1,
		OSDiskSizeGB: 0,
	}

	if scope.infraPool.Spec.OSDiskSizeGB != nil {
		defaultPoolSpec.OSDiskSizeGB = *scope.infraPool.Spec.OSDiskSizeGB
	}
	if scope.ownerPool.Spec.Replicas != nil {
		defaultPoolSpec.Replicas = *scope.ownerPool.Spec.Replicas
	}

	managedClusterSpec.AgentPools = []managedclusters.PoolSpec{defaultPoolSpec}

	if azure.ResourceNotFound(err) {
		// We are creating this cluster for the first time.
		if err := r.managedClustersSvc.Reconcile(ctx, managedClusterSpec); err != nil {
			return errors.Wrapf(err, "failed to reconcile managed cluster %s", scope.aksCluster.Name)
		}
		scope.log.Info("reconciled managed cluster successfully")
		return nil
	}

	// Fetched newly created credentials
	managedClusterResult, err := r.managedClustersSvc.Get(ctx, managedClusterSpec)
	if err != nil {
		return err
	}

	managedCluster, ok := managedClusterResult.(containerservice.ManagedCluster)
	if !ok {
		return fmt.Errorf("expected containerservice ManagedCluster object")
	}

	scope.aksCluster.Spec.ControlPlaneEndpoint = clusterv1.APIEndpoint{
		Host: *managedCluster.ManagedClusterProperties.Fqdn,
		Port: 443,
	}

	// Fetched newly created credentials
	data, err := r.managedClustersSvc.GetCredentials(ctx, managedClusterSpec)
	if err != nil {
		return err
	}

	// Construct and store secret
	kubeconfig := makeKubeconfig(scope.ownerCluster)
	_, err = controllerutil.CreateOrUpdate(ctx, r.kubeclient, kubeconfig, func() error {
		kubeconfig.Data = map[string][]byte{
			secret.KubeconfigDataName: data,
		}
		return nil
	})

	// if err != nil {
	// 	scope.log.Error(err)
	// }

	return err
}

// Delete reconciles all the services in pre determined order
func (r *azureManagedClusterReconciler) Delete(ctx context.Context, scope *ManagedClusterContext) error {
	if err := r.managedClustersSvc.Delete(ctx, nil); err != nil {
		return errors.Wrapf(err, "failed to delete managed cluster %s", scope.aksCluster.Name)
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
