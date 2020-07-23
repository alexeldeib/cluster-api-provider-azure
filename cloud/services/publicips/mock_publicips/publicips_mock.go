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

// Package mock_publicips is a generated GoMock package.
package mock_publicips

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

// MockPublicIPScope is a mock of PublicIPScope interface.
type MockPublicIPScope struct {
	ctrl     *gomock.Controller
	recorder *MockPublicIPScopeMockRecorder
}

// MockPublicIPScopeMockRecorder is the mock recorder for MockPublicIPScope.
type MockPublicIPScopeMockRecorder struct {
	mock *MockPublicIPScope
}

// NewMockPublicIPScope creates a new mock instance.
func NewMockPublicIPScope(ctrl *gomock.Controller) *MockPublicIPScope {
	mock := &MockPublicIPScope{ctrl: ctrl}
	mock.recorder = &MockPublicIPScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPublicIPScope) EXPECT() *MockPublicIPScopeMockRecorder {
	return m.recorder
}

// Info mocks base method.
func (m *MockPublicIPScope) Info(msg string, keysAndValues ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{msg}
	for _, a := range keysAndValues {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info.
func (mr *MockPublicIPScopeMockRecorder) Info(msg interface{}, keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{msg}, keysAndValues...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockPublicIPScope)(nil).Info), varargs...)
}

// Enabled mocks base method.
func (m *MockPublicIPScope) Enabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Enabled indicates an expected call of Enabled.
func (mr *MockPublicIPScopeMockRecorder) Enabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enabled", reflect.TypeOf((*MockPublicIPScope)(nil).Enabled))
}

// Error mocks base method.
func (m *MockPublicIPScope) Error(err error, msg string, keysAndValues ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{err, msg}
	for _, a := range keysAndValues {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error.
func (mr *MockPublicIPScopeMockRecorder) Error(err, msg interface{}, keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{err, msg}, keysAndValues...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockPublicIPScope)(nil).Error), varargs...)
}

// V mocks base method.
func (m *MockPublicIPScope) V(level int) logr.InfoLogger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V", level)
	ret0, _ := ret[0].(logr.InfoLogger)
	return ret0
}

// V indicates an expected call of V.
func (mr *MockPublicIPScopeMockRecorder) V(level interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V", reflect.TypeOf((*MockPublicIPScope)(nil).V), level)
}

// WithValues mocks base method.
func (m *MockPublicIPScope) WithValues(keysAndValues ...interface{}) logr.Logger {
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
func (mr *MockPublicIPScopeMockRecorder) WithValues(keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithValues", reflect.TypeOf((*MockPublicIPScope)(nil).WithValues), keysAndValues...)
}

// WithName mocks base method.
func (m *MockPublicIPScope) WithName(name string) logr.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithName", name)
	ret0, _ := ret[0].(logr.Logger)
	return ret0
}

// WithName indicates an expected call of WithName.
func (mr *MockPublicIPScopeMockRecorder) WithName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithName", reflect.TypeOf((*MockPublicIPScope)(nil).WithName), name)
}

// BaseURI mocks base method.
func (m *MockPublicIPScope) BaseURI() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseURI")
	ret0, _ := ret[0].(string)
	return ret0
}

// BaseURI indicates an expected call of BaseURI.
func (mr *MockPublicIPScopeMockRecorder) BaseURI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseURI", reflect.TypeOf((*MockPublicIPScope)(nil).BaseURI))
}

// Authorizer mocks base method.
func (m *MockPublicIPScope) Authorizer() autorest.Authorizer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorizer")
	ret0, _ := ret[0].(autorest.Authorizer)
	return ret0
}

// Authorizer indicates an expected call of Authorizer.
func (mr *MockPublicIPScopeMockRecorder) Authorizer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorizer", reflect.TypeOf((*MockPublicIPScope)(nil).Authorizer))
}

