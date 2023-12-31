// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-application-networking-k8s/pkg/model/core (interfaces: Stack)

// Package core is a generated GoMock package.
package core

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStack is a mock of Stack interface.
type MockStack struct {
	ctrl     *gomock.Controller
	recorder *MockStackMockRecorder
}

// MockStackMockRecorder is the mock recorder for MockStack.
type MockStackMockRecorder struct {
	mock *MockStack
}

// NewMockStack creates a new mock instance.
func NewMockStack(ctrl *gomock.Controller) *MockStack {
	mock := &MockStack{ctrl: ctrl}
	mock.recorder = &MockStackMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStack) EXPECT() *MockStackMockRecorder {
	return m.recorder
}

// AddDependency mocks base method.
func (m *MockStack) AddDependency(arg0, arg1 Resource) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDependency", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddDependency indicates an expected call of AddDependency.
func (mr *MockStackMockRecorder) AddDependency(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDependency", reflect.TypeOf((*MockStack)(nil).AddDependency), arg0, arg1)
}

// AddResource mocks base method.
func (m *MockStack) AddResource(arg0 Resource) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddResource", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddResource indicates an expected call of AddResource.
func (mr *MockStackMockRecorder) AddResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddResource", reflect.TypeOf((*MockStack)(nil).AddResource), arg0)
}

// GetResource mocks base method.
func (m *MockStack) GetResource(arg0 string, arg1 Resource) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResource", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetResource indicates an expected call of GetResource.
func (mr *MockStackMockRecorder) GetResource(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResource", reflect.TypeOf((*MockStack)(nil).GetResource), arg0, arg1)
}

// ListResources mocks base method.
func (m *MockStack) ListResources(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListResources", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListResources indicates an expected call of ListResources.
func (mr *MockStackMockRecorder) ListResources(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResources", reflect.TypeOf((*MockStack)(nil).ListResources), arg0)
}

// StackID mocks base method.
func (m *MockStack) StackID() StackID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StackID")
	ret0, _ := ret[0].(StackID)
	return ret0
}

// StackID indicates an expected call of StackID.
func (mr *MockStackMockRecorder) StackID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StackID", reflect.TypeOf((*MockStack)(nil).StackID))
}

// TopologicalTraversal mocks base method.
func (m *MockStack) TopologicalTraversal(arg0 ResourceVisitor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TopologicalTraversal", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// TopologicalTraversal indicates an expected call of TopologicalTraversal.
func (mr *MockStackMockRecorder) TopologicalTraversal(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TopologicalTraversal", reflect.TypeOf((*MockStack)(nil).TopologicalTraversal), arg0)
}
