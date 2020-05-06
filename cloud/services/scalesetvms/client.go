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

package scalesetvms

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-12-01/compute"
	"github.com/Azure/go-autorest/autorest"
	azure "sigs.k8s.io/cluster-api-provider-azure/cloud"
)

// Client wraps go-sdk
type Client interface {
	ListInstances(context.Context, string, string) ([]string, error)
}

// AzureClient contains the Azure go-sdk Client
type AzureClient struct {
	scalesetvms compute.VirtualMachineScaleSetVMsClient
}

var _ Client = &AzureClient{}

// NewClient creates a new VMSS client from subscription ID.
func NewClient(subscriptionID string, authorizer autorest.Authorizer) *AzureClient {
	c := newVirtualMachineScaleSetVMsClient(subscriptionID, authorizer)
	return &AzureClient{c}
}

// newVirtualMachineScaleSetVMsClient creates a new vmss client from subscription ID.
func newVirtualMachineScaleSetVMsClient(subscriptionID string, authorizer autorest.Authorizer) compute.VirtualMachineScaleSetVMsClient {
	scaleSetVMsClient := compute.NewVirtualMachineScaleSetVMsClient(subscriptionID)
	scaleSetVMsClient.Authorizer = authorizer
	scaleSetVMsClient.AddToUserAgent(azure.UserAgent)
	return scaleSetVMsClient
}

// Lists all instance IDs in a VM scale set.
func (ac *AzureClient) ListInstances(ctx context.Context, resourceGroupName, name string) ([]string, error) {
	itr, err := ac.scalesetvms.ListComplete(ctx, resourceGroupName, name, "", "", "")
	var instances []string
	for ; itr.NotDone(); err = itr.NextWithContext(ctx) {
		if err != nil {
			return nil, fmt.Errorf("failed to iterate vm scale sets [%w]", err)
		}
		vm := itr.Value()
		instances = append(instances, fmt.Sprintf("azure://%s", *vm.ID))
	}
	return instances, nil
}
