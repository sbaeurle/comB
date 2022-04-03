// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sbaeurle/comb/metrics/outputs (interfaces: Output)

// Package mock_outputs is a generated GoMock package.
package mock_outputs

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOutput is a mock of Output interface.
type MockOutput struct {
	ctrl     *gomock.Controller
	recorder *MockOutputMockRecorder
}

// MockOutputMockRecorder is the mock recorder for MockOutput.
type MockOutputMockRecorder struct {
	mock *MockOutput
}

// NewMockOutput creates a new mock instance.
func NewMockOutput(ctrl *gomock.Controller) *MockOutput {
	mock := &MockOutput{ctrl: ctrl}
	mock.recorder = &MockOutputMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOutput) EXPECT() *MockOutputMockRecorder {
	return m.recorder
}

// ChangePath mocks base method.
func (m *MockOutput) ChangePath(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePath", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePath indicates an expected call of ChangePath.
func (mr *MockOutputMockRecorder) ChangePath(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePath", reflect.TypeOf((*MockOutput)(nil).ChangePath), arg0)
}

// WriteResult mocks base method.
func (m *MockOutput) WriteResult(arg0 map[string]float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteResult", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteResult indicates an expected call of WriteResult.
func (mr *MockOutputMockRecorder) WriteResult(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteResult", reflect.TypeOf((*MockOutput)(nil).WriteResult), arg0)
}
