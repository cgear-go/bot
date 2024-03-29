// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cgear-go/bot/discord/command (interfaces: Command)

// Package commandmock is a generated GoMock package.
package commandmock

import (
	reflect "reflect"

	client "github.com/cgear-go/bot/discord/client"
	command "github.com/cgear-go/bot/discord/command"
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

// Execute mocks base method.
func (m *MockCommand) Execute(arg0 client.Client, arg1 command.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockCommandMockRecorder) Execute(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockCommand)(nil).Execute), arg0, arg1)
}

// Name mocks base method.
func (m *MockCommand) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockCommandMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockCommand)(nil).Name))
}
