// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/common (interfaces: ToolsFinder,ToolsFindEntity,ToolsURLGetter,APIHostPortsForAgentsGetter,ToolsStorageGetter,AgentTooler)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/tools_mock.go github.com/juju/juju/apiserver/common ToolsFinder,ToolsFindEntity,ToolsURLGetter,APIHostPortsForAgentsGetter,ToolsStorageGetter,AgentTooler
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	common "github.com/juju/juju/apiserver/common"
	network "github.com/juju/juju/core/network"
	state "github.com/juju/juju/state"
	binarystorage "github.com/juju/juju/state/binarystorage"
	tools "github.com/juju/juju/tools"
	names "github.com/juju/names/v5"
	version "github.com/juju/version/v2"
	gomock "go.uber.org/mock/gomock"
)

// MockToolsFinder is a mock of ToolsFinder interface.
type MockToolsFinder struct {
	ctrl     *gomock.Controller
	recorder *MockToolsFinderMockRecorder
}

// MockToolsFinderMockRecorder is the mock recorder for MockToolsFinder.
type MockToolsFinderMockRecorder struct {
	mock *MockToolsFinder
}

// NewMockToolsFinder creates a new mock instance.
func NewMockToolsFinder(ctrl *gomock.Controller) *MockToolsFinder {
	mock := &MockToolsFinder{ctrl: ctrl}
	mock.recorder = &MockToolsFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockToolsFinder) EXPECT() *MockToolsFinderMockRecorder {
	return m.recorder
}

