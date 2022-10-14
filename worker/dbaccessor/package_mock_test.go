// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/worker/dbaccessor (interfaces: Logger,DBApp)

// Package dbaccessor is a generated GoMock package.
package dbaccessor

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	loggo "github.com/juju/loggo"
)

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Debugf mocks base method.
func (m *MockLogger) Debugf(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf.
func (mr *MockLoggerMockRecorder) Debugf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockLogger)(nil).Debugf), varargs...)
}

// Errorf mocks base method.
func (m *MockLogger) Errorf(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockLoggerMockRecorder) Errorf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockLogger)(nil).Errorf), varargs...)
}

// Infof mocks base method.
func (m *MockLogger) Infof(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof.
func (mr *MockLoggerMockRecorder) Infof(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockLogger)(nil).Infof), varargs...)
}

// Logf mocks base method.
func (m *MockLogger) Logf(arg0 loggo.Level, arg1 string, arg2 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Logf", varargs...)
}

// Logf indicates an expected call of Logf.
func (mr *MockLoggerMockRecorder) Logf(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logf", reflect.TypeOf((*MockLogger)(nil).Logf), varargs...)
}

// Tracef mocks base method.
func (m *MockLogger) Tracef(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Tracef", varargs...)
}

// Tracef indicates an expected call of Tracef.
func (mr *MockLoggerMockRecorder) Tracef(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tracef", reflect.TypeOf((*MockLogger)(nil).Tracef), varargs...)
}

// Warningf mocks base method.
func (m *MockLogger) Warningf(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warningf", varargs...)
}

// Warningf indicates an expected call of Warningf.
func (mr *MockLoggerMockRecorder) Warningf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warningf", reflect.TypeOf((*MockLogger)(nil).Warningf), varargs...)
}

// MockDBApp is a mock of DBApp interface.
type MockDBApp struct {
	ctrl     *gomock.Controller
	recorder *MockDBAppMockRecorder
}

// MockDBAppMockRecorder is the mock recorder for MockDBApp.
type MockDBAppMockRecorder struct {
	mock *MockDBApp
}

// NewMockDBApp creates a new mock instance.
func NewMockDBApp(ctrl *gomock.Controller) *MockDBApp {
	mock := &MockDBApp{ctrl: ctrl}
	mock.recorder = &MockDBAppMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBApp) EXPECT() *MockDBAppMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockDBApp) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockDBAppMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDBApp)(nil).Close))
}

// Handover mocks base method.
func (m *MockDBApp) Handover(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handover", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Handover indicates an expected call of Handover.
func (mr *MockDBAppMockRecorder) Handover(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handover", reflect.TypeOf((*MockDBApp)(nil).Handover), arg0)
}

// ID mocks base method.
func (m *MockDBApp) ID() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockDBAppMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockDBApp)(nil).ID))
}

// Open mocks base method.
func (m *MockDBApp) Open(arg0 context.Context, arg1 string) (*sql.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", arg0, arg1)
	ret0, _ := ret[0].(*sql.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open.
func (mr *MockDBAppMockRecorder) Open(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockDBApp)(nil).Open), arg0, arg1)
}

// Ready mocks base method.
func (m *MockDBApp) Ready(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ready", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ready indicates an expected call of Ready.
func (mr *MockDBAppMockRecorder) Ready(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ready", reflect.TypeOf((*MockDBApp)(nil).Ready), arg0)
}
