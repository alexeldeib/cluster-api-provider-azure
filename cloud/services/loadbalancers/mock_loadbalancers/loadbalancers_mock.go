/*
Copyright The Kubernetes Authors.

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

// Code generated by MockGen. DO NOT EDIT.
// Source: ../service.go

// Package mock_loadbalancers is a generated GoMock package.
package mock_loadbalancers

import (
	autorest "github.com/Azure/go-autorest/autorest"
	logr "github.com/go-logr/logr"
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	reflect "reflect"
	v1alpha3 "sigs.k8s.io/cluster-api-provider-azure/api/v1alpha3"
	azure "sigs.k8s.io/cluster-api-provider-azure/cloud"
	v1alpha30 "sigs.k8s.io/cluster-api/api/v1alpha3"
)

// MockLBScope is a mock of LBScope interface.
type MockLBScope struct {
	ctrl     *gomock.Controller
	recorder *MockLBScopeMockRecorder
}

// MockLBScopeMockRecorder is the mock recorder for MockLBScope.
type MockLBScopeMockRecorder struct {
	mock *MockLBScope
}

// NewMockLBScope creates a new mock instance.
func NewMockLBScope(ctrl *gomock.Controller) *MockLBScope {
	mock := &MockLBScope{ctrl: ctrl}
	mock.recorder = &MockLBScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLBScope) EXPECT() *MockLBScopeMockRecorder {
	return m.recorder
}

// GetNamespace mocks base method.
func (m *MockLBScope) GetNamespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetNamespace indicates an expected call of GetNamespace.
func (mr *MockLBScopeMockRecorder) GetNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockLBScope)(nil).GetNamespace))
}

// SetNamespace mocks base method.
func (m *MockLBScope) SetNamespace(namespace string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetNamespace", namespace)
}

// SetNamespace indicates an expected call of SetNamespace.
func (mr *MockLBScopeMockRecorder) SetNamespace(namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNamespace", reflect.TypeOf((*MockLBScope)(nil).SetNamespace), namespace)
}

// GetName mocks base method.
func (m *MockLBScope) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockLBScopeMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockLBScope)(nil).GetName))
}

// SetName mocks base method.
func (m *MockLBScope) SetName(name string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetName", name)
}

// SetName indicates an expected call of SetName.
func (mr *MockLBScopeMockRecorder) SetName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetName", reflect.TypeOf((*MockLBScope)(nil).SetName), name)
}

// GetGenerateName mocks base method.
func (m *MockLBScope) GetGenerateName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGenerateName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetGenerateName indicates an expected call of GetGenerateName.
func (mr *MockLBScopeMockRecorder) GetGenerateName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenerateName", reflect.TypeOf((*MockLBScope)(nil).GetGenerateName))
}

// SetGenerateName mocks base method.
func (m *MockLBScope) SetGenerateName(name string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetGenerateName", name)
}

// SetGenerateName indicates an expected call of SetGenerateName.
func (mr *MockLBScopeMockRecorder) SetGenerateName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGenerateName", reflect.TypeOf((*MockLBScope)(nil).SetGenerateName), name)
}

// GetUID mocks base method.
func (m *MockLBScope) GetUID() types.UID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUID")
	ret0, _ := ret[0].(types.UID)
	return ret0
}

// GetUID indicates an expected call of GetUID.
func (mr *MockLBScopeMockRecorder) GetUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUID", reflect.TypeOf((*MockLBScope)(nil).GetUID))
}

// SetUID mocks base method.
func (m *MockLBScope) SetUID(uid types.UID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUID", uid)
}

// SetUID indicates an expected call of SetUID.
func (mr *MockLBScopeMockRecorder) SetUID(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUID", reflect.TypeOf((*MockLBScope)(nil).SetUID), uid)
}

// GetResourceVersion mocks base method.
func (m *MockLBScope) GetResourceVersion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetResourceVersion indicates an expected call of GetResourceVersion.
func (mr *MockLBScopeMockRecorder) GetResourceVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceVersion", reflect.TypeOf((*MockLBScope)(nil).GetResourceVersion))
}

// SetResourceVersion mocks base method.
func (m *MockLBScope) SetResourceVersion(version string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetResourceVersion", version)
}

// SetResourceVersion indicates an expected call of SetResourceVersion.
func (mr *MockLBScopeMockRecorder) SetResourceVersion(version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetResourceVersion", reflect.TypeOf((*MockLBScope)(nil).SetResourceVersion), version)
}

// GetGeneration mocks base method.
func (m *MockLBScope) GetGeneration() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGeneration")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetGeneration indicates an expected call of GetGeneration.
func (mr *MockLBScopeMockRecorder) GetGeneration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGeneration", reflect.TypeOf((*MockLBScope)(nil).GetGeneration))
}

// SetGeneration mocks base method.
func (m *MockLBScope) SetGeneration(generation int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetGeneration", generation)
}

// SetGeneration indicates an expected call of SetGeneration.
func (mr *MockLBScopeMockRecorder) SetGeneration(generation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGeneration", reflect.TypeOf((*MockLBScope)(nil).SetGeneration), generation)
}

// GetSelfLink mocks base method.
func (m *MockLBScope) GetSelfLink() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSelfLink")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSelfLink indicates an expected call of GetSelfLink.
func (mr *MockLBScopeMockRecorder) GetSelfLink() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSelfLink", reflect.TypeOf((*MockLBScope)(nil).GetSelfLink))
}

// SetSelfLink mocks base method.
func (m *MockLBScope) SetSelfLink(selfLink string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSelfLink", selfLink)
}

// SetSelfLink indicates an expected call of SetSelfLink.
func (mr *MockLBScopeMockRecorder) SetSelfLink(selfLink interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSelfLink", reflect.TypeOf((*MockLBScope)(nil).SetSelfLink), selfLink)
}

// GetCreationTimestamp mocks base method.
func (m *MockLBScope) GetCreationTimestamp() v1.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCreationTimestamp")
	ret0, _ := ret[0].(v1.Time)
	return ret0
}

// GetCreationTimestamp indicates an expected call of GetCreationTimestamp.
func (mr *MockLBScopeMockRecorder) GetCreationTimestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreationTimestamp", reflect.TypeOf((*MockLBScope)(nil).GetCreationTimestamp))
}

// SetCreationTimestamp mocks base method.
func (m *MockLBScope) SetCreationTimestamp(timestamp v1.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCreationTimestamp", timestamp)
}

// SetCreationTimestamp indicates an expected call of SetCreationTimestamp.
func (mr *MockLBScopeMockRecorder) SetCreationTimestamp(timestamp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCreationTimestamp", reflect.TypeOf((*MockLBScope)(nil).SetCreationTimestamp), timestamp)
}

// GetDeletionTimestamp mocks base method.
func (m *MockLBScope) GetDeletionTimestamp() *v1.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeletionTimestamp")
	ret0, _ := ret[0].(*v1.Time)
	return ret0
}

// GetDeletionTimestamp indicates an expected call of GetDeletionTimestamp.
func (mr *MockLBScopeMockRecorder) GetDeletionTimestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeletionTimestamp", reflect.TypeOf((*MockLBScope)(nil).GetDeletionTimestamp))
}

// SetDeletionTimestamp mocks base method.
func (m *MockLBScope) SetDeletionTimestamp(timestamp *v1.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDeletionTimestamp", timestamp)
}

// SetDeletionTimestamp indicates an expected call of SetDeletionTimestamp.
func (mr *MockLBScopeMockRecorder) SetDeletionTimestamp(timestamp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeletionTimestamp", reflect.TypeOf((*MockLBScope)(nil).SetDeletionTimestamp), timestamp)
}

// GetDeletionGracePeriodSeconds mocks base method.
func (m *MockLBScope) GetDeletionGracePeriodSeconds() *int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeletionGracePeriodSeconds")
	ret0, _ := ret[0].(*int64)
	return ret0
}

// GetDeletionGracePeriodSeconds indicates an expected call of GetDeletionGracePeriodSeconds.
func (mr *MockLBScopeMockRecorder) GetDeletionGracePeriodSeconds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeletionGracePeriodSeconds", reflect.TypeOf((*MockLBScope)(nil).GetDeletionGracePeriodSeconds))
}

// SetDeletionGracePeriodSeconds mocks base method.
func (m *MockLBScope) SetDeletionGracePeriodSeconds(arg0 *int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDeletionGracePeriodSeconds", arg0)
}

// SetDeletionGracePeriodSeconds indicates an expected call of SetDeletionGracePeriodSeconds.
func (mr *MockLBScopeMockRecorder) SetDeletionGracePeriodSeconds(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeletionGracePeriodSeconds", reflect.TypeOf((*MockLBScope)(nil).SetDeletionGracePeriodSeconds), arg0)
}

// GetLabels mocks base method.
func (m *MockLBScope) GetLabels() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLabels")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetLabels indicates an expected call of GetLabels.
func (mr *MockLBScopeMockRecorder) GetLabels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLabels", reflect.TypeOf((*MockLBScope)(nil).GetLabels))
}

// SetLabels mocks base method.
func (m *MockLBScope) SetLabels(labels map[string]string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLabels", labels)
}

// SetLabels indicates an expected call of SetLabels.
func (mr *MockLBScopeMockRecorder) SetLabels(labels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLabels", reflect.TypeOf((*MockLBScope)(nil).SetLabels), labels)
}

// GetAnnotations mocks base method.
func (m *MockLBScope) GetAnnotations() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAnnotations")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetAnnotations indicates an expected call of GetAnnotations.
func (mr *MockLBScopeMockRecorder) GetAnnotations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnnotations", reflect.TypeOf((*MockLBScope)(nil).GetAnnotations))
}

// SetAnnotations mocks base method.
func (m *MockLBScope) SetAnnotations(annotations map[string]string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAnnotations", annotations)
}

// SetAnnotations indicates an expected call of SetAnnotations.
func (mr *MockLBScopeMockRecorder) SetAnnotations(annotations interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAnnotations", reflect.TypeOf((*MockLBScope)(nil).SetAnnotations), annotations)
}

// GetFinalizers mocks base method.
func (m *MockLBScope) GetFinalizers() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFinalizers")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetFinalizers indicates an expected call of GetFinalizers.
func (mr *MockLBScopeMockRecorder) GetFinalizers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFinalizers", reflect.TypeOf((*MockLBScope)(nil).GetFinalizers))
}

// SetFinalizers mocks base method.
func (m *MockLBScope) SetFinalizers(finalizers []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetFinalizers", finalizers)
}

// SetFinalizers indicates an expected call of SetFinalizers.
func (mr *MockLBScopeMockRecorder) SetFinalizers(finalizers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFinalizers", reflect.TypeOf((*MockLBScope)(nil).SetFinalizers), finalizers)
}

// GetOwnerReferences mocks base method.
func (m *MockLBScope) GetOwnerReferences() []v1.OwnerReference {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOwnerReferences")
	ret0, _ := ret[0].([]v1.OwnerReference)
	return ret0
}

// GetOwnerReferences indicates an expected call of GetOwnerReferences.
func (mr *MockLBScopeMockRecorder) GetOwnerReferences() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOwnerReferences", reflect.TypeOf((*MockLBScope)(nil).GetOwnerReferences))
}

// SetOwnerReferences mocks base method.
func (m *MockLBScope) SetOwnerReferences(arg0 []v1.OwnerReference) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetOwnerReferences", arg0)
}

// SetOwnerReferences indicates an expected call of SetOwnerReferences.
func (mr *MockLBScopeMockRecorder) SetOwnerReferences(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOwnerReferences", reflect.TypeOf((*MockLBScope)(nil).SetOwnerReferences), arg0)
}

// GetClusterName mocks base method.
func (m *MockLBScope) GetClusterName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetClusterName indicates an expected call of GetClusterName.
func (mr *MockLBScopeMockRecorder) GetClusterName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterName", reflect.TypeOf((*MockLBScope)(nil).GetClusterName))
}

// SetClusterName mocks base method.
func (m *MockLBScope) SetClusterName(clusterName string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetClusterName", clusterName)
}

// SetClusterName indicates an expected call of SetClusterName.
func (mr *MockLBScopeMockRecorder) SetClusterName(clusterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClusterName", reflect.TypeOf((*MockLBScope)(nil).SetClusterName), clusterName)
}

// GetManagedFields mocks base method.
func (m *MockLBScope) GetManagedFields() []v1.ManagedFieldsEntry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetManagedFields")
	ret0, _ := ret[0].([]v1.ManagedFieldsEntry)
	return ret0
}

// GetManagedFields indicates an expected call of GetManagedFields.
func (mr *MockLBScopeMockRecorder) GetManagedFields() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManagedFields", reflect.TypeOf((*MockLBScope)(nil).GetManagedFields))
}

// SetManagedFields mocks base method.
func (m *MockLBScope) SetManagedFields(managedFields []v1.ManagedFieldsEntry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetManagedFields", managedFields)
}

// SetManagedFields indicates an expected call of SetManagedFields.
func (mr *MockLBScopeMockRecorder) SetManagedFields(managedFields interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetManagedFields", reflect.TypeOf((*MockLBScope)(nil).SetManagedFields), managedFields)
}

// GetObjectKind mocks base method.
func (m *MockLBScope) GetObjectKind() schema.ObjectKind {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectKind")
	ret0, _ := ret[0].(schema.ObjectKind)
	return ret0
}

// GetObjectKind indicates an expected call of GetObjectKind.
func (mr *MockLBScopeMockRecorder) GetObjectKind() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectKind", reflect.TypeOf((*MockLBScope)(nil).GetObjectKind))
}

// DeepCopyObject mocks base method.
func (m *MockLBScope) DeepCopyObject() runtime.Object {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeepCopyObject")
	ret0, _ := ret[0].(runtime.Object)
	return ret0
}

// DeepCopyObject indicates an expected call of DeepCopyObject.
func (mr *MockLBScopeMockRecorder) DeepCopyObject() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeepCopyObject", reflect.TypeOf((*MockLBScope)(nil).DeepCopyObject))
}

// LoadBalancerName mocks base method.
func (m *MockLBScope) LoadBalancerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadBalancerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// LoadBalancerName indicates an expected call of LoadBalancerName.
func (mr *MockLBScopeMockRecorder) LoadBalancerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadBalancerName", reflect.TypeOf((*MockLBScope)(nil).LoadBalancerName))
}

// Network mocks base method.
func (m *MockLBScope) Network() *v1alpha3.Network {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Network")
	ret0, _ := ret[0].(*v1alpha3.Network)
	return ret0
}

// Network indicates an expected call of Network.
func (mr *MockLBScopeMockRecorder) Network() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Network", reflect.TypeOf((*MockLBScope)(nil).Network))
}

// Vnet mocks base method.
func (m *MockLBScope) Vnet() *v1alpha3.VnetSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Vnet")
	ret0, _ := ret[0].(*v1alpha3.VnetSpec)
	return ret0
}

// Vnet indicates an expected call of Vnet.
func (mr *MockLBScopeMockRecorder) Vnet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Vnet", reflect.TypeOf((*MockLBScope)(nil).Vnet))
}

// IsVnetManaged mocks base method.
func (m *MockLBScope) IsVnetManaged() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsVnetManaged")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsVnetManaged indicates an expected call of IsVnetManaged.
func (mr *MockLBScopeMockRecorder) IsVnetManaged() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsVnetManaged", reflect.TypeOf((*MockLBScope)(nil).IsVnetManaged))
}

// Subnets mocks base method.
func (m *MockLBScope) Subnets() v1alpha3.Subnets {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subnets")
	ret0, _ := ret[0].(v1alpha3.Subnets)
	return ret0
}

// Subnets indicates an expected call of Subnets.
func (mr *MockLBScopeMockRecorder) Subnets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subnets", reflect.TypeOf((*MockLBScope)(nil).Subnets))
}

// NodeSubnet mocks base method.
func (m *MockLBScope) NodeSubnet() *v1alpha3.SubnetSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeSubnet")
	ret0, _ := ret[0].(*v1alpha3.SubnetSpec)
	return ret0
}

// NodeSubnet indicates an expected call of NodeSubnet.
func (mr *MockLBScopeMockRecorder) NodeSubnet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeSubnet", reflect.TypeOf((*MockLBScope)(nil).NodeSubnet))
}

// ControlPlaneSubnet mocks base method.
func (m *MockLBScope) ControlPlaneSubnet() *v1alpha3.SubnetSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControlPlaneSubnet")
	ret0, _ := ret[0].(*v1alpha3.SubnetSpec)
	return ret0
}

// ControlPlaneSubnet indicates an expected call of ControlPlaneSubnet.
func (mr *MockLBScopeMockRecorder) ControlPlaneSubnet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControlPlaneSubnet", reflect.TypeOf((*MockLBScope)(nil).ControlPlaneSubnet))
}

// RouteTable mocks base method.
func (m *MockLBScope) RouteTable() *v1alpha3.RouteTable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RouteTable")
	ret0, _ := ret[0].(*v1alpha3.RouteTable)
	return ret0
}

// RouteTable indicates an expected call of RouteTable.
func (mr *MockLBScopeMockRecorder) RouteTable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteTable", reflect.TypeOf((*MockLBScope)(nil).RouteTable))
}

// SubscriptionID mocks base method.
func (m *MockLBScope) SubscriptionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SubscriptionID indicates an expected call of SubscriptionID.
func (mr *MockLBScopeMockRecorder) SubscriptionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionID", reflect.TypeOf((*MockLBScope)(nil).SubscriptionID))
}

// ResourceGroup mocks base method.
func (m *MockLBScope) ResourceGroup() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceGroup")
	ret0, _ := ret[0].(string)
	return ret0
}

// ResourceGroup indicates an expected call of ResourceGroup.
func (mr *MockLBScopeMockRecorder) ResourceGroup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceGroup", reflect.TypeOf((*MockLBScope)(nil).ResourceGroup))
}

// ClusterName mocks base method.
func (m *MockLBScope) ClusterName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClusterName indicates an expected call of ClusterName.
func (mr *MockLBScopeMockRecorder) ClusterName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterName", reflect.TypeOf((*MockLBScope)(nil).ClusterName))
}

// Location mocks base method.
func (m *MockLBScope) Location() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Location")
	ret0, _ := ret[0].(string)
	return ret0
}

// Location indicates an expected call of Location.
func (mr *MockLBScopeMockRecorder) Location() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Location", reflect.TypeOf((*MockLBScope)(nil).Location))
}

// SetFailureDomain mocks base method.
func (m *MockLBScope) SetFailureDomain(id string, spec v1alpha30.FailureDomainSpec) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetFailureDomain", id, spec)
}

// SetFailureDomain indicates an expected call of SetFailureDomain.
func (mr *MockLBScopeMockRecorder) SetFailureDomain(id, spec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFailureDomain", reflect.TypeOf((*MockLBScope)(nil).SetFailureDomain), id, spec)
}

// AdditionalTags mocks base method.
func (m *MockLBScope) AdditionalTags() v1alpha3.Tags {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdditionalTags")
	ret0, _ := ret[0].(v1alpha3.Tags)
	return ret0
}

// AdditionalTags indicates an expected call of AdditionalTags.
func (mr *MockLBScopeMockRecorder) AdditionalTags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdditionalTags", reflect.TypeOf((*MockLBScope)(nil).AdditionalTags))
}

// ClientID mocks base method.
func (m *MockLBScope) ClientID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientID indicates an expected call of ClientID.
func (mr *MockLBScopeMockRecorder) ClientID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientID", reflect.TypeOf((*MockLBScope)(nil).ClientID))
}

// ClientSecret mocks base method.
func (m *MockLBScope) ClientSecret() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientSecret")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientSecret indicates an expected call of ClientSecret.
func (mr *MockLBScopeMockRecorder) ClientSecret() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientSecret", reflect.TypeOf((*MockLBScope)(nil).ClientSecret))
}

// CloudEnvironment mocks base method.
func (m *MockLBScope) CloudEnvironment() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudEnvironment")
	ret0, _ := ret[0].(string)
	return ret0
}

// CloudEnvironment indicates an expected call of CloudEnvironment.
func (mr *MockLBScopeMockRecorder) CloudEnvironment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudEnvironment", reflect.TypeOf((*MockLBScope)(nil).CloudEnvironment))
}

// TenantID mocks base method.
func (m *MockLBScope) TenantID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantID")
	ret0, _ := ret[0].(string)
	return ret0
}

// TenantID indicates an expected call of TenantID.
func (mr *MockLBScopeMockRecorder) TenantID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantID", reflect.TypeOf((*MockLBScope)(nil).TenantID))
}

// BaseURI mocks base method.
func (m *MockLBScope) BaseURI() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseURI")
	ret0, _ := ret[0].(string)
	return ret0
}

// BaseURI indicates an expected call of BaseURI.
func (mr *MockLBScopeMockRecorder) BaseURI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseURI", reflect.TypeOf((*MockLBScope)(nil).BaseURI))
}

// Authorizer mocks base method.
func (m *MockLBScope) Authorizer() autorest.Authorizer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorizer")
	ret0, _ := ret[0].(autorest.Authorizer)
	return ret0
}

// Authorizer indicates an expected call of Authorizer.
func (mr *MockLBScopeMockRecorder) Authorizer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorizer", reflect.TypeOf((*MockLBScope)(nil).Authorizer))
}

// Info mocks base method.
func (m *MockLBScope) Info(msg string, keysAndValues ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{msg}
	for _, a := range keysAndValues {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info.
func (mr *MockLBScopeMockRecorder) Info(msg interface{}, keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{msg}, keysAndValues...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockLBScope)(nil).Info), varargs...)
}

// Enabled mocks base method.
func (m *MockLBScope) Enabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Enabled indicates an expected call of Enabled.
func (mr *MockLBScopeMockRecorder) Enabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enabled", reflect.TypeOf((*MockLBScope)(nil).Enabled))
}

// Error mocks base method.
func (m *MockLBScope) Error(err error, msg string, keysAndValues ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{err, msg}
	for _, a := range keysAndValues {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error.
func (mr *MockLBScopeMockRecorder) Error(err, msg interface{}, keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{err, msg}, keysAndValues...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLBScope)(nil).Error), varargs...)
}

// V mocks base method.
func (m *MockLBScope) V(level int) logr.InfoLogger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V", level)
	ret0, _ := ret[0].(logr.InfoLogger)
	return ret0
}

// V indicates an expected call of V.
func (mr *MockLBScopeMockRecorder) V(level interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V", reflect.TypeOf((*MockLBScope)(nil).V), level)
}

// WithValues mocks base method.
func (m *MockLBScope) WithValues(keysAndValues ...interface{}) logr.Logger {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range keysAndValues {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WithValues", varargs...)
	ret0, _ := ret[0].(logr.Logger)
	return ret0
}

// WithValues indicates an expected call of WithValues.
func (mr *MockLBScopeMockRecorder) WithValues(keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithValues", reflect.TypeOf((*MockLBScope)(nil).WithValues), keysAndValues...)
}

// WithName mocks base method.
func (m *MockLBScope) WithName(name string) logr.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithName", name)
	ret0, _ := ret[0].(logr.Logger)
	return ret0
}

// WithName indicates an expected call of WithName.
func (mr *MockLBScopeMockRecorder) WithName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithName", reflect.TypeOf((*MockLBScope)(nil).WithName), name)
}

// LBSpecs mocks base method.
func (m *MockLBScope) LBSpecs() []azure.LBSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LBSpecs")
	ret0, _ := ret[0].([]azure.LBSpec)
	return ret0
}

// LBSpecs indicates an expected call of LBSpecs.
func (mr *MockLBScopeMockRecorder) LBSpecs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LBSpecs", reflect.TypeOf((*MockLBScope)(nil).LBSpecs))
}
