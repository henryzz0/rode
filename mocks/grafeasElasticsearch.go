// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rode/grafeas-elasticsearch/go/v1beta1/storage/filtering (interfaces: Filterer)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	filtering "github.com/rode/grafeas-elasticsearch/go/v1beta1/storage/filtering"
	reflect "reflect"
)

// MockFilterer is a mock of Filterer interface
type MockFilterer struct {
	ctrl     *gomock.Controller
	recorder *MockFiltererMockRecorder
}

// MockFiltererMockRecorder is the mock recorder for MockFilterer
type MockFiltererMockRecorder struct {
	mock *MockFilterer
}

// NewMockFilterer creates a new mock instance
func NewMockFilterer(ctrl *gomock.Controller) *MockFilterer {
	mock := &MockFilterer{ctrl: ctrl}
	mock.recorder = &MockFiltererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFilterer) EXPECT() *MockFiltererMockRecorder {
	return m.recorder
}

// ParseExpression mocks base method
func (m *MockFilterer) ParseExpression(arg0 string) (*filtering.Query, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseExpression", arg0)
	ret0, _ := ret[0].(*filtering.Query)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseExpression indicates an expected call of ParseExpression
func (mr *MockFiltererMockRecorder) ParseExpression(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseExpression", reflect.TypeOf((*MockFilterer)(nil).ParseExpression), arg0)
}