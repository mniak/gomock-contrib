// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mniak/gomock-contrib/internal/testing (interfaces: MatcherGotFormatter)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMatcherGotFormatter is a mock of MatcherGotFormatter interface.
type MockMatcherGotFormatter struct {
	ctrl     *gomock.Controller
	recorder *MockMatcherGotFormatterMockRecorder
}

// MockMatcherGotFormatterMockRecorder is the mock recorder for MockMatcherGotFormatter.
type MockMatcherGotFormatterMockRecorder struct {
	mock *MockMatcherGotFormatter
}

// NewMockMatcherGotFormatter creates a new mock instance.
func NewMockMatcherGotFormatter(ctrl *gomock.Controller) *MockMatcherGotFormatter {
	mock := &MockMatcherGotFormatter{ctrl: ctrl}
	mock.recorder = &MockMatcherGotFormatterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMatcherGotFormatter) EXPECT() *MockMatcherGotFormatterMockRecorder {
	return m.recorder
}

// Got mocks base method.
func (m *MockMatcherGotFormatter) Got(arg0 interface{}) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Got", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Got indicates an expected call of Got.
func (mr *MockMatcherGotFormatterMockRecorder) Got(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Got", reflect.TypeOf((*MockMatcherGotFormatter)(nil).Got), arg0)
}

// Matches mocks base method.
func (m *MockMatcherGotFormatter) Matches(arg0 interface{}) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Matches", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Matches indicates an expected call of Matches.
func (mr *MockMatcherGotFormatterMockRecorder) Matches(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Matches", reflect.TypeOf((*MockMatcherGotFormatter)(nil).Matches), arg0)
}

// String mocks base method.
func (m *MockMatcherGotFormatter) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockMatcherGotFormatterMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockMatcherGotFormatter)(nil).String))
}
