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

// Package mock_virtualnetworks is a generated GoMock package.
package mock_virtualnetworks

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

// MockVNetScope is a mock of VNetScope interface.
type MockVNetScope struct {
	ctrl     *gomock.Controller
	recorder *MockVNetScopeMockRecorder
}

// MockVNetScopeMockRecorder is the mock recorder for MockVNetScope.
type MockVNetScopeMockRecorder struct {
	mock *MockVNetScope
}

// NewMockVNetScope creates a new mock instance.
func NewMockVNetScope(ctrl *gomock.Controller) *MockVNetScope {
	mock := &MockVNetScope{ctrl: ctrl}
	mock.recorder = &MockVNetScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVNetScope) EXPECT() *MockVNetScopeMockRecorder {
	return m.recorder
}

// Info mocks base method.
func (m *MockVNetScope) Info(msg string, keysAndValues ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{msg}
	for _, a := range keysAndValues {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info.
func (mr *MockVNetScopeMockRecorder) Info(msg interface{}, keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{msg}, keysAndValues...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockVNetScope)(nil).Info), varargs...)
}

// Enabled mocks base method.
func (m *MockVNetScope) Enabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Enabled indicates an expected call of Enabled.
func (mr *MockVNetScopeMockRecorder) Enabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enabled", reflect.TypeOf((*MockVNetScope)(nil).Enabled))
}

// Error mocks base method.
func (m *MockVNetScope) Error(err error, msg string, keysAndValues ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{err, msg}
	for _, a := range keysAndValues {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error.
func (mr *MockVNetScopeMockRecorder) Error(err, msg interface{}, keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{err, msg}, keysAndValues...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockVNetScope)(nil).Error), varargs...)
}

// V mocks base method.
func (m *MockVNetScope) V(level int) logr.InfoLogger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V", level)
	ret0, _ := ret[0].(logr.InfoLogger)
	return ret0
}

// V indicates an expected call of V.
func (mr *MockVNetScopeMockRecorder) V(level interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V", reflect.TypeOf((*MockVNetScope)(nil).V), level)
}

// WithValues mocks base method.
func (m *MockVNetScope) WithValues(keysAndValues ...interface{}) logr.Logger {
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
func (mr *MockVNetScopeMockRecorder) WithValues(keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithValues", reflect.TypeOf((*MockVNetScope)(nil).WithValues), keysAndValues...)
}

// WithName mocks base method.
func (m *MockVNetScope) WithName(name string) logr.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithName", name)
	ret0, _ := ret[0].(logr.Logger)
	return ret0
}

// WithName indicates an expected call of WithName.
func (mr *MockVNetScopeMockRecorder) WithName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithName", reflect.TypeOf((*MockVNetScope)(nil).WithName), name)
}

// GetNamespace mocks base method.
func (m *MockVNetScope) GetNamespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetNamespace indicates an expected call of GetNamespace.
func (mr *MockVNetScopeMockRecorder) GetNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockVNetScope)(nil).GetNamespace))
}

// SetNamespace mocks base method.
func (m *MockVNetScope) SetNamespace(namespace string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetNamespace", namespace)
}

// SetNamespace indicates an expected call of SetNamespace.
func (mr *MockVNetScopeMockRecorder) SetNamespace(namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNamespace", reflect.TypeOf((*MockVNetScope)(nil).SetNamespace), namespace)
}

// GetName mocks base method.
func (m *MockVNetScope) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockVNetScopeMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockVNetScope)(nil).GetName))
}

// SetName mocks base method.
func (m *MockVNetScope) SetName(name string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetName", name)
}

// SetName indicates an expected call of SetName.
func (mr *MockVNetScopeMockRecorder) SetName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetName", reflect.TypeOf((*MockVNetScope)(nil).SetName), name)
}

// GetGenerateName mocks base method.
func (m *MockVNetScope) GetGenerateName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGenerateName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetGenerateName indicates an expected call of GetGenerateName.
func (mr *MockVNetScopeMockRecorder) GetGenerateName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenerateName", reflect.TypeOf((*MockVNetScope)(nil).GetGenerateName))
}

