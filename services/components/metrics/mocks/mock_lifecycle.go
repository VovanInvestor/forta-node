// Code generated by MockGen. DO NOT EDIT.
// Source: services/components/metrics/lifecycle.go

// Package mock_metrics is a generated GoMock package.
package mock_metrics

import (
	reflect "reflect"

	domain "github.com/forta-network/forta-core-go/domain"
	config "github.com/forta-network/forta-node/config"
	gomock "github.com/golang/mock/gomock"
)

// MockLifecycle is a mock of Lifecycle interface.
type MockLifecycle struct {
	ctrl     *gomock.Controller
	recorder *MockLifecycleMockRecorder
}

// MockLifecycleMockRecorder is the mock recorder for MockLifecycle.
type MockLifecycleMockRecorder struct {
	mock *MockLifecycle
}

// NewMockLifecycle creates a new mock instance.
func NewMockLifecycle(ctrl *gomock.Controller) *MockLifecycle {
	mock := &MockLifecycle{ctrl: ctrl}
	mock.recorder = &MockLifecycleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLifecycle) EXPECT() *MockLifecycleMockRecorder {
	return m.recorder
}

// ActionRestart mocks base method.
func (m *MockLifecycle) ActionRestart(arg0 ...config.AgentConfig) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "ActionRestart", varargs...)
}

// ActionRestart indicates an expected call of ActionRestart.
func (mr *MockLifecycleMockRecorder) ActionRestart(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionRestart", reflect.TypeOf((*MockLifecycle)(nil).ActionRestart), arg0...)
}

// ActionSubscribe mocks base method.
func (m *MockLifecycle) ActionSubscribe(arg0 []domain.CombinerBotSubscription) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ActionSubscribe", arg0)
}

// ActionSubscribe indicates an expected call of ActionSubscribe.
func (mr *MockLifecycleMockRecorder) ActionSubscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionSubscribe", reflect.TypeOf((*MockLifecycle)(nil).ActionSubscribe), arg0)
}

// ActionUnsubscribe mocks base method.
func (m *MockLifecycle) ActionUnsubscribe(arg0 []domain.CombinerBotSubscription) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ActionUnsubscribe", arg0)
}

// ActionUnsubscribe indicates an expected call of ActionUnsubscribe.
func (mr *MockLifecycleMockRecorder) ActionUnsubscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionUnsubscribe", reflect.TypeOf((*MockLifecycle)(nil).ActionUnsubscribe), arg0)
}

// ActionUpdate mocks base method.
func (m *MockLifecycle) ActionUpdate(arg0 ...config.AgentConfig) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "ActionUpdate", varargs...)
}

// ActionUpdate indicates an expected call of ActionUpdate.
func (mr *MockLifecycleMockRecorder) ActionUpdate(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionUpdate", reflect.TypeOf((*MockLifecycle)(nil).ActionUpdate), arg0...)
}

// BotError mocks base method.
func (m *MockLifecycle) BotError(metricName string, err error, botID ...string) {
	m.ctrl.T.Helper()
	varargs := []interface{}{metricName, err}
	for _, a := range botID {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "BotError", varargs...)
}

// BotError indicates an expected call of BotError.
func (mr *MockLifecycleMockRecorder) BotError(metricName, err interface{}, botID ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{metricName, err}, botID...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BotError", reflect.TypeOf((*MockLifecycle)(nil).BotError), varargs...)
}

// ClientClose mocks base method.
func (m *MockLifecycle) ClientClose(arg0 ...config.AgentConfig) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "ClientClose", varargs...)
}

// ClientClose indicates an expected call of ClientClose.
func (mr *MockLifecycleMockRecorder) ClientClose(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientClose", reflect.TypeOf((*MockLifecycle)(nil).ClientClose), arg0...)
}

// ClientDial mocks base method.
func (m *MockLifecycle) ClientDial(arg0 ...config.AgentConfig) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "ClientDial", varargs...)
}

// ClientDial indicates an expected call of ClientDial.
func (mr *MockLifecycleMockRecorder) ClientDial(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientDial", reflect.TypeOf((*MockLifecycle)(nil).ClientDial), arg0...)
}

// FailureDial mocks base method.
func (m *MockLifecycle) FailureDial(arg0 error, arg1 ...config.AgentConfig) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "FailureDial", varargs...)
}

// FailureDial indicates an expected call of FailureDial.
func (mr *MockLifecycleMockRecorder) FailureDial(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailureDial", reflect.TypeOf((*MockLifecycle)(nil).FailureDial), varargs...)
}

// FailureInitialize mocks base method.
func (m *MockLifecycle) FailureInitialize(arg0 error, arg1 ...config.AgentConfig) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "FailureInitialize", varargs...)
}

// FailureInitialize indicates an expected call of FailureInitialize.
func (mr *MockLifecycleMockRecorder) FailureInitialize(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailureInitialize", reflect.TypeOf((*MockLifecycle)(nil).FailureInitialize), varargs...)
}

// FailureInitializeResponse mocks base method.
func (m *MockLifecycle) FailureInitializeResponse(arg0 error, arg1 ...config.AgentConfig) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "FailureInitializeResponse", varargs...)
}

// FailureInitializeResponse indicates an expected call of FailureInitializeResponse.
func (mr *MockLifecycleMockRecorder) FailureInitializeResponse(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailureInitializeResponse", reflect.TypeOf((*MockLifecycle)(nil).FailureInitializeResponse), varargs...)
}

// FailureInitializeValidate mocks base method.
func (m *MockLifecycle) FailureInitializeValidate(arg0 error, arg1 ...config.AgentConfig) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "FailureInitializeValidate", varargs...)
}

// FailureInitializeValidate indicates an expected call of FailureInitializeValidate.
func (mr *MockLifecycleMockRecorder) FailureInitializeValidate(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailureInitializeValidate", reflect.TypeOf((*MockLifecycle)(nil).FailureInitializeValidate), varargs...)
}

// FailureLaunch mocks base method.
func (m *MockLifecycle) FailureLaunch(arg0 error, arg1 ...config.AgentConfig) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "FailureLaunch", varargs...)
}

// FailureLaunch indicates an expected call of FailureLaunch.
func (mr *MockLifecycleMockRecorder) FailureLaunch(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailureLaunch", reflect.TypeOf((*MockLifecycle)(nil).FailureLaunch), varargs...)
}

// FailurePull mocks base method.
func (m *MockLifecycle) FailurePull(arg0 error, arg1 ...config.AgentConfig) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "FailurePull", varargs...)
}

// FailurePull indicates an expected call of FailurePull.
func (mr *MockLifecycleMockRecorder) FailurePull(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailurePull", reflect.TypeOf((*MockLifecycle)(nil).FailurePull), varargs...)
}

// FailureStop mocks base method.
func (m *MockLifecycle) FailureStop(arg0 error, arg1 ...config.AgentConfig) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "FailureStop", varargs...)
}

// FailureStop indicates an expected call of FailureStop.
func (mr *MockLifecycleMockRecorder) FailureStop(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailureStop", reflect.TypeOf((*MockLifecycle)(nil).FailureStop), varargs...)
}

// FailureTooManyErrs mocks base method.
func (m *MockLifecycle) FailureTooManyErrs(arg0 error, arg1 ...config.AgentConfig) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "FailureTooManyErrs", varargs...)
}

// FailureTooManyErrs indicates an expected call of FailureTooManyErrs.
func (mr *MockLifecycleMockRecorder) FailureTooManyErrs(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailureTooManyErrs", reflect.TypeOf((*MockLifecycle)(nil).FailureTooManyErrs), varargs...)
}

// StatusActive mocks base method.
func (m *MockLifecycle) StatusActive(arg0 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StatusActive", arg0)
}

// StatusActive indicates an expected call of StatusActive.
func (mr *MockLifecycleMockRecorder) StatusActive(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusActive", reflect.TypeOf((*MockLifecycle)(nil).StatusActive), arg0)
}

// StatusAttached mocks base method.
func (m *MockLifecycle) StatusAttached(arg0 ...config.AgentConfig) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "StatusAttached", varargs...)
}

// StatusAttached indicates an expected call of StatusAttached.
func (mr *MockLifecycleMockRecorder) StatusAttached(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusAttached", reflect.TypeOf((*MockLifecycle)(nil).StatusAttached), arg0...)
}

// StatusInactive mocks base method.
func (m *MockLifecycle) StatusInactive(arg0 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StatusInactive", arg0)
}

// StatusInactive indicates an expected call of StatusInactive.
func (mr *MockLifecycleMockRecorder) StatusInactive(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusInactive", reflect.TypeOf((*MockLifecycle)(nil).StatusInactive), arg0)
}

// StatusInitialized mocks base method.
func (m *MockLifecycle) StatusInitialized(arg0 ...config.AgentConfig) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "StatusInitialized", varargs...)
}

// StatusInitialized indicates an expected call of StatusInitialized.
func (mr *MockLifecycleMockRecorder) StatusInitialized(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusInitialized", reflect.TypeOf((*MockLifecycle)(nil).StatusInitialized), arg0...)
}

// StatusRunning mocks base method.
func (m *MockLifecycle) StatusRunning(arg0 ...config.AgentConfig) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "StatusRunning", varargs...)
}

// StatusRunning indicates an expected call of StatusRunning.
func (mr *MockLifecycleMockRecorder) StatusRunning(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusRunning", reflect.TypeOf((*MockLifecycle)(nil).StatusRunning), arg0...)
}

// StatusStopping mocks base method.
func (m *MockLifecycle) StatusStopping(arg0 ...config.AgentConfig) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "StatusStopping", varargs...)
}

// StatusStopping indicates an expected call of StatusStopping.
func (mr *MockLifecycleMockRecorder) StatusStopping(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusStopping", reflect.TypeOf((*MockLifecycle)(nil).StatusStopping), arg0...)
}

// SystemError mocks base method.
func (m *MockLifecycle) SystemError(metricName string, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SystemError", metricName, err)
}

// SystemError indicates an expected call of SystemError.
func (mr *MockLifecycleMockRecorder) SystemError(metricName, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SystemError", reflect.TypeOf((*MockLifecycle)(nil).SystemError), metricName, err)
}

// SystemStatus mocks base method.
func (m *MockLifecycle) SystemStatus(metricName, details string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SystemStatus", metricName, details)
}

// SystemStatus indicates an expected call of SystemStatus.
func (mr *MockLifecycleMockRecorder) SystemStatus(metricName, details interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SystemStatus", reflect.TypeOf((*MockLifecycle)(nil).SystemStatus), metricName, details)
}
