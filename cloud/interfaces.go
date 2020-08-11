/*
Copyright 2018 The Kubernetes Authors.

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

package azure

import (
	"context"

	"github.com/Azure/go-autorest/autorest"
	infrav1 "sigs.k8s.io/cluster-api-provider-azure/api/v1alpha3"
	expv1 "sigs.k8s.io/cluster-api-provider-azure/exp/api/v1alpha3"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha3"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// Service is a generic interface used by components offering a type of service.
// Example: virtualnetworks service would offer Reconcile/Delete methods.
type Service interface {
	Reconcile(ctx context.Context) error
	Delete(ctx context.Context) error
}

// OldService is a generic interface for services that have not yet been refactored.
// Once all services have been converted to use Service, this should be removed.
// Example: virtualnetworks service would offer Reconcile/Delete methods.
type OldService interface {
	Reconcile(ctx context.Context, spec interface{}) error
	Delete(ctx context.Context, spec interface{}) error
}

// GetterService is a temporary interface used by components which still require Get methods.
// Once all components move to storing provider information within the relevant
// Cluster/Machine specs, this interface should be removed.
type GetterService interface {
	Get(ctx context.Context, spec interface{}) (interface{}, error)
}

// CredentialGetter is a Service which knows how to retrieve credentials for an Azure
// resource in a resource group.
type CredentialGetter interface {
	Service
	GetCredentials(ctx context.Context, group string, cluster string) ([]byte, error)
}

// Authorizer is an interface which can get the subscription ID, base URI, and authorizer for an Azure service.
type Authorizer interface {
	ClientID() string
	ClientSecret() string
	CloudEnvironment() string
	TenantID() string
	BaseURI() string
	Authorizer() autorest.Authorizer
}

// SubscriptionAuthroizer contains a subscription ID and a means to
// authorize access to it.
type SubscriptionAuthorizer interface {
	Authorizer
	SubscriptionID() string
}

// AuthorizedClusterDescriber describes a complete cluster and can
// authorize requests to mutate/get the backing resources.
type AuthorizedClusterDescriber interface {
	Authorizer
	ClusterDescriber
}

// ClusterDescriber is an interface which can get common Azure Cluster information.
type ClusterDescriber interface {
	controllerutil.Object
	NetworkDescriber
	ClusterName() string
	SubscriptionID() string
	ControlPlaneResourceGroup() string
	NodeResourceGroup() string
	Location() string
	SetFailureDomain(string, clusterv1.FailureDomainSpec)
	AdditionalTags() infrav1.Tags
}

// NetworkDescriber describes the vnet and subnet configuration for a cluster
// abstracted because it is implemented by managed and unmanaged cluster.
type NetworkDescriber interface {
	LoadBalancer() string
	OutboundPool() string
	Network() *infrav1.Network
	Subnets() infrav1.Subnets
	Vnet() *infrav1.VnetSpec
	IsVnetManaged() bool
	NodeSubnet() *infrav1.SubnetSpec
	ControlPlaneSubnet() *infrav1.SubnetSpec
	RouteTable() *infrav1.RouteTable
}

var _ ClusterDescriber = new(infrav1.AzureCluster)
var _ ClusterDescriber = new(expv1.AzureManagedControlPlane)