// SetGenerateName mocks base method.
func (m *MockVNetScope) SetGenerateName(name string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetGenerateName", name)
}

// SetGenerateName indicates an expected call of SetGenerateName.
func (mr *MockVNetScopeMockRecorder) SetGenerateName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGenerateName", reflect.TypeOf((*MockVNetScope)(nil).SetGenerateName), name)
}

// GetUID mocks base method.
func (m *MockVNetScope) GetUID() types.UID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUID")
	ret0, _ := ret[0].(types.UID)
	return ret0
}

// GetUID indicates an expected call of GetUID.
func (mr *MockVNetScopeMockRecorder) GetUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUID", reflect.TypeOf((*MockVNetScope)(nil).GetUID))
}

// SetUID mocks base method.
func (m *MockVNetScope) SetUID(uid types.UID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUID", uid)
}

// SetUID indicates an expected call of SetUID.
func (mr *MockVNetScopeMockRecorder) SetUID(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUID", reflect.TypeOf((*MockVNetScope)(nil).SetUID), uid)
}

// GetResourceVersion mocks base method.
func (m *MockVNetScope) GetResourceVersion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetResourceVersion indicates an expected call of GetResourceVersion.
func (mr *MockVNetScopeMockRecorder) GetResourceVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceVersion", reflect.TypeOf((*MockVNetScope)(nil).GetResourceVersion))
}

// SetResourceVersion mocks base method.
func (m *MockVNetScope) SetResourceVersion(version string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetResourceVersion", version)
}

// SetResourceVersion indicates an expected call of SetResourceVersion.
func (mr *MockVNetScopeMockRecorder) SetResourceVersion(version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetResourceVersion", reflect.TypeOf((*MockVNetScope)(nil).SetResourceVersion), version)
}

// GetGeneration mocks base method.
func (m *MockVNetScope) GetGeneration() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGeneration")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetGeneration indicates an expected call of GetGeneration.
func (mr *MockVNetScopeMockRecorder) GetGeneration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGeneration", reflect.TypeOf((*MockVNetScope)(nil).GetGeneration))
}

// SetGeneration mocks base method.
func (m *MockVNetScope) SetGeneration(generation int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetGeneration", generation)
}

// SetGeneration indicates an expected call of SetGeneration.
func (mr *MockVNetScopeMockRecorder) SetGeneration(generation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGeneration", reflect.TypeOf((*MockVNetScope)(nil).SetGeneration), generation)
}

// GetSelfLink mocks base method.
func (m *MockVNetScope) GetSelfLink() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSelfLink")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSelfLink indicates an expected call of GetSelfLink.
func (mr *MockVNetScopeMockRecorder) GetSelfLink() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSelfLink", reflect.TypeOf((*MockVNetScope)(nil).GetSelfLink))
}

// SetSelfLink mocks base method.
func (m *MockVNetScope) SetSelfLink(selfLink string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSelfLink", selfLink)
}

// SetSelfLink indicates an expected call of SetSelfLink.
func (mr *MockVNetScopeMockRecorder) SetSelfLink(selfLink interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSelfLink", reflect.TypeOf((*MockVNetScope)(nil).SetSelfLink), selfLink)
}

// GetCreationTimestamp mocks base method.
func (m *MockVNetScope) GetCreationTimestamp() v1.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCreationTimestamp")
	ret0, _ := ret[0].(v1.Time)
	return ret0
}

// GetCreationTimestamp indicates an expected call of GetCreationTimestamp.
func (mr *MockVNetScopeMockRecorder) GetCreationTimestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreationTimestamp", reflect.TypeOf((*MockVNetScope)(nil).GetCreationTimestamp))
}

// SetCreationTimestamp mocks base method.
func (m *MockVNetScope) SetCreationTimestamp(timestamp v1.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCreationTimestamp", timestamp)
}

