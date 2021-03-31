// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jonathanarnault/cgear-go/go/discord/reaction (interfaces: Reaction)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	client "github.com/jonathanarnault/cgear-go/go/discord/client"
	reaction "github.com/jonathanarnault/cgear-go/go/discord/reaction"
)

// MockReaction is a mock of Reaction interface.
type MockReaction struct {
	ctrl     *gomock.Controller
	recorder *MockReactionMockRecorder
}

// MockReactionMockRecorder is the mock recorder for MockReaction.
type MockReactionMockRecorder struct {
	mock *MockReaction
}

// NewMockReaction creates a new mock instance.
func NewMockReaction(ctrl *gomock.Controller) *MockReaction {
	mock := &MockReaction{ctrl: ctrl}
	mock.recorder = &MockReactionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReaction) EXPECT() *MockReactionMockRecorder {
	return m.recorder
}

// Added mocks base method.
func (m *MockReaction) Added(arg0 client.Client, arg1 reaction.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Added", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Added indicates an expected call of Added.
func (mr *MockReactionMockRecorder) Added(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Added", reflect.TypeOf((*MockReaction)(nil).Added), arg0, arg1)
}

// Emoji mocks base method.
func (m *MockReaction) Emoji() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Emoji")
	ret0, _ := ret[0].(string)
	return ret0
}

// Emoji indicates an expected call of Emoji.
func (mr *MockReactionMockRecorder) Emoji() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Emoji", reflect.TypeOf((*MockReaction)(nil).Emoji))
}

// Removed mocks base method.
func (m *MockReaction) Removed(arg0 client.Client, arg1 reaction.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Removed", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Removed indicates an expected call of Removed.
func (mr *MockReactionMockRecorder) Removed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Removed", reflect.TypeOf((*MockReaction)(nil).Removed), arg0, arg1)
}
