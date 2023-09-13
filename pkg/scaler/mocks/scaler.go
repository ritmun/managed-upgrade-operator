// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/managed-upgrade-operator/pkg/scaler (interfaces: Scaler)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	logr "github.com/go-logr/logr"
	gomock "github.com/golang/mock/gomock"
	drain "github.com/openshift/managed-upgrade-operator/pkg/drain"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockScaler is a mock of Scaler interface.
type MockScaler struct {
	ctrl     *gomock.Controller
	recorder *MockScalerMockRecorder
}

// MockScalerMockRecorder is the mock recorder for MockScaler.
type MockScalerMockRecorder struct {
	mock *MockScaler
}

// NewMockScaler creates a new mock instance.
func NewMockScaler(ctrl *gomock.Controller) *MockScaler {
	mock := &MockScaler{ctrl: ctrl}
	mock.recorder = &MockScalerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScaler) EXPECT() *MockScalerMockRecorder {
	return m.recorder
}

// CanScale mocks base method.
func (m *MockScaler) CanScale(arg0 client.Client, arg1 logr.Logger) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanScale", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CanScale indicates an expected call of CanScale.
func (mr *MockScalerMockRecorder) CanScale(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanScale", reflect.TypeOf((*MockScaler)(nil).CanScale), arg0, arg1)
}

// EnsureScaleDownNodes mocks base method.
func (m *MockScaler) EnsureScaleDownNodes(arg0 client.Client, arg1 drain.NodeDrainStrategy, arg2 logr.Logger) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureScaleDownNodes", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureScaleDownNodes indicates an expected call of EnsureScaleDownNodes.
func (mr *MockScalerMockRecorder) EnsureScaleDownNodes(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureScaleDownNodes", reflect.TypeOf((*MockScaler)(nil).EnsureScaleDownNodes), arg0, arg1, arg2)
}

// EnsureScaleUpNodes mocks base method.
func (m *MockScaler) EnsureScaleUpNodes(arg0 client.Client, arg1 time.Duration, arg2 logr.Logger) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureScaleUpNodes", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureScaleUpNodes indicates an expected call of EnsureScaleUpNodes.
func (mr *MockScalerMockRecorder) EnsureScaleUpNodes(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureScaleUpNodes", reflect.TypeOf((*MockScaler)(nil).EnsureScaleUpNodes), arg0, arg1, arg2)
}
