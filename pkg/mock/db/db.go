// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_backends is a generated GoMock package.
package mock_backends

import (
	_go "github.com/alibaba/morphling/api/v1alpha1/grpc_storage/go"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStorageBackend is a mock of StorageBackend interface
type MockStorageBackend struct {
	ctrl     *gomock.Controller
	recorder *MockStorageBackendMockRecorder
}

// MockStorageBackendMockRecorder is the mock recorder for MockStorageBackend
type MockStorageBackendMockRecorder struct {
	mock *MockStorageBackend
}

// NewMockStorageBackend creates a new mock instance
func NewMockStorageBackend(ctrl *gomock.Controller) *MockStorageBackend {
	mock := &MockStorageBackend{ctrl: ctrl}
	mock.recorder = &MockStorageBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorageBackend) EXPECT() *MockStorageBackendMockRecorder {
	return m.recorder
}

// Initialize mocks base method
func (m *MockStorageBackend) Initialize() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize")
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize
func (mr *MockStorageBackendMockRecorder) Initialize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockStorageBackend)(nil).Initialize))
}

// Close mocks base method
func (m *MockStorageBackend) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockStorageBackendMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStorageBackend)(nil).Close))
}

// Name mocks base method
func (m *MockStorageBackend) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockStorageBackendMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockStorageBackend)(nil).Name))
}

// SaveTrialResult mocks base method
func (m *MockStorageBackend) SaveTrialResult(observationLog *_go.SaveResultRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTrialResult", observationLog)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveTrialResult indicates an expected call of SaveTrialResult
func (mr *MockStorageBackendMockRecorder) SaveTrialResult(observationLog interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTrialResult", reflect.TypeOf((*MockStorageBackend)(nil).SaveTrialResult), observationLog)
}

// GetTrialResult mocks base method
func (m *MockStorageBackend) GetTrialResult(request *_go.GetResultRequest) (*_go.GetResultReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrialResult", request)
	ret0, _ := ret[0].(*_go.GetResultReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrialResult indicates an expected call of GetTrialResult
func (mr *MockStorageBackendMockRecorder) GetTrialResult(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrialResult", reflect.TypeOf((*MockStorageBackend)(nil).GetTrialResult), request)
}