// GetNamespace mocks base method.
func (m *MockPublicIPScope) GetNamespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetNamespace indicates an expected call of GetNamespace.
func (mr *MockPublicIPScopeMockRecorder) GetNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockPublicIPScope)(nil).GetNamespace))
}

// SetNamespace mocks base method.
func (m *MockPublicIPScope) SetNamespace(namespace string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetNamespace", namespace)
}

// SetNamespace indicates an expected call of SetNamespace.
func (mr *MockPublicIPScopeMockRecorder) SetNamespace(namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNamespace", reflect.TypeOf((*MockPublicIPScope)(nil).SetNamespace), namespace)
}

// GetName mocks base method.
func (m *MockPublicIPScope) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockPublicIPScopeMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockPublicIPScope)(nil).GetName))
}

// SetName mocks base method.
func (m *MockPublicIPScope) SetName(name string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetName", name)
}

// SetName indicates an expected call of SetName.
func (mr *MockPublicIPScopeMockRecorder) SetName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetName", reflect.TypeOf((*MockPublicIPScope)(nil).SetName), name)
}

// GetGenerateName mocks base method.
func (m *MockPublicIPScope) GetGenerateName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGenerateName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetGenerateName indicates an expected call of GetGenerateName.
func (mr *MockPublicIPScopeMockRecorder) GetGenerateName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenerateName", reflect.TypeOf((*MockPublicIPScope)(nil).GetGenerateName))
}

// SetGenerateName mocks base method.
func (m *MockPublicIPScope) SetGenerateName(name string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetGenerateName", name)
}

// SetGenerateName indicates an expected call of SetGenerateName.
func (mr *MockPublicIPScopeMockRecorder) SetGenerateName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGenerateName", reflect.TypeOf((*MockPublicIPScope)(nil).SetGenerateName), name)
}

// GetUID mocks base method.
func (m *MockPublicIPScope) GetUID() types.UID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUID")
	ret0, _ := ret[0].(types.UID)
	return ret0
}

// GetUID indicates an expected call of GetUID.
func (mr *MockPublicIPScopeMockRecorder) GetUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUID", reflect.TypeOf((*MockPublicIPScope)(nil).GetUID))
}

// SetUID mocks base method.
func (m *MockPublicIPScope) SetUID(uid types.UID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUID", uid)
}

// SetUID indicates an expected call of SetUID.
func (mr *MockPublicIPScopeMockRecorder) SetUID(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUID", reflect.TypeOf((*MockPublicIPScope)(nil).SetUID), uid)
}

// GetResourceVersion mocks base method.
func (m *MockPublicIPScope) GetResourceVersion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetResourceVersion indicates an expected call of GetResourceVersion.
func (mr *MockPublicIPScopeMockRecorder) GetResourceVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceVersion", reflect.TypeOf((*MockPublicIPScope)(nil).GetResourceVersion))
}

// SetResourceVersion mocks base method.
func (m *MockPublicIPScope) SetResourceVersion(version string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetResourceVersion", version)
}

// SetResourceVersion indicates an expected call of SetResourceVersion.
func (mr *MockPublicIPScopeMockRecorder) SetResourceVersion(version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetResourceVersion", reflect.TypeOf((*MockPublicIPScope)(nil).SetResourceVersion), version)
}

// GetGeneration mocks base method.
func (m *MockPublicIPScope) GetGeneration() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGeneration")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetGeneration indicates an expected call of GetGeneration.
func (mr *MockPublicIPScopeMockRecorder) GetGeneration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGeneration", reflect.TypeOf((*MockPublicIPScope)(nil).GetGeneration))
}

// SetGeneration mocks base method.
func (m *MockPublicIPScope) SetGeneration(generation int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetGeneration", generation)
}

// SetGeneration indicates an expected call of SetGeneration.
func (mr *MockPublicIPScopeMockRecorder) SetGeneration(generation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGeneration", reflect.TypeOf((*MockPublicIPScope)(nil).SetGeneration), generation)
}