// FindAgents mocks base method.
func (m *MockToolsFinder) FindAgents(arg0 common.FindAgentsParams) (tools.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAgents", arg0)
	ret0, _ := ret[0].(tools.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAgents indicates an expected call of FindAgents.
func (mr *MockToolsFinderMockRecorder) FindAgents(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAgents", reflect.TypeOf((*MockToolsFinder)(nil).FindAgents), arg0)
}

// MockToolsFindEntity is a mock of ToolsFindEntity interface.
type MockToolsFindEntity struct {
	ctrl     *gomock.Controller
	recorder *MockToolsFindEntityMockRecorder
}

// MockToolsFindEntityMockRecorder is the mock recorder for MockToolsFindEntity.
type MockToolsFindEntityMockRecorder struct {
	mock *MockToolsFindEntity
}

// NewMockToolsFindEntity creates a new mock instance.
func NewMockToolsFindEntity(ctrl *gomock.Controller) *MockToolsFindEntity {
	mock := &MockToolsFindEntity{ctrl: ctrl}
	mock.recorder = &MockToolsFindEntityMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockToolsFindEntity) EXPECT() *MockToolsFindEntityMockRecorder {
	return m.recorder
}

// FindEntity mocks base method.
func (m *MockToolsFindEntity) FindEntity(arg0 names.Tag) (state.Entity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEntity", arg0)
	ret0, _ := ret[0].(state.Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEntity indicates an expected call of FindEntity.
func (mr *MockToolsFindEntityMockRecorder) FindEntity(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEntity", reflect.TypeOf((*MockToolsFindEntity)(nil).FindEntity), arg0)
}

// MockToolsURLGetter is a mock of ToolsURLGetter interface.
type MockToolsURLGetter struct {
	ctrl     *gomock.Controller
	recorder *MockToolsURLGetterMockRecorder
}

// MockToolsURLGetterMockRecorder is the mock recorder for MockToolsURLGetter.
type MockToolsURLGetterMockRecorder struct {
	mock *MockToolsURLGetter
}

// NewMockToolsURLGetter creates a new mock instance.
func NewMockToolsURLGetter(ctrl *gomock.Controller) *MockToolsURLGetter {
	mock := &MockToolsURLGetter{ctrl: ctrl}
	mock.recorder = &MockToolsURLGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockToolsURLGetter) EXPECT() *MockToolsURLGetterMockRecorder {
	return m.recorder
}

// ToolsURLs mocks base method.
func (m *MockToolsURLGetter) ToolsURLs(arg0 version.Binary) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToolsURLs", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToolsURLs indicates an expected call of ToolsURLs.
func (mr *MockToolsURLGetterMockRecorder) ToolsURLs(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToolsURLs", reflect.TypeOf((*MockToolsURLGetter)(nil).ToolsURLs), arg0)
}

// MockAPIHostPortsForAgentsGetter is a mock of APIHostPortsForAgentsGetter interface.
type MockAPIHostPortsForAgentsGetter struct {
	ctrl     *gomock.Controller
	recorder *MockAPIHostPortsForAgentsGetterMockRecorder
}

// MockAPIHostPortsForAgentsGetterMockRecorder is the mock recorder for MockAPIHostPortsForAgentsGetter.
type MockAPIHostPortsForAgentsGetterMockRecorder struct {
	mock *MockAPIHostPortsForAgentsGetter
}

// NewMockAPIHostPortsForAgentsGetter creates a new mock instance.
func NewMockAPIHostPortsForAgentsGetter(ctrl *gomock.Controller) *MockAPIHostPortsForAgentsGetter {
	mock := &MockAPIHostPortsForAgentsGetter{ctrl: ctrl}
	mock.recorder = &MockAPIHostPortsForAgentsGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPIHostPortsForAgentsGetter) EXPECT() *MockAPIHostPortsForAgentsGetterMockRecorder {
	return m.recorder
}

// APIHostPortsForAgents mocks base method.
func (m *MockAPIHostPortsForAgentsGetter) APIHostPortsForAgents() ([]network.SpaceHostPorts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIHostPortsForAgents")
	ret0, _ := ret[0].([]network.SpaceHostPorts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// APIHostPortsForAgents indicates an expected call of APIHostPortsForAgents.
func (mr *MockAPIHostPortsForAgentsGetterMockRecorder) APIHostPortsForAgents() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIHostPortsForAgents", reflect.TypeOf((*MockAPIHostPortsForAgentsGetter)(nil).APIHostPortsForAgents))
}

// MockToolsStorageGetter is a mock of ToolsStorageGetter interface.
type MockToolsStorageGetter struct {
	ctrl     *gomock.Controller
	recorder *MockToolsStorageGetterMockRecorder
}

// MockToolsStorageGetterMockRecorder is the mock recorder for MockToolsStorageGetter.
type MockToolsStorageGetterMockRecorder struct {
	mock *MockToolsStorageGetter
}

// NewMockToolsStorageGetter creates a new mock instance.
func NewMockToolsStorageGetter(ctrl *gomock.Controller) *MockToolsStorageGetter {
	mock := &MockToolsStorageGetter{ctrl: ctrl}
	mock.recorder = &MockToolsStorageGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockToolsStorageGetter) EXPECT() *MockToolsStorageGetterMockRecorder {
	return m.recorder
}

// ToolsStorage mocks base method.
func (m *MockToolsStorageGetter) ToolsStorage() (binarystorage.StorageCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToolsStorage")
	ret0, _ := ret[0].(binarystorage.StorageCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToolsStorage indicates an expected call of ToolsStorage.
func (mr *MockToolsStorageGetterMockRecorder) ToolsStorage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToolsStorage", reflect.TypeOf((*MockToolsStorageGetter)(nil).ToolsStorage))
}

// MockAgentTooler is a mock of AgentTooler interface.
type MockAgentTooler struct {
	ctrl     *gomock.Controller
	recorder *MockAgentToolerMockRecorder
}

// MockAgentToolerMockRecorder is the mock recorder for MockAgentTooler.
type MockAgentToolerMockRecorder struct {
	mock *MockAgentTooler
}

// NewMockAgentTooler creates a new mock instance.
func NewMockAgentTooler(ctrl *gomock.Controller) *MockAgentTooler {
	mock := &MockAgentTooler{ctrl: ctrl}
	mock.recorder = &MockAgentToolerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAgentTooler) EXPECT() *MockAgentToolerMockRecorder {
	return m.recorder
}

// AgentTools mocks base method.
func (m *MockAgentTooler) AgentTools() (*tools.Tools, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AgentTools")
	ret0, _ := ret[0].(*tools.Tools)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AgentTools indicates an expected call of AgentTools.
func (mr *MockAgentToolerMockRecorder) AgentTools() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AgentTools", reflect.TypeOf((*MockAgentTooler)(nil).AgentTools))
}

// SetAgentVersion mocks base method.
func (m *MockAgentTooler) SetAgentVersion(arg0 version.Binary) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAgentVersion", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetAgentVersion indicates an expected call of SetAgentVersion.
func (mr *MockAgentToolerMockRecorder) SetAgentVersion(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAgentVersion", reflect.TypeOf((*MockAgentTooler)(nil).SetAgentVersion), arg0)
}

// Tag mocks base method.
func (m *MockAgentTooler) Tag() names.Tag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tag")
	ret0, _ := ret[0].(names.Tag)
	return ret0
}

// Tag indicates an expected call of Tag.
func (mr *MockAgentToolerMockRecorder) Tag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tag", reflect.TypeOf((*MockAgentTooler)(nil).Tag))
}