// SetCreationTimestamp indicates an expected call of SetCreationTimestamp.
func (mr *MockVNetScopeMockRecorder) SetCreationTimestamp(timestamp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCreationTimestamp", reflect.TypeOf((*MockVNetScope)(nil).SetCreationTimestamp), timestamp)
}

// GetDeletionTimestamp mocks base method.
func (m *MockVNetScope) GetDeletionTimestamp() *v1.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeletionTimestamp")
	ret0, _ := ret[0].(*v1.Time)
	return ret0
}

// GetDeletionTimestamp indicates an expected call of GetDeletionTimestamp.
func (mr *MockVNetScopeMockRecorder) GetDeletionTimestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeletionTimestamp", reflect.TypeOf((*MockVNetScope)(nil).GetDeletionTimestamp))
}

// SetDeletionTimestamp mocks base method.
func (m *MockVNetScope) SetDeletionTimestamp(timestamp *v1.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDeletionTimestamp", timestamp)
}

// SetDeletionTimestamp indicates an expected call of SetDeletionTimestamp.
func (mr *MockVNetScopeMockRecorder) SetDeletionTimestamp(timestamp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeletionTimestamp", reflect.TypeOf((*MockVNetScope)(nil).SetDeletionTimestamp), timestamp)
}

// GetDeletionGracePeriodSeconds mocks base method.
func (m *MockVNetScope) GetDeletionGracePeriodSeconds() *int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeletionGracePeriodSeconds")
	ret0, _ := ret[0].(*int64)
	return ret0
}

// GetDeletionGracePeriodSeconds indicates an expected call of GetDeletionGracePeriodSeconds.
func (mr *MockVNetScopeMockRecorder) GetDeletionGracePeriodSeconds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeletionGracePeriodSeconds", reflect.TypeOf((*MockVNetScope)(nil).GetDeletionGracePeriodSeconds))
}

// SetDeletionGracePeriodSeconds mocks base method.
func (m *MockVNetScope) SetDeletionGracePeriodSeconds(arg0 *int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDeletionGracePeriodSeconds", arg0)
}

// SetDeletionGracePeriodSeconds indicates an expected call of SetDeletionGracePeriodSeconds.
func (mr *MockVNetScopeMockRecorder) SetDeletionGracePeriodSeconds(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeletionGracePeriodSeconds", reflect.TypeOf((*MockVNetScope)(nil).SetDeletionGracePeriodSeconds), arg0)
}

// GetLabels mocks base method.
func (m *MockVNetScope) GetLabels() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLabels")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetLabels indicates an expected call of GetLabels.
func (mr *MockVNetScopeMockRecorder) GetLabels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLabels", reflect.TypeOf((*MockVNetScope)(nil).GetLabels))
}

// SetLabels mocks base method.
func (m *MockVNetScope) SetLabels(labels map[string]string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLabels", labels)
}

// SetLabels indicates an expected call of SetLabels.
func (mr *MockVNetScopeMockRecorder) SetLabels(labels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLabels", reflect.TypeOf((*MockVNetScope)(nil).SetLabels), labels)
}

// GetAnnotations mocks base method.
func (m *MockVNetScope) GetAnnotations() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAnnotations")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetAnnotations indicates an expected call of GetAnnotations.
func (mr *MockVNetScopeMockRecorder) GetAnnotations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnnotations", reflect.TypeOf((*MockVNetScope)(nil).GetAnnotations))
}

// SetAnnotations mocks base method.
func (m *MockVNetScope) SetAnnotations(annotations map[string]string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAnnotations", annotations)
}

// SetAnnotations indicates an expected call of SetAnnotations.
func (mr *MockVNetScopeMockRecorder) SetAnnotations(annotations interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAnnotations", reflect.TypeOf((*MockVNetScope)(nil).SetAnnotations), annotations)
}

// GetFinalizers mocks base method.
func (m *MockVNetScope) GetFinalizers() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFinalizers")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetFinalizers indicates an expected call of GetFinalizers.
func (mr *MockVNetScopeMockRecorder) GetFinalizers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFinalizers", reflect.TypeOf((*MockVNetScope)(nil).GetFinalizers))
}