// GetSelfLink mocks base method.
func (m *MockPublicIPScope) GetSelfLink() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSelfLink")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSelfLink indicates an expected call of GetSelfLink.
func (mr *MockPublicIPScopeMockRecorder) GetSelfLink() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSelfLink", reflect.TypeOf((*MockPublicIPScope)(nil).GetSelfLink))
}

// SetSelfLink mocks base method.
func (m *MockPublicIPScope) SetSelfLink(selfLink string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSelfLink", selfLink)
}

// SetSelfLink indicates an expected call of SetSelfLink.
func (mr *MockPublicIPScopeMockRecorder) SetSelfLink(selfLink interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSelfLink", reflect.TypeOf((*MockPublicIPScope)(nil).SetSelfLink), selfLink)
}

// GetCreationTimestamp mocks base method.
func (m *MockPublicIPScope) GetCreationTimestamp() v1.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCreationTimestamp")
	ret0, _ := ret[0].(v1.Time)
	return ret0
}

// GetCreationTimestamp indicates an expected call of GetCreationTimestamp.
func (mr *MockPublicIPScopeMockRecorder) GetCreationTimestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreationTimestamp", reflect.TypeOf((*MockPublicIPScope)(nil).GetCreationTimestamp))
}

// SetCreationTimestamp mocks base method.
func (m *MockPublicIPScope) SetCreationTimestamp(timestamp v1.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCreationTimestamp", timestamp)
}

// SetCreationTimestamp indicates an expected call of SetCreationTimestamp.
func (mr *MockPublicIPScopeMockRecorder) SetCreationTimestamp(timestamp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCreationTimestamp", reflect.TypeOf((*MockPublicIPScope)(nil).SetCreationTimestamp), timestamp)
}

// GetDeletionTimestamp mocks base method.
func (m *MockPublicIPScope) GetDeletionTimestamp() *v1.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeletionTimestamp")
	ret0, _ := ret[0].(*v1.Time)
	return ret0
}

// GetDeletionTimestamp indicates an expected call of GetDeletionTimestamp.
func (mr *MockPublicIPScopeMockRecorder) GetDeletionTimestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeletionTimestamp", reflect.TypeOf((*MockPublicIPScope)(nil).GetDeletionTimestamp))
}

// SetDeletionTimestamp mocks base method.
func (m *MockPublicIPScope) SetDeletionTimestamp(timestamp *v1.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDeletionTimestamp", timestamp)
}

// SetDeletionTimestamp indicates an expected call of SetDeletionTimestamp.
func (mr *MockPublicIPScopeMockRecorder) SetDeletionTimestamp(timestamp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeletionTimestamp", reflect.TypeOf((*MockPublicIPScope)(nil).SetDeletionTimestamp), timestamp)
}

// GetDeletionGracePeriodSeconds mocks base method.
func (m *MockPublicIPScope) GetDeletionGracePeriodSeconds() *int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeletionGracePeriodSeconds")
	ret0, _ := ret[0].(*int64)
	return ret0
}

// GetDeletionGracePeriodSeconds indicates an expected call of GetDeletionGracePeriodSeconds.
func (mr *MockPublicIPScopeMockRecorder) GetDeletionGracePeriodSeconds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeletionGracePeriodSeconds", reflect.TypeOf((*MockPublicIPScope)(nil).GetDeletionGracePeriodSeconds))
}

// SetDeletionGracePeriodSeconds mocks base method.
func (m *MockPublicIPScope) SetDeletionGracePeriodSeconds(arg0 *int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDeletionGracePeriodSeconds", arg0)
}

// SetDeletionGracePeriodSeconds indicates an expected call of SetDeletionGracePeriodSeconds.
func (mr *MockPublicIPScopeMockRecorder) SetDeletionGracePeriodSeconds(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeletionGracePeriodSeconds", reflect.TypeOf((*MockPublicIPScope)(nil).SetDeletionGracePeriodSeconds), arg0)
}

// GetLabels mocks base method.
func (m *MockPublicIPScope) GetLabels() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLabels")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetLabels indicates an expected call of GetLabels.
func (mr *MockPublicIPScopeMockRecorder) GetLabels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLabels", reflect.TypeOf((*MockPublicIPScope)(nil).GetLabels))
}

