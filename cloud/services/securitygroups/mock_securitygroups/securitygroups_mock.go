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

// Package mock_securitygroups is a generated GoMock package.
package mock_securitygroups

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

// MockNSGScope is a mock of NSGScope interface.
type MockNSGScope struct {
	ctrl     *gomock.Controller
	recorder *MockNSGScopeMockRecorder
}

// MockNSGScopeMockRecorder is the mock recorder for MockNSGScope.
type MockNSGScopeMockRecorder struct {
	mock *MockNSGScope
}

// NewMockNSGScope creates a new mock instance.
func NewMockNSGScope(ctrl *gomock.Controller) *MockNSGScope {
	mock := &MockNSGScope{ctrl: ctrl}
	mock.recorder = &MockNSGScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNSGScope) EXPECT() *MockNSGScopeMockRecorder {
	return m.recorder
}

// GetNamespace mocks base method.
func (m *MockNSGScope) GetNamespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetNamespace indicates an expected call of GetNamespace.
func (mr *MockNSGScopeMockRecorder) GetNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockNSGScope)(nil).GetNamespace))
}

// SetNamespace mocks base method.
func (m *MockNSGScope) SetNamespace(namespace string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetNamespace", namespace)
}

// SetNamespace indicates an expected call of SetNamespace.
func (mr *MockNSGScopeMockRecorder) SetNamespace(namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNamespace", reflect.TypeOf((*MockNSGScope)(nil).SetNamespace), namespace)
}

// GetName mocks base method.
func (m *MockNSGScope) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockNSGScopeMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockNSGScope)(nil).GetName))
}

// SetName mocks base method.
func (m *MockNSGScope) SetName(name string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetName", name)
}

// SetName indicates an expected call of SetName.
func (mr *MockNSGScopeMockRecorder) SetName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetName", reflect.TypeOf((*MockNSGScope)(nil).SetName), name)
}

// GetGenerateName mocks base method.
func (m *MockNSGScope) GetGenerateName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGenerateName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetGenerateName indicates an expected call of GetGenerateName.
func (mr *MockNSGScopeMockRecorder) GetGenerateName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenerateName", reflect.TypeOf((*MockNSGScope)(nil).GetGenerateName))
}

// SetGenerateName mocks base method.
func (m *MockNSGScope) SetGenerateName(name string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetGenerateName", name)
}

// SetGenerateName indicates an expected call of SetGenerateName.
func (mr *MockNSGScopeMockRecorder) SetGenerateName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGenerateName", reflect.TypeOf((*MockNSGScope)(nil).SetGenerateName), name)
}

// GetUID mocks base method.
func (m *MockNSGScope) GetUID() types.UID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUID")
	ret0, _ := ret[0].(types.UID)
	return ret0
}

// GetUID indicates an expected call of GetUID.
func (mr *MockNSGScopeMockRecorder) GetUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUID", reflect.TypeOf((*MockNSGScope)(nil).GetUID))
}

// SetUID mocks base method.
func (m *MockNSGScope) SetUID(uid types.UID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUID", uid)
}

// SetUID indicates an expected call of SetUID.
func (mr *MockNSGScopeMockRecorder) SetUID(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUID", reflect.TypeOf((*MockNSGScope)(nil).SetUID), uid)
}

// GetResourceVersion mocks base method.
func (m *MockNSGScope) GetResourceVersion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetResourceVersion indicates an expected call of GetResourceVersion.
func (mr *MockNSGScopeMockRecorder) GetResourceVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceVersion", reflect.TypeOf((*MockNSGScope)(nil).GetResourceVersion))
}

// SetResourceVersion mocks base method.
func (m *MockNSGScope) SetResourceVersion(version string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetResourceVersion", version)
}

// SetResourceVersion indicates an expected call of SetResourceVersion.
func (mr *MockNSGScopeMockRecorder) SetResourceVersion(version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetResourceVersion", reflect.TypeOf((*MockNSGScope)(nil).SetResourceVersion), version)
}

