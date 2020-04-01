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

package v1alpha3

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha3"
)

// AzureManagedClusterSpec defines the desired state of AzureManagedCluster
type AzureManagedClusterSpec struct {
	// Version defines the desired Kubernetes version.
	// +kubebuilder:validation:MinLength:=2
	// +kubebuilder:validation:Pattern:=^(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)([-0-9a-zA-Z_\.+]*)?$
	Version string `json:"version"`

	// ResourceGroup is the name of the Azure resource group for this AKS Cluster.
	ResourceGroup string `json:"resourceGroup"`

	// Location is a string matching one of the canonical Azure region names. Examples: "westus2", "eastus".
	Location string `json:"location"`

	// ControlPlaneEndpoint represents the endpoint used to communicate with the control plane.
	// +optional
	ControlPlaneEndpoint clusterv1.APIEndpoint `json:"controlPlaneEndpoint"`

	// AdditionalTags is an optional set of tags to add to Azure resources managed by the Azure provider, in addition to the
	// ones added by default.
	// +optional
	AdditionalTags Tags `json:"additionalTags,omitempty"`

	// LoadBalancerSKU for the managed cluster. Possible values include: 'Standard', 'Basic'. Defaults to standard.
	// +kubebuilder:validation:Enum=Standard;Basic
	LoadBalancerSKU *string `json:"loadBalancerSku,omitempty"`

	// NetworkPlugin used for building Kubernetes network. Possible values include: 'Azure', 'Kubenet'. Defaults to Azure.
	// +kubebuilder:validation:Enum=Azure;Kubenet
	NetworkPlugin *string `json:"networkPlugin,omitempty"`

	// NetworkPolicy used for building Kubernetes network. Possible values include: 'NetworkPolicyCalico', 'NetworkPolicyAzure'
	// +kubebuilder:validation:Enum=NetworkPolicyCalico;NetworkPolicyAzure
	NetworkPolicy *string `json:"networkPolicy,omitempty"`

	// SSHPublicKey is a string literal containing an ssh public key.
	SSHPublicKey string `json:"sshPublicKey"`

	// DefaultPoolRef is the specification for the default pool, without which an AKS cluster cannot be created.
	// TODO(ace): consider defaulting and making optional pointer?
	DefaultPoolRef corev1.LocalObjectReference `json:"defaultPoolRef"`
}

// AzureManagedClusterStatus defines the observed state of AzureManagedCluster
type AzureManagedClusterStatus struct {
	// Ready is true when the provider resource is ready.
	// +optional
	Ready bool `json:"ready,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=azuremanagedclusters,scope=Namespaced,categories=cluster-api,shortName=amc
// +kubebuilder:storageversion
// +kubebuilder:subresource:status

// AzureManagedCluster is the Schema for the azuremanagedclusters API
type AzureManagedCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureManagedClusterSpec   `json:"spec,omitempty"`
	Status AzureManagedClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureManagedClusterList contains a list of AzureManagedCluster
type AzureManagedClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureManagedCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AzureManagedCluster{}, &AzureManagedClusterList{})
}