// SetLabels mocks base method.
func (m *MockPublicIPScope) SetLabels(labels map[string]string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLabels", labels)
}

// SetLabels indicates an expected call of SetLabels.
func (mr *MockPublicIPScopeMockRecorder) SetLabels(labels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLabels", reflect.TypeOf((*MockPublicIPScope)(nil).SetLabels), labels)
}

// GetAnnotations mocks base method.
func (m *MockPublicIPScope) GetAnnotations() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAnnotations")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetAnnotations indicates an expected call of GetAnnotations.
func (mr *MockPublicIPScopeMockRecorder) GetAnnotations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnnotations", reflect.TypeOf((*MockPublicIPScope)(nil).GetAnnotations))
}

// SetAnnotations mocks base method.
func (m *MockPublicIPScope) SetAnnotations(annotations map[string]string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAnnotations", annotations)
}

// SetAnnotations indicates an expected call of SetAnnotations.
func (mr *MockPublicIPScopeMockRecorder) SetAnnotations(annotations interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAnnotations", reflect.TypeOf((*MockPublicIPScope)(nil).SetAnnotations), annotations)
}

// GetFinalizers mocks base method.
func (m *MockPublicIPScope) GetFinalizers() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFinalizers")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetFinalizers indicates an expected call of GetFinalizers.
func (mr *MockPublicIPScopeMockRecorder) GetFinalizers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFinalizers", reflect.TypeOf((*MockPublicIPScope)(nil).GetFinalizers))
}

// SetFinalizers mocks base method.
func (m *MockPublicIPScope) SetFinalizers(finalizers []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetFinalizers", finalizers)
}

// SetFinalizers indicates an expected call of SetFinalizers.
func (mr *MockPublicIPScopeMockRecorder) SetFinalizers(finalizers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFinalizers", reflect.TypeOf((*MockPublicIPScope)(nil).SetFinalizers), finalizers)
}

// GetOwnerReferences mocks base method.
func (m *MockPublicIPScope) GetOwnerReferences() []v1.OwnerReference {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOwnerReferences")
	ret0, _ := ret[0].([]v1.OwnerReference)
	return ret0
}

// GetOwnerReferences indicates an expected call of GetOwnerReferences.
func (mr *MockPublicIPScopeMockRecorder) GetOwnerReferences() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOwnerReferences", reflect.TypeOf((*MockPublicIPScope)(nil).GetOwnerReferences))
}

// SetOwnerReferences mocks base method.
func (m *MockPublicIPScope) SetOwnerReferences(arg0 []v1.OwnerReference) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetOwnerReferences", arg0)
}

// SetOwnerReferences indicates an expected call of SetOwnerReferences.
func (mr *MockPublicIPScopeMockRecorder) SetOwnerReferences(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOwnerReferences", reflect.TypeOf((*MockPublicIPScope)(nil).SetOwnerReferences), arg0)
}

// GetClusterName mocks base method.
func (m *MockPublicIPScope) GetClusterName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetClusterName indicates an expected call of GetClusterName.
func (mr *MockPublicIPScopeMockRecorder) GetClusterName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterName", reflect.TypeOf((*MockPublicIPScope)(nil).GetClusterName))
}

// SetClusterName mocks base method.
func (m *MockPublicIPScope) SetClusterName(clusterName string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetClusterName", clusterName)
}

// SetClusterName indicates an expected call of SetClusterName.
func (mr *MockPublicIPScopeMockRecorder) SetClusterName(clusterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClusterName", reflect.TypeOf((*MockPublicIPScope)(nil).SetClusterName), clusterName)
}

// GetManagedFields mocks base method.
func (m *MockPublicIPScope) GetManagedFields() []v1.ManagedFieldsEntry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetManagedFields")
	ret0, _ := ret[0].([]v1.ManagedFieldsEntry)
	return ret0
}