// SetFinalizers mocks base method.
func (m *MockVNetScope) SetFinalizers(finalizers []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetFinalizers", finalizers)
}

// SetFinalizers indicates an expected call of SetFinalizers.
func (mr *MockVNetScopeMockRecorder) SetFinalizers(finalizers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFinalizers", reflect.TypeOf((*MockVNetScope)(nil).SetFinalizers), finalizers)
}

// GetOwnerReferences mocks base method.
func (m *MockVNetScope) GetOwnerReferences() []v1.OwnerReference {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOwnerReferences")
	ret0, _ := ret[0].([]v1.OwnerReference)
	return ret0
}

// GetOwnerReferences indicates an expected call of GetOwnerReferences.
func (mr *MockVNetScopeMockRecorder) GetOwnerReferences() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOwnerReferences", reflect.TypeOf((*MockVNetScope)(nil).GetOwnerReferences))
}

// SetOwnerReferences mocks base method.
func (m *MockVNetScope) SetOwnerReferences(arg0 []v1.OwnerReference) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetOwnerReferences", arg0)
}

// SetOwnerReferences indicates an expected call of SetOwnerReferences.
func (mr *MockVNetScopeMockRecorder) SetOwnerReferences(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOwnerReferences", reflect.TypeOf((*MockVNetScope)(nil).SetOwnerReferences), arg0)
}

// GetClusterName mocks base method.
func (m *MockVNetScope) GetClusterName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetClusterName indicates an expected call of GetClusterName.
func (mr *MockVNetScopeMockRecorder) GetClusterName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterName", reflect.TypeOf((*MockVNetScope)(nil).GetClusterName))
}

// SetClusterName mocks base method.
func (m *MockVNetScope) SetClusterName(clusterName string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetClusterName", clusterName)
}

// SetClusterName indicates an expected call of SetClusterName.
func (mr *MockVNetScopeMockRecorder) SetClusterName(clusterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClusterName", reflect.TypeOf((*MockVNetScope)(nil).SetClusterName), clusterName)
}

// GetManagedFields mocks base method.
func (m *MockVNetScope) GetManagedFields() []v1.ManagedFieldsEntry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetManagedFields")
	ret0, _ := ret[0].([]v1.ManagedFieldsEntry)
	return ret0
}

// GetManagedFields indicates an expected call of GetManagedFields.
func (mr *MockVNetScopeMockRecorder) GetManagedFields() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManagedFields", reflect.TypeOf((*MockVNetScope)(nil).GetManagedFields))
}

// SetManagedFields mocks base method.
func (m *MockVNetScope) SetManagedFields(managedFields []v1.ManagedFieldsEntry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetManagedFields", managedFields)
}

// SetManagedFields indicates an expected call of SetManagedFields.
func (mr *MockVNetScopeMockRecorder) SetManagedFields(managedFields interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetManagedFields", reflect.TypeOf((*MockVNetScope)(nil).SetManagedFields), managedFields)
}

// GetObjectKind mocks base method.
func (m *MockVNetScope) GetObjectKind() schema.ObjectKind {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectKind")
	ret0, _ := ret[0].(schema.ObjectKind)
	return ret0
}

// GetObjectKind indicates an expected call of GetObjectKind.
func (mr *MockVNetScopeMockRecorder) GetObjectKind() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectKind", reflect.TypeOf((*MockVNetScope)(nil).GetObjectKind))
}

// DeepCopyObject mocks base method.
func (m *MockVNetScope) DeepCopyObject() runtime.Object {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeepCopyObject")
	ret0, _ := ret[0].(runtime.Object)
	return ret0
}

// DeepCopyObject indicates an expected call of DeepCopyObject.
func (mr *MockVNetScopeMockRecorder) DeepCopyObject() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeepCopyObject", reflect.TypeOf((*MockVNetScope)(nil).DeepCopyObject))
}

