// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces/database/ldap_handler.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockLdapHandler is a mock of LdapHandler interface
type MockLdapHandler struct {
	ctrl     *gomock.Controller
	recorder *MockLdapHandlerMockRecorder
}

// MockLdapHandlerMockRecorder is the mock recorder for MockLdapHandler
type MockLdapHandlerMockRecorder struct {
	mock *MockLdapHandler
}

// NewMockLdapHandler creates a new mock instance
func NewMockLdapHandler(ctrl *gomock.Controller) *MockLdapHandler {
	mock := &MockLdapHandler{ctrl: ctrl}
	mock.recorder = &MockLdapHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLdapHandler) EXPECT() *MockLdapHandlerMockRecorder {
	return m.recorder
}

// AddRequest mocks base method
func (m *MockLdapHandler) AddRequest(arg0 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRequest", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRequest indicates an expected call of AddRequest
func (mr *MockLdapHandlerMockRecorder) AddRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRequest", reflect.TypeOf((*MockLdapHandler)(nil).AddRequest), arg0)
}

// UpdateRequest mocks base method
func (m *MockLdapHandler) UpdateRequest(arg0 string, arg1 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRequest", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRequest indicates an expected call of UpdateRequest
func (mr *MockLdapHandlerMockRecorder) UpdateRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRequest", reflect.TypeOf((*MockLdapHandler)(nil).UpdateRequest), arg0, arg1)
}

// DeleteRequest mocks base method
func (m *MockLdapHandler) DeleteRequest(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRequest", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRequest indicates an expected call of DeleteRequest
func (mr *MockLdapHandlerMockRecorder) DeleteRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRequest", reflect.TypeOf((*MockLdapHandler)(nil).DeleteRequest), arg0)
}

// SearchRequest mocks base method
func (m *MockLdapHandler) SearchRequest(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchRequest", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SearchRequest indicates an expected call of SearchRequest
func (mr *MockLdapHandlerMockRecorder) SearchRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchRequest", reflect.TypeOf((*MockLdapHandler)(nil).SearchRequest), arg0)
}