// GetManagedFields indicates an expected call of GetManagedFields.
func (mr *MockPublicIPScopeMockRecorder) GetManagedFields() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManagedFields", reflect.TypeOf((*MockPublicIPScope)(nil).GetManagedFields))
}

// SetManagedFields mocks base method.
func (m *MockPublicIPScope) SetManagedFields(managedFields []v1.ManagedFieldsEntry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetManagedFields", managedFields)
}

// SetManagedFields indicates an expected call of SetManagedFields.
func (mr *MockPublicIPScopeMockRecorder) SetManagedFields(managedFields interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetManagedFields", reflect.TypeOf((*MockPublicIPScope)(nil).SetManagedFields), managedFields)
}

// GetObjectKind mocks base method.
func (m *MockPublicIPScope) GetObjectKind() schema.ObjectKind {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectKind")
	ret0, _ := ret[0].(schema.ObjectKind)
	return ret0
}

// GetObjectKind indicates an expected call of GetObjectKind.
func (mr *MockPublicIPScopeMockRecorder) GetObjectKind() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectKind", reflect.TypeOf((*MockPublicIPScope)(nil).GetObjectKind))
}

// DeepCopyObject mocks base method.
func (m *MockPublicIPScope) DeepCopyObject() runtime.Object {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeepCopyObject")
	ret0, _ := ret[0].(runtime.Object)
	return ret0
}

// DeepCopyObject indicates an expected call of DeepCopyObject.
func (mr *MockPublicIPScopeMockRecorder) DeepCopyObject() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeepCopyObject", reflect.TypeOf((*MockPublicIPScope)(nil).DeepCopyObject))
}

// LoadBalancerName mocks base method.
func (m *MockPublicIPScope) LoadBalancerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadBalancerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// LoadBalancerName indicates an expected call of LoadBalancerName.
func (mr *MockPublicIPScopeMockRecorder) LoadBalancerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadBalancerName", reflect.TypeOf((*MockPublicIPScope)(nil).LoadBalancerName))
}

// Network mocks base method.
func (m *MockPublicIPScope) Network() *v1alpha3.Network {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Network")
	ret0, _ := ret[0].(*v1alpha3.Network)
	return ret0
}

// Network indicates an expected call of Network.
func (mr *MockPublicIPScopeMockRecorder) Network() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Network", reflect.TypeOf((*MockPublicIPScope)(nil).Network))
}

// Vnet mocks base method.
func (m *MockPublicIPScope) Vnet() *v1alpha3.VnetSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Vnet")
	ret0, _ := ret[0].(*v1alpha3.VnetSpec)
	return ret0
}

// Vnet indicates an expected call of Vnet.
func (mr *MockPublicIPScopeMockRecorder) Vnet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Vnet", reflect.TypeOf((*MockPublicIPScope)(nil).Vnet))
}

// IsVnetManaged mocks base method.
func (m *MockPublicIPScope) IsVnetManaged() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsVnetManaged")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsVnetManaged indicates an expected call of IsVnetManaged.
func (mr *MockPublicIPScopeMockRecorder) IsVnetManaged() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsVnetManaged", reflect.TypeOf((*MockPublicIPScope)(nil).IsVnetManaged))
}

// Subnets mocks base method.
func (m *MockPublicIPScope) Subnets() v1alpha3.Subnets {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subnets")
	ret0, _ := ret[0].(v1alpha3.Subnets)
	return ret0
}

// Subnets indicates an expected call of Subnets.
func (mr *MockPublicIPScopeMockRecorder) Subnets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subnets", reflect.TypeOf((*MockPublicIPScope)(nil).Subnets))
}

// NodeSubnet mocks base method.
func (m *MockPublicIPScope) NodeSubnet() *v1alpha3.SubnetSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeSubnet")
	ret0, _ := ret[0].(*v1alpha3.SubnetSpec)
	return ret0
}

// NodeSubnet indicates an expected call of NodeSubnet.
func (mr *MockPublicIPScopeMockRecorder) NodeSubnet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeSubnet", reflect.TypeOf((*MockPublicIPScope)(nil).NodeSubnet))
}