// LoadBalancerName mocks base method.
func (m *MockVNetScope) LoadBalancerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadBalancerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// LoadBalancerName indicates an expected call of LoadBalancerName.
func (mr *MockVNetScopeMockRecorder) LoadBalancerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadBalancerName", reflect.TypeOf((*MockVNetScope)(nil).LoadBalancerName))
}

// Network mocks base method.
func (m *MockVNetScope) Network() *v1alpha3.Network {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Network")
	ret0, _ := ret[0].(*v1alpha3.Network)
	return ret0
}

// Network indicates an expected call of Network.
func (mr *MockVNetScopeMockRecorder) Network() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Network", reflect.TypeOf((*MockVNetScope)(nil).Network))
}

// Vnet mocks base method.
func (m *MockVNetScope) Vnet() *v1alpha3.VnetSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Vnet")
	ret0, _ := ret[0].(*v1alpha3.VnetSpec)
	return ret0
}

// Vnet indicates an expected call of Vnet.
func (mr *MockVNetScopeMockRecorder) Vnet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Vnet", reflect.TypeOf((*MockVNetScope)(nil).Vnet))
}

// IsVnetManaged mocks base method.
func (m *MockVNetScope) IsVnetManaged() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsVnetManaged")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsVnetManaged indicates an expected call of IsVnetManaged.
func (mr *MockVNetScopeMockRecorder) IsVnetManaged() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsVnetManaged", reflect.TypeOf((*MockVNetScope)(nil).IsVnetManaged))
}

// Subnets mocks base method.
func (m *MockVNetScope) Subnets() v1alpha3.Subnets {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subnets")
	ret0, _ := ret[0].(v1alpha3.Subnets)
	return ret0
}

// Subnets indicates an expected call of Subnets.
func (mr *MockVNetScopeMockRecorder) Subnets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subnets", reflect.TypeOf((*MockVNetScope)(nil).Subnets))
}

// NodeSubnet mocks base method.
func (m *MockVNetScope) NodeSubnet() *v1alpha3.SubnetSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeSubnet")
	ret0, _ := ret[0].(*v1alpha3.SubnetSpec)
	return ret0
}

// NodeSubnet indicates an expected call of NodeSubnet.
func (mr *MockVNetScopeMockRecorder) NodeSubnet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeSubnet", reflect.TypeOf((*MockVNetScope)(nil).NodeSubnet))
}

// ControlPlaneSubnet mocks base method.
func (m *MockVNetScope) ControlPlaneSubnet() *v1alpha3.SubnetSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControlPlaneSubnet")
	ret0, _ := ret[0].(*v1alpha3.SubnetSpec)
	return ret0
}

// ControlPlaneSubnet indicates an expected call of ControlPlaneSubnet.
func (mr *MockVNetScopeMockRecorder) ControlPlaneSubnet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControlPlaneSubnet", reflect.TypeOf((*MockVNetScope)(nil).ControlPlaneSubnet))
}

// RouteTable mocks base method.
func (m *MockVNetScope) RouteTable() *v1alpha3.RouteTable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RouteTable")
	ret0, _ := ret[0].(*v1alpha3.RouteTable)
	return ret0
}

// RouteTable indicates an expected call of RouteTable.
func (mr *MockVNetScopeMockRecorder) RouteTable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteTable", reflect.TypeOf((*MockVNetScope)(nil).RouteTable))
}

// SubscriptionID mocks base method.
func (m *MockVNetScope) SubscriptionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SubscriptionID indicates an expected call of SubscriptionID.
func (mr *MockVNetScopeMockRecorder) SubscriptionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionID", reflect.TypeOf((*MockVNetScope)(nil).SubscriptionID))
}

// ResourceGroup mocks base method.
func (m *MockVNetScope) ResourceGroup() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceGroup")
	ret0, _ := ret[0].(string)
	return ret0
}

// ResourceGroup indicates an expected call of ResourceGroup.
func (mr *MockVNetScopeMockRecorder) ResourceGroup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceGroup", reflect.TypeOf((*MockVNetScope)(nil).ResourceGroup))
}

