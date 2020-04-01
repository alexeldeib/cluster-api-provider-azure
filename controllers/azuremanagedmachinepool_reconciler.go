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

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-12-01/compute"
	"github.com/Azure/go-autorest/autorest"
	"github.com/pkg/errors"
	azure "sigs.k8s.io/cluster-api-provider-azure/cloud"
	"sigs.k8s.io/cluster-api-provider-azure/cloud/services/agentpools"
	"sigs.k8s.io/cluster-api-provider-azure/cloud/services/scalesets"
	"sigs.k8s.io/cluster-api-provider-azure/cloud/services/scalesetvms"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// azureManagedMachinePoolReconciler are list of services required by cluster controller
type azureManagedMachinePoolReconciler struct {
	kubeclient     client.Client
	agentPoolsSvc  azure.GetterService
	scaleSetVMsSvc NodeLister
	scaleSetsSvc   Lister
}

// NodeLister is a service interface exclusively for returning a list of VMSS instance provider IDs.
type NodeLister interface {
	ListInstances(context.Context, interface{}) ([]string, error)
}

// Lister is a service interface for returning generic lists.
type Lister interface {
	List(context.Context, interface{}) ([]interface{}, error)
}

// newAzureManagedMachinePoolReconciler populates all the services based on input scope
func newAzureManagedMachinePoolReconciler(subscriptionID string, authorizer autorest.Authorizer, kubeclient client.Client) *azureManagedMachinePoolReconciler {
	return &azureManagedMachinePoolReconciler{
		kubeclient:     kubeclient,
		agentPoolsSvc:  agentpools.NewService(authorizer, subscriptionID),
		scaleSetVMsSvc: scalesetvms.NewService(authorizer, subscriptionID),
		scaleSetsSvc:   scalesets.NewService(authorizer, subscriptionID),
	}
}

// Reconcile reconciles all the services in pre determined order
func (r *azureManagedMachinePoolReconciler) Reconcile(ctx context.Context, scope *ManagedClusterContext) error {
	scope.log.Info("reconciling machine pool")
	agentPoolSpec := &agentpools.Spec{
		Name:          scope.infraPool.Name,
		ResourceGroup: scope.aksCluster.Spec.ResourceGroup,
		Cluster:       scope.aksCluster.Name,
		SKU:           scope.infraPool.Spec.SKU,
		Replicas:      1,
	}

	if scope.infraPool.Spec.OSDiskSizeGB != nil {
		agentPoolSpec.OSDiskSizeGB = *scope.infraPool.Spec.OSDiskSizeGB
	}

	if scope.ownerPool.Spec.Replicas != nil {
		agentPoolSpec.Replicas = *scope.ownerPool.Spec.Replicas
	}

	if err := r.agentPoolsSvc.Reconcile(ctx, agentPoolSpec); err != nil {
		return errors.Wrapf(err, "failed to reconcile machine pool %s", scope.infraPool.Name)
	}

	nodeResourceGroup := fmt.Sprintf("MC_%s_%s_%s", scope.aksCluster.Spec.ResourceGroup, scope.aksCluster.Name, scope.aksCluster.Spec.Location)
	vmss, err := r.scaleSetsSvc.List(ctx, &scalesets.Spec{ResourceGroup: nodeResourceGroup})
	if err != nil {
		return errors.Wrapf(err, "failed to list vmss in resource group %s", nodeResourceGroup)
	}

	var match *compute.VirtualMachineScaleSet
	for _, ss := range vmss {
		ss := ss
		switch scaleset := ss.(type) {
		case compute.VirtualMachineScaleSet:
			if scaleset.Tags["poolName"] != nil && *scaleset.Tags["poolName"] == scope.infraPool.Name {
				match = &scaleset
				break
			}
		default:
			return errors.New("expected vmss but found wrong interface type")
		}
	}

	if match == nil {
		return errors.New("failed to find vm scale set matchinf pool")
	}

	providerIDs, err := r.scaleSetVMsSvc.ListInstances(ctx, &scalesetvms.Spec{
		Name:          *match.Name,
		ResourceGroup: nodeResourceGroup,
	})
	if err != nil {
		return errors.Wrapf(err, "failed to reconcile machine pool %s", scope.infraPool.Name)
	}

	scope.infraPool.Spec.ProviderIDList = providerIDs
	scope.infraPool.Status.Replicas = int32(len(providerIDs))
	scope.infraPool.Status.Ready = true

	scope.log.Info("reconciled machine pool successfully")
	return nil
}

// Delete reconciles all the services in pre determined order
func (r *azureManagedMachinePoolReconciler) Delete(ctx context.Context, scope *ManagedClusterContext) error {
	agentPoolSpec := agentpools.Spec{
		Name:          scope.infraPool.Name,
		ResourceGroup: scope.aksCluster.Spec.ResourceGroup,
		Cluster:       scope.aksCluster.Name,
		SKU:           scope.infraPool.Spec.SKU,
	}

	if err := r.agentPoolsSvc.Delete(ctx, agentPoolSpec); err != nil {
		return errors.Wrapf(err, "failed to delete machine pool %s", scope.infraPool.Name)
	}

	return nil
}