// GetGeneration mocks base method.
func (m *MockNSGScope) GetGeneration() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGeneration")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetGeneration indicates an expected call of GetGeneration.
func (mr *MockNSGScopeMockRecorder) GetGeneration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGeneration", reflect.TypeOf((*MockNSGScope)(nil).GetGeneration))
}

// SetGeneration mocks base method.
func (m *MockNSGScope) SetGeneration(generation int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetGeneration", generation)
}

// SetGeneration indicates an expected call of SetGeneration.
func (mr *MockNSGScopeMockRecorder) SetGeneration(generation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGeneration", reflect.TypeOf((*MockNSGScope)(nil).SetGeneration), generation)
}

// GetSelfLink mocks base method.
func (m *MockNSGScope) GetSelfLink() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSelfLink")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSelfLink indicates an expected call of GetSelfLink.
func (mr *MockNSGScopeMockRecorder) GetSelfLink() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSelfLink", reflect.TypeOf((*MockNSGScope)(nil).GetSelfLink))
}

// SetSelfLink mocks base method.
func (m *MockNSGScope) SetSelfLink(selfLink string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSelfLink", selfLink)
}

// SetSelfLink indicates an expected call of SetSelfLink.
func (mr *MockNSGScopeMockRecorder) SetSelfLink(selfLink interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSelfLink", reflect.TypeOf((*MockNSGScope)(nil).SetSelfLink), selfLink)
}

// GetCreationTimestamp mocks base method.
func (m *MockNSGScope) GetCreationTimestamp() v1.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCreationTimestamp")
	ret0, _ := ret[0].(v1.Time)
	return ret0
}

// GetCreationTimestamp indicates an expected call of GetCreationTimestamp.
func (mr *MockNSGScopeMockRecorder) GetCreationTimestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreationTimestamp", reflect.TypeOf((*MockNSGScope)(nil).GetCreationTimestamp))
}

// SetCreationTimestamp mocks base method.
func (m *MockNSGScope) SetCreationTimestamp(timestamp v1.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCreationTimestamp", timestamp)
}

// SetCreationTimestamp indicates an expected call of SetCreationTimestamp.
func (mr *MockNSGScopeMockRecorder) SetCreationTimestamp(timestamp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCreationTimestamp", reflect.TypeOf((*MockNSGScope)(nil).SetCreationTimestamp), timestamp)
}

// GetDeletionTimestamp mocks base method.
func (m *MockNSGScope) GetDeletionTimestamp() *v1.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeletionTimestamp")
	ret0, _ := ret[0].(*v1.Time)
	return ret0
}

// GetDeletionTimestamp indicates an expected call of GetDeletionTimestamp.
func (mr *MockNSGScopeMockRecorder) GetDeletionTimestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeletionTimestamp", reflect.TypeOf((*MockNSGScope)(nil).GetDeletionTimestamp))
}

// SetDeletionTimestamp mocks base method.
func (m *MockNSGScope) SetDeletionTimestamp(timestamp *v1.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDeletionTimestamp", timestamp)
}

// SetDeletionTimestamp indicates an expected call of SetDeletionTimestamp.
func (mr *MockNSGScopeMockRecorder) SetDeletionTimestamp(timestamp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeletionTimestamp", reflect.TypeOf((*MockNSGScope)(nil).SetDeletionTimestamp), timestamp)
}

// GetDeletionGracePeriodSeconds mocks base method.
func (m *MockNSGScope) GetDeletionGracePeriodSeconds() *int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeletionGracePeriodSeconds")
	ret0, _ := ret[0].(*int64)
	return ret0
}

// GetDeletionGracePeriodSeconds indicates an expected call of GetDeletionGracePeriodSeconds.
func (mr *MockNSGScopeMockRecorder) GetDeletionGracePeriodSeconds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeletionGracePeriodSeconds", reflect.TypeOf((*MockNSGScope)(nil).GetDeletionGracePeriodSeconds))
}

// SetDeletionGracePeriodSeconds mocks base method.
func (m *MockNSGScope) SetDeletionGracePeriodSeconds(arg0 *int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDeletionGracePeriodSeconds", arg0)
}

