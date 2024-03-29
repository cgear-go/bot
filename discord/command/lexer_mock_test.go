// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cgear-go/bot/discord/command (interfaces: Lexer)

// Package command is a generated GoMock package.
package command

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLexer is a mock of Lexer interface.
type MockLexer struct {
	ctrl     *gomock.Controller
	recorder *MockLexerMockRecorder
}

// MockLexerMockRecorder is the mock recorder for MockLexer.
type MockLexerMockRecorder struct {
	mock *MockLexer
}

// NewMockLexer creates a new mock instance.
func NewMockLexer(ctrl *gomock.Controller) *MockLexer {
	mock := &MockLexer{ctrl: ctrl}
	mock.recorder = &MockLexerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLexer) EXPECT() *MockLexerMockRecorder {
	return m.recorder
}

// HasNext mocks base method.
func (m *MockLexer) HasNext() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasNext")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasNext indicates an expected call of HasNext.
func (mr *MockLexerMockRecorder) HasNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasNext", reflect.TypeOf((*MockLexer)(nil).HasNext))
}

// Next mocks base method.
func (m *MockLexer) Next() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Next indicates an expected call of Next.
func (mr *MockLexerMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockLexer)(nil).Next))
}