// ControlPlaneSubnet mocks base method.
func (m *MockPublicIPScope) ControlPlaneSubnet() *v1alpha3.SubnetSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControlPlaneSubnet")
	ret0, _ := ret[0].(*v1alpha3.SubnetSpec)
	return ret0
}

// ControlPlaneSubnet indicates an expected call of ControlPlaneSubnet.
func (mr *MockPublicIPScopeMockRecorder) ControlPlaneSubnet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControlPlaneSubnet", reflect.TypeOf((*MockPublicIPScope)(nil).ControlPlaneSubnet))
}

// RouteTable mocks base method.
func (m *MockPublicIPScope) RouteTable() *v1alpha3.RouteTable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RouteTable")
	ret0, _ := ret[0].(*v1alpha3.RouteTable)
	return ret0
}

// RouteTable indicates an expected call of RouteTable.
func (mr *MockPublicIPScopeMockRecorder) RouteTable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteTable", reflect.TypeOf((*MockPublicIPScope)(nil).RouteTable))
}

// SubscriptionID mocks base method.
func (m *MockPublicIPScope) SubscriptionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SubscriptionID indicates an expected call of SubscriptionID.
func (mr *MockPublicIPScopeMockRecorder) SubscriptionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionID", reflect.TypeOf((*MockPublicIPScope)(nil).SubscriptionID))
}

// ResourceGroup mocks base method.
func (m *MockPublicIPScope) ResourceGroup() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceGroup")
	ret0, _ := ret[0].(string)
	return ret0
}

// ResourceGroup indicates an expected call of ResourceGroup.
func (mr *MockPublicIPScopeMockRecorder) ResourceGroup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceGroup", reflect.TypeOf((*MockPublicIPScope)(nil).ResourceGroup))
}

// ClusterName mocks base method.
func (m *MockPublicIPScope) ClusterName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClusterName indicates an expected call of ClusterName.
func (mr *MockPublicIPScopeMockRecorder) ClusterName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterName", reflect.TypeOf((*MockPublicIPScope)(nil).ClusterName))
}

// Location mocks base method.
func (m *MockPublicIPScope) Location() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Location")
	ret0, _ := ret[0].(string)
	return ret0
}

// Location indicates an expected call of Location.
func (mr *MockPublicIPScopeMockRecorder) Location() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Location", reflect.TypeOf((*MockPublicIPScope)(nil).Location))
}

// SetFailureDomain mocks base method.
func (m *MockPublicIPScope) SetFailureDomain(id string, spec v1alpha30.FailureDomainSpec) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetFailureDomain", id, spec)
}

// SetFailureDomain indicates an expected call of SetFailureDomain.
func (mr *MockPublicIPScopeMockRecorder) SetFailureDomain(id, spec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFailureDomain", reflect.TypeOf((*MockPublicIPScope)(nil).SetFailureDomain), id, spec)
}

// AdditionalTags mocks base method.
func (m *MockPublicIPScope) AdditionalTags() v1alpha3.Tags {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdditionalTags")
	ret0, _ := ret[0].(v1alpha3.Tags)
	return ret0
}

// AdditionalTags indicates an expected call of AdditionalTags.
func (mr *MockPublicIPScopeMockRecorder) AdditionalTags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdditionalTags", reflect.TypeOf((*MockPublicIPScope)(nil).AdditionalTags))
}

// PublicIPSpecs mocks base method.
func (m *MockPublicIPScope) PublicIPSpecs() []azure.PublicIPSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublicIPSpecs")
	ret0, _ := ret[0].([]azure.PublicIPSpec)
	return ret0
}

// PublicIPSpecs indicates an expected call of PublicIPSpecs.
func (mr *MockPublicIPScopeMockRecorder) PublicIPSpecs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicIPSpecs", reflect.TypeOf((*MockPublicIPScope)(nil).PublicIPSpecs))
}
