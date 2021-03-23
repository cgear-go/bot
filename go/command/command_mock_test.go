// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jonathanarnault/cgear-go/go/command (interfaces: Command)

// Package command is a generated GoMock package.
package command

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCommand is a mock of Command interface.
type MockCommand struct {
	ctrl     *gomock.Controller
	recorder *MockCommandMockRecorder
}

// MockCommandMockRecorder is the mock recorder for MockCommand.
type MockCommandMockRecorder struct {
	mock *MockCommand
}

// NewMockCommand creates a new mock instance.
func NewMockCommand(ctrl *gomock.Controller) *MockCommand {
	mock := &MockCommand{ctrl: ctrl}
	mock.recorder = &MockCommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommand) EXPECT() *MockCommandMockRecorder {
	return m.recorder
}

// AddInt mocks base method.
func (m *MockCommand) AddInt(arg0 string) Command {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddInt", arg0)
	ret0, _ := ret[0].(Command)
	return ret0
}

// AddInt indicates an expected call of AddInt.
func (mr *MockCommandMockRecorder) AddInt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddInt", reflect.TypeOf((*MockCommand)(nil).AddInt), arg0)
}

// AddResolver mocks base method.
func (m *MockCommand) AddResolver(arg0 CommandFn) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddResolver", arg0)
}

// AddResolver indicates an expected call of AddResolver.
func (mr *MockCommandMockRecorder) AddResolver(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddResolver", reflect.TypeOf((*MockCommand)(nil).AddResolver), arg0)
}

// AddRest mocks base method.
func (m *MockCommand) AddRest(arg0 string) Command {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRest", arg0)
	ret0, _ := ret[0].(Command)
	return ret0
}

// AddRest indicates an expected call of AddRest.
func (mr *MockCommandMockRecorder) AddRest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRest", reflect.TypeOf((*MockCommand)(nil).AddRest), arg0)
}

// AddString mocks base method.
func (m *MockCommand) AddString(arg0 string) Command {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddString", arg0)
	ret0, _ := ret[0].(Command)
	return ret0
}

// AddString indicates an expected call of AddString.
func (mr *MockCommandMockRecorder) AddString(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddString", reflect.TypeOf((*MockCommand)(nil).AddString), arg0)
}

// execute mocks base method.
func (m *MockCommand) execute(arg0 context.Context, arg1 Parser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "execute", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// execute indicates an expected call of execute.
func (mr *MockCommandMockRecorder) execute(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "execute", reflect.TypeOf((*MockCommand)(nil).execute), arg0, arg1)
}
