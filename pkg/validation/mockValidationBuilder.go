// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/managed-upgrade-operator/pkg/validation (interfaces: ValidationBuilder)

// Package validation is a generated GoMock package.
package validation

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockValidationBuilder is a mock of ValidationBuilder interface
type MockValidationBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockValidationBuilderMockRecorder
}

// MockValidationBuilderMockRecorder is the mock recorder for MockValidationBuilder
type MockValidationBuilderMockRecorder struct {
	mock *MockValidationBuilder
}

// NewMockValidationBuilder creates a new mock instance
func NewMockValidationBuilder(ctrl *gomock.Controller) *MockValidationBuilder {
	mock := &MockValidationBuilder{ctrl: ctrl}
	mock.recorder = &MockValidationBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockValidationBuilder) EXPECT() *MockValidationBuilderMockRecorder {
	return m.recorder
}

// NewClient mocks base method
func (m *MockValidationBuilder) NewClient() (Validator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewClient")
	ret0, _ := ret[0].(Validator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewClient indicates an expected call of NewClient
func (mr *MockValidationBuilderMockRecorder) NewClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewClient", reflect.TypeOf((*MockValidationBuilder)(nil).NewClient))
}