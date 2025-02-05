// Code generated by MockGen. DO NOT EDIT.
// Source: ../../notifications-client/notifier/notifier.go

// Package mock_notifier is a generated GoMock package.
package mock_notifier

import (
	context "context"
	reflect "reflect"

	api "github.com/chef/automate/components/notifications-client/api"
	gomock "github.com/golang/mock/gomock"
)

// MockNotifier is a mock of Notifier interface.
type MockNotifier struct {
	ctrl     *gomock.Controller
	recorder *MockNotifierMockRecorder
}

// MockNotifierMockRecorder is the mock recorder for MockNotifier.
type MockNotifierMockRecorder struct {
	mock *MockNotifier
}

// NewMockNotifier creates a new mock instance.
func NewMockNotifier(ctrl *gomock.Controller) *MockNotifier {
	mock := &MockNotifier{ctrl: ctrl}
	mock.recorder = &MockNotifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotifier) EXPECT() *MockNotifierMockRecorder {
	return m.recorder
}

// QueueSize mocks base method.
func (m *MockNotifier) QueueSize() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueSize")
	ret0, _ := ret[0].(int)
	return ret0
}

// QueueSize indicates an expected call of QueueSize.
func (mr *MockNotifierMockRecorder) QueueSize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueSize", reflect.TypeOf((*MockNotifier)(nil).QueueSize))
}

// Send mocks base method.
func (m *MockNotifier) Send(arg0 context.Context, arg1 *api.Event) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Send", arg0, arg1)
}

// Send indicates an expected call of Send.
func (mr *MockNotifierMockRecorder) Send(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockNotifier)(nil).Send), arg0, arg1)
}
