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

package scalesets

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-12-01/compute"
	"github.com/Azure/go-autorest/autorest"
	azure "sigs.k8s.io/cluster-api-provider-azure/cloud"
)

// Client wraps go-sdk
type Client interface {
	List(context.Context, string) ([]interface{}, error)
}

// AzureClient contains the Azure go-sdk Client
type AzureClient struct {
	scalesets compute.VirtualMachineScaleSetsClient
}

var _ Client = &AzureClient{}

// NewClient creates a new VMSS client from subscription ID.
func NewClient(subscriptionID string, authorizer autorest.Authorizer) *AzureClient {
	c := newVirtualMachineScaleSetsClient(subscriptionID, authorizer)
	return &AzureClient{c}
}

// newVirtualMachineScaleSetsClient creates a new vmss client from subscription ID.
func newVirtualMachineScaleSetsClient(subscriptionID string, authorizer autorest.Authorizer) compute.VirtualMachineScaleSetsClient {
	scaleSetsClient := compute.NewVirtualMachineScaleSetsClient(subscriptionID)
	scaleSetsClient.Authorizer = authorizer
	scaleSetsClient.AddToUserAgent(azure.UserAgent)
	return scaleSetsClient
}

// Lists all scale sets in a resource group.
func (ac *AzureClient) List(ctx context.Context, resourceGroupName string) ([]interface{}, error) {
	itr, err := ac.scalesets.ListComplete(ctx, resourceGroupName)
	var instances []interface{}
	for ; itr.NotDone(); err = itr.NextWithContext(ctx) {
		if err != nil {
			return nil, fmt.Errorf("failed to iterate vm scale sets [%w]", err)
		}
		instances = append(instances, itr.Value())
	}
	return instances, nil
}
