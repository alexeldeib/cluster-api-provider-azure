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

package roleassignments

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/2019-03-01/authorization/mgmt/authorization"
	"github.com/Azure/go-autorest/autorest"
	azure "sigs.k8s.io/cluster-api-provider-azure/cloud"
)

// Client wraps go-sdk
type Client interface {
	Create(context.Context, string, string, authorization.RoleAssignmentCreateParameters) (authorization.RoleAssignment, error)
}

// AzureClient contains the Azure go-sdk Client
type AzureClient struct {
	roleassignments authorization.RoleAssignmentsClient
}

var _ Client = &AzureClient{}

// NewClient creates a new role assignment client from subscription ID.
func NewClient(auth azure.SubscriptionAuthorizer) *AzureClient {
	c := newRoleAssignmentClient(auth.SubscriptionID(), auth.BaseURI(), auth.Authorizer())
	return &AzureClient{c}
}

// newRoleAssignmentClient creates a role assignments client from subscription ID.
func newRoleAssignmentClient(subscriptionID string, baseURI string, authorizer autorest.Authorizer) authorization.RoleAssignmentsClient {
	roleClient := authorization.NewRoleAssignmentsClientWithBaseURI(baseURI, subscriptionID)
	roleClient.Authorizer = authorizer
	roleClient.AddToUserAgent(azure.UserAgent())
	return roleClient
}

// Create creates a role assignment.
// Parameters:
// scope - the scope of the role assignment to create. The scope can be any REST resource instance. For
// example, use '/subscriptions/{subscription-id}/' for a subscription,
// '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}' for a resource group, and
// '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/providers/{resource-provider}/{resource-type}/{resource-name}'
// for a resource.
// roleAssignmentName - the name of the role assignment to create. It can be any valid GUID.
// parameters - parameters for the role assignment.
func (ac *AzureClient) Create(ctx context.Context, scope string, roleAssignmentName string, parameters authorization.RoleAssignmentCreateParameters) (authorization.RoleAssignment, error) {
	return ac.roleassignments.Create(ctx, scope, roleAssignmentName, parameters)
}
