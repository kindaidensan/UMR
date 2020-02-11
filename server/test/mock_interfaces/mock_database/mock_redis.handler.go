// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces/database/redis.handler.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRedisHandler is a mock of RedisHandler interface
type MockRedisHandler struct {
	ctrl     *gomock.Controller
	recorder *MockRedisHandlerMockRecorder
}

// MockRedisHandlerMockRecorder is the mock recorder for MockRedisHandler
type MockRedisHandlerMockRecorder struct {
	mock *MockRedisHandler
}

// NewMockRedisHandler creates a new mock instance
func NewMockRedisHandler(ctrl *gomock.Controller) *MockRedisHandler {
	mock := &MockRedisHandler{ctrl: ctrl}
	mock.recorder = &MockRedisHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRedisHandler) EXPECT() *MockRedisHandlerMockRecorder {
	return m.recorder
}

// Set mocks base method
func (m *MockRedisHandler) Set(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockRedisHandlerMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockRedisHandler)(nil).Set), arg0, arg1)
}

// Get mocks base method
func (m *MockRedisHandler) Get(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRedisHandlerMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRedisHandler)(nil).Get), arg0)
}

// RPush mocks base method
func (m *MockRedisHandler) RPush(arg0 string, arg1 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RPush", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RPush indicates an expected call of RPush
func (mr *MockRedisHandlerMockRecorder) RPush(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RPush", reflect.TypeOf((*MockRedisHandler)(nil).RPush), arg0, arg1)
}

// LPop mocks base method
func (m *MockRedisHandler) LPop(arg0 string, arg1 int) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LPop", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LPop indicates an expected call of LPop
func (mr *MockRedisHandlerMockRecorder) LPop(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LPop", reflect.TypeOf((*MockRedisHandler)(nil).LPop), arg0, arg1)
}