// SetDeletionGracePeriodSeconds indicates an expected call of SetDeletionGracePeriodSeconds.
func (mr *MockNSGScopeMockRecorder) SetDeletionGracePeriodSeconds(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeletionGracePeriodSeconds", reflect.TypeOf((*MockNSGScope)(nil).SetDeletionGracePeriodSeconds), arg0)
}

// GetLabels mocks base method.
func (m *MockNSGScope) GetLabels() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLabels")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetLabels indicates an expected call of GetLabels.
func (mr *MockNSGScopeMockRecorder) GetLabels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLabels", reflect.TypeOf((*MockNSGScope)(nil).GetLabels))
}

// SetLabels mocks base method.
func (m *MockNSGScope) SetLabels(labels map[string]string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLabels", labels)
}

// SetLabels indicates an expected call of SetLabels.
func (mr *MockNSGScopeMockRecorder) SetLabels(labels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLabels", reflect.TypeOf((*MockNSGScope)(nil).SetLabels), labels)
}

// GetAnnotations mocks base method.
func (m *MockNSGScope) GetAnnotations() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAnnotations")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetAnnotations indicates an expected call of GetAnnotations.
func (mr *MockNSGScopeMockRecorder) GetAnnotations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnnotations", reflect.TypeOf((*MockNSGScope)(nil).GetAnnotations))
}

// SetAnnotations mocks base method.
func (m *MockNSGScope) SetAnnotations(annotations map[string]string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAnnotations", annotations)
}

// SetAnnotations indicates an expected call of SetAnnotations.
func (mr *MockNSGScopeMockRecorder) SetAnnotations(annotations interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAnnotations", reflect.TypeOf((*MockNSGScope)(nil).SetAnnotations), annotations)
}

// GetFinalizers mocks base method.
func (m *MockNSGScope) GetFinalizers() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFinalizers")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetFinalizers indicates an expected call of GetFinalizers.
func (mr *MockNSGScopeMockRecorder) GetFinalizers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFinalizers", reflect.TypeOf((*MockNSGScope)(nil).GetFinalizers))
}

// SetFinalizers mocks base method.
func (m *MockNSGScope) SetFinalizers(finalizers []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetFinalizers", finalizers)
}

// SetFinalizers indicates an expected call of SetFinalizers.
func (mr *MockNSGScopeMockRecorder) SetFinalizers(finalizers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFinalizers", reflect.TypeOf((*MockNSGScope)(nil).SetFinalizers), finalizers)
}

// GetOwnerReferences mocks base method.
func (m *MockNSGScope) GetOwnerReferences() []v1.OwnerReference {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOwnerReferences")
	ret0, _ := ret[0].([]v1.OwnerReference)
	return ret0
}

// GetOwnerReferences indicates an expected call of GetOwnerReferences.
func (mr *MockNSGScopeMockRecorder) GetOwnerReferences() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOwnerReferences", reflect.TypeOf((*MockNSGScope)(nil).GetOwnerReferences))
}

// SetOwnerReferences mocks base method.
func (m *MockNSGScope) SetOwnerReferences(arg0 []v1.OwnerReference) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetOwnerReferences", arg0)
}

// SetOwnerReferences indicates an expected call of SetOwnerReferences.
func (mr *MockNSGScopeMockRecorder) SetOwnerReferences(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOwnerReferences", reflect.TypeOf((*MockNSGScope)(nil).SetOwnerReferences), arg0)
}

// GetClusterName mocks base method.
func (m *MockNSGScope) GetClusterName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetClusterName indicates an expected call of GetClusterName.
func (mr *MockNSGScopeMockRecorder) GetClusterName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterName", reflect.TypeOf((*MockNSGScope)(nil).GetClusterName))
}

// SetClusterName mocks base method.
func (m *MockNSGScope) SetClusterName(clusterName string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetClusterName", clusterName)
}

// SetClusterName indicates an expected call of SetClusterName.
func (mr *MockNSGScopeMockRecorder) SetClusterName(clusterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClusterName", reflect.TypeOf((*MockNSGScope)(nil).SetClusterName), clusterName)
}