// ClusterName mocks base method.
func (m *MockVNetScope) ClusterName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClusterName indicates an expected call of ClusterName.
func (mr *MockVNetScopeMockRecorder) ClusterName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterName", reflect.TypeOf((*MockVNetScope)(nil).ClusterName))
}

// Location mocks base method.
func (m *MockVNetScope) Location() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Location")
	ret0, _ := ret[0].(string)
	return ret0
}

// Location indicates an expected call of Location.
func (mr *MockVNetScopeMockRecorder) Location() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Location", reflect.TypeOf((*MockVNetScope)(nil).Location))
}

// SetFailureDomain mocks base method.
func (m *MockVNetScope) SetFailureDomain(id string, spec v1alpha30.FailureDomainSpec) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetFailureDomain", id, spec)
}

// SetFailureDomain indicates an expected call of SetFailureDomain.
func (mr *MockVNetScopeMockRecorder) SetFailureDomain(id, spec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFailureDomain", reflect.TypeOf((*MockVNetScope)(nil).SetFailureDomain), id, spec)
}

// AdditionalTags mocks base method.
func (m *MockVNetScope) AdditionalTags() v1alpha3.Tags {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdditionalTags")
	ret0, _ := ret[0].(v1alpha3.Tags)
	return ret0
}

// AdditionalTags indicates an expected call of AdditionalTags.
func (mr *MockVNetScopeMockRecorder) AdditionalTags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdditionalTags", reflect.TypeOf((*MockVNetScope)(nil).AdditionalTags))
}

// ClientID mocks base method.
func (m *MockVNetScope) ClientID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientID indicates an expected call of ClientID.
func (mr *MockVNetScopeMockRecorder) ClientID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientID", reflect.TypeOf((*MockVNetScope)(nil).ClientID))
}

// ClientSecret mocks base method.
func (m *MockVNetScope) ClientSecret() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientSecret")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientSecret indicates an expected call of ClientSecret.
func (mr *MockVNetScopeMockRecorder) ClientSecret() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientSecret", reflect.TypeOf((*MockVNetScope)(nil).ClientSecret))
}

// CloudEnvironment mocks base method.
func (m *MockVNetScope) CloudEnvironment() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudEnvironment")
	ret0, _ := ret[0].(string)
	return ret0
}

// CloudEnvironment indicates an expected call of CloudEnvironment.
func (mr *MockVNetScopeMockRecorder) CloudEnvironment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudEnvironment", reflect.TypeOf((*MockVNetScope)(nil).CloudEnvironment))
}

// TenantID mocks base method.
func (m *MockVNetScope) TenantID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantID")
	ret0, _ := ret[0].(string)
	return ret0
}

// TenantID indicates an expected call of TenantID.
func (mr *MockVNetScopeMockRecorder) TenantID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantID", reflect.TypeOf((*MockVNetScope)(nil).TenantID))
}

// BaseURI mocks base method.
func (m *MockVNetScope) BaseURI() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseURI")
	ret0, _ := ret[0].(string)
	return ret0
}

// BaseURI indicates an expected call of BaseURI.
func (mr *MockVNetScopeMockRecorder) BaseURI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseURI", reflect.TypeOf((*MockVNetScope)(nil).BaseURI))
}

// Authorizer mocks base method.
func (m *MockVNetScope) Authorizer() autorest.Authorizer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorizer")
	ret0, _ := ret[0].(autorest.Authorizer)
	return ret0
}

// Authorizer indicates an expected call of Authorizer.
func (mr *MockVNetScopeMockRecorder) Authorizer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorizer", reflect.TypeOf((*MockVNetScope)(nil).Authorizer))
}

// VNetSpecs mocks base method.
func (m *MockVNetScope) VNetSpecs() []azure.VNetSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VNetSpecs")
	ret0, _ := ret[0].([]azure.VNetSpec)
	return ret0
}

// VNetSpecs indicates an expected call of VNetSpecs.
func (mr *MockVNetScopeMockRecorder) VNetSpecs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VNetSpecs", reflect.TypeOf((*MockVNetScope)(nil).VNetSpecs))
}