// GetManagedFields mocks base method.
func (m *MockNSGScope) GetManagedFields() []v1.ManagedFieldsEntry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetManagedFields")
	ret0, _ := ret[0].([]v1.ManagedFieldsEntry)
	return ret0
}

// GetManagedFields indicates an expected call of GetManagedFields.
func (mr *MockNSGScopeMockRecorder) GetManagedFields() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManagedFields", reflect.TypeOf((*MockNSGScope)(nil).GetManagedFields))
}

// SetManagedFields mocks base method.
func (m *MockNSGScope) SetManagedFields(managedFields []v1.ManagedFieldsEntry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetManagedFields", managedFields)
}

// SetManagedFields indicates an expected call of SetManagedFields.
func (mr *MockNSGScopeMockRecorder) SetManagedFields(managedFields interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetManagedFields", reflect.TypeOf((*MockNSGScope)(nil).SetManagedFields), managedFields)
}

// GetObjectKind mocks base method.
func (m *MockNSGScope) GetObjectKind() schema.ObjectKind {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectKind")
	ret0, _ := ret[0].(schema.ObjectKind)
	return ret0
}

// GetObjectKind indicates an expected call of GetObjectKind.
func (mr *MockNSGScopeMockRecorder) GetObjectKind() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectKind", reflect.TypeOf((*MockNSGScope)(nil).GetObjectKind))
}

// DeepCopyObject mocks base method.
func (m *MockNSGScope) DeepCopyObject() runtime.Object {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeepCopyObject")
	ret0, _ := ret[0].(runtime.Object)
	return ret0
}

// DeepCopyObject indicates an expected call of DeepCopyObject.
func (mr *MockNSGScopeMockRecorder) DeepCopyObject() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeepCopyObject", reflect.TypeOf((*MockNSGScope)(nil).DeepCopyObject))
}

// LoadBalancerName mocks base method.
func (m *MockNSGScope) LoadBalancerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadBalancerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// LoadBalancerName indicates an expected call of LoadBalancerName.
func (mr *MockNSGScopeMockRecorder) LoadBalancerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadBalancerName", reflect.TypeOf((*MockNSGScope)(nil).LoadBalancerName))
}

// Network mocks base method.
func (m *MockNSGScope) Network() *v1alpha3.Network {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Network")
	ret0, _ := ret[0].(*v1alpha3.Network)
	return ret0
}

// Network indicates an expected call of Network.
func (mr *MockNSGScopeMockRecorder) Network() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Network", reflect.TypeOf((*MockNSGScope)(nil).Network))
}

// Vnet mocks base method.
func (m *MockNSGScope) Vnet() *v1alpha3.VnetSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Vnet")
	ret0, _ := ret[0].(*v1alpha3.VnetSpec)
	return ret0
}

// Vnet indicates an expected call of Vnet.
func (mr *MockNSGScopeMockRecorder) Vnet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Vnet", reflect.TypeOf((*MockNSGScope)(nil).Vnet))
}

// IsVnetManaged mocks base method.
func (m *MockNSGScope) IsVnetManaged() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsVnetManaged")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsVnetManaged indicates an expected call of IsVnetManaged.
func (mr *MockNSGScopeMockRecorder) IsVnetManaged() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsVnetManaged", reflect.TypeOf((*MockNSGScope)(nil).IsVnetManaged))
}

// Subnets mocks base method.
func (m *MockNSGScope) Subnets() v1alpha3.Subnets {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subnets")
	ret0, _ := ret[0].(v1alpha3.Subnets)
	return ret0
}

// Subnets indicates an expected call of Subnets.
func (mr *MockNSGScopeMockRecorder) Subnets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subnets", reflect.TypeOf((*MockNSGScope)(nil).Subnets))
}

// NodeSubnet mocks base method.
func (m *MockNSGScope) NodeSubnet() *v1alpha3.SubnetSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeSubnet")
	ret0, _ := ret[0].(*v1alpha3.SubnetSpec)
	return ret0
}

// NodeSubnet indicates an expected call of NodeSubnet.
func (mr *MockNSGScopeMockRecorder) NodeSubnet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeSubnet", reflect.TypeOf((*MockNSGScope)(nil).NodeSubnet))
}

// ControlPlaneSubnet mocks base method.
func (m *MockNSGScope) ControlPlaneSubnet() *v1alpha3.SubnetSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControlPlaneSubnet")
	ret0, _ := ret[0].(*v1alpha3.SubnetSpec)
	return ret0
}

// ControlPlaneSubnet indicates an expected call of ControlPlaneSubnet.
func (mr *MockNSGScopeMockRecorder) ControlPlaneSubnet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControlPlaneSubnet", reflect.TypeOf((*MockNSGScope)(nil).ControlPlaneSubnet))
}

// RouteTable mocks base method.
func (m *MockNSGScope) RouteTable() *v1alpha3.RouteTable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RouteTable")
	ret0, _ := ret[0].(*v1alpha3.RouteTable)
	return ret0
}

// RouteTable indicates an expected call of RouteTable.
func (mr *MockNSGScopeMockRecorder) RouteTable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteTable", reflect.TypeOf((*MockNSGScope)(nil).RouteTable))
}

// SubscriptionID mocks base method.
func (m *MockNSGScope) SubscriptionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SubscriptionID indicates an expected call of SubscriptionID.
func (mr *MockNSGScopeMockRecorder) SubscriptionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionID", reflect.TypeOf((*MockNSGScope)(nil).SubscriptionID))
}

// ResourceGroup mocks base method.
func (m *MockNSGScope) ResourceGroup() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceGroup")
	ret0, _ := ret[0].(string)
	return ret0
}

// ResourceGroup indicates an expected call of ResourceGroup.
func (mr *MockNSGScopeMockRecorder) ResourceGroup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceGroup", reflect.TypeOf((*MockNSGScope)(nil).ResourceGroup))
}

// ClusterName mocks base method.
func (m *MockNSGScope) ClusterName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClusterName indicates an expected call of ClusterName.
func (mr *MockNSGScopeMockRecorder) ClusterName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterName", reflect.TypeOf((*MockNSGScope)(nil).ClusterName))
}

// Location mocks base method.
func (m *MockNSGScope) Location() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Location")
	ret0, _ := ret[0].(string)
	return ret0
}

// Location indicates an expected call of Location.
func (mr *MockNSGScopeMockRecorder) Location() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Location", reflect.TypeOf((*MockNSGScope)(nil).Location))
}

// SetFailureDomain mocks base method.
func (m *MockNSGScope) SetFailureDomain(id string, spec v1alpha30.FailureDomainSpec) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetFailureDomain", id, spec)
}

// SetFailureDomain indicates an expected call of SetFailureDomain.
func (mr *MockNSGScopeMockRecorder) SetFailureDomain(id, spec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFailureDomain", reflect.TypeOf((*MockNSGScope)(nil).SetFailureDomain), id, spec)
}

// AdditionalTags mocks base method.
func (m *MockNSGScope) AdditionalTags() v1alpha3.Tags {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdditionalTags")
	ret0, _ := ret[0].(v1alpha3.Tags)
	return ret0
}

// AdditionalTags indicates an expected call of AdditionalTags.
func (mr *MockNSGScopeMockRecorder) AdditionalTags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdditionalTags", reflect.TypeOf((*MockNSGScope)(nil).AdditionalTags))
}

// ClientID mocks base method.
func (m *MockNSGScope) ClientID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientID indicates an expected call of ClientID.
func (mr *MockNSGScopeMockRecorder) ClientID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientID", reflect.TypeOf((*MockNSGScope)(nil).ClientID))
}

// ClientSecret mocks base method.
func (m *MockNSGScope) ClientSecret() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientSecret")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientSecret indicates an expected call of ClientSecret.
func (mr *MockNSGScopeMockRecorder) ClientSecret() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientSecret", reflect.TypeOf((*MockNSGScope)(nil).ClientSecret))
}

// CloudEnvironment mocks base method.
func (m *MockNSGScope) CloudEnvironment() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudEnvironment")
	ret0, _ := ret[0].(string)
	return ret0
}

// CloudEnvironment indicates an expected call of CloudEnvironment.
func (mr *MockNSGScopeMockRecorder) CloudEnvironment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudEnvironment", reflect.TypeOf((*MockNSGScope)(nil).CloudEnvironment))
}

// TenantID mocks base method.
func (m *MockNSGScope) TenantID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantID")
	ret0, _ := ret[0].(string)
	return ret0
}

// TenantID indicates an expected call of TenantID.
func (mr *MockNSGScopeMockRecorder) TenantID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantID", reflect.TypeOf((*MockNSGScope)(nil).TenantID))
}

// BaseURI mocks base method.
func (m *MockNSGScope) BaseURI() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseURI")
	ret0, _ := ret[0].(string)
	return ret0
}

// BaseURI indicates an expected call of BaseURI.
func (mr *MockNSGScopeMockRecorder) BaseURI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseURI", reflect.TypeOf((*MockNSGScope)(nil).BaseURI))
}

// Authorizer mocks base method.
func (m *MockNSGScope) Authorizer() autorest.Authorizer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorizer")
	ret0, _ := ret[0].(autorest.Authorizer)
	return ret0
}

// Authorizer indicates an expected call of Authorizer.
func (mr *MockNSGScopeMockRecorder) Authorizer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorizer", reflect.TypeOf((*MockNSGScope)(nil).Authorizer))
}

// Info mocks base method.
func (m *MockNSGScope) Info(msg string, keysAndValues ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{msg}
	for _, a := range keysAndValues {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info.
func (mr *MockNSGScopeMockRecorder) Info(msg interface{}, keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{msg}, keysAndValues...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockNSGScope)(nil).Info), varargs...)
}

// Enabled mocks base method.
func (m *MockNSGScope) Enabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Enabled indicates an expected call of Enabled.
func (mr *MockNSGScopeMockRecorder) Enabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enabled", reflect.TypeOf((*MockNSGScope)(nil).Enabled))
}

// Error mocks base method.
func (m *MockNSGScope) Error(err error, msg string, keysAndValues ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{err, msg}
	for _, a := range keysAndValues {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error.
func (mr *MockNSGScopeMockRecorder) Error(err, msg interface{}, keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{err, msg}, keysAndValues...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockNSGScope)(nil).Error), varargs...)
}

// V mocks base method.
func (m *MockNSGScope) V(level int) logr.InfoLogger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V", level)
	ret0, _ := ret[0].(logr.InfoLogger)
	return ret0
}

// V indicates an expected call of V.
func (mr *MockNSGScopeMockRecorder) V(level interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V", reflect.TypeOf((*MockNSGScope)(nil).V), level)
}

// WithValues mocks base method.
func (m *MockNSGScope) WithValues(keysAndValues ...interface{}) logr.Logger {
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
func (mr *MockNSGScopeMockRecorder) WithValues(keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithValues", reflect.TypeOf((*MockNSGScope)(nil).WithValues), keysAndValues...)
}

// WithName mocks base method.
func (m *MockNSGScope) WithName(name string) logr.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithName", name)
	ret0, _ := ret[0].(logr.Logger)
	return ret0
}

// WithName indicates an expected call of WithName.
func (mr *MockNSGScopeMockRecorder) WithName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithName", reflect.TypeOf((*MockNSGScope)(nil).WithName), name)
}

// NSGSpecs mocks base method.
func (m *MockNSGScope) NSGSpecs() []azure.NSGSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NSGSpecs")
	ret0, _ := ret[0].([]azure.NSGSpec)
	return ret0
}

// NSGSpecs indicates an expected call of NSGSpecs.
func (mr *MockNSGScopeMockRecorder) NSGSpecs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NSGSpecs", reflect.TypeOf((*MockNSGScope)(nil).NSGSpecs))
}
