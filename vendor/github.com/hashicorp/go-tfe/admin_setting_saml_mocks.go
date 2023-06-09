// Code generated by MockGen. DO NOT EDIT.
// Source: admin_setting_saml.go

// Package tfe is a generated GoMock package.
package tfe

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSAMLSettings is a mock of SAMLSettings interface.
type MockSAMLSettings struct {
	ctrl     *gomock.Controller
	recorder *MockSAMLSettingsMockRecorder
}

// MockSAMLSettingsMockRecorder is the mock recorder for MockSAMLSettings.
type MockSAMLSettingsMockRecorder struct {
	mock *MockSAMLSettings
}

// NewMockSAMLSettings creates a new mock instance.
func NewMockSAMLSettings(ctrl *gomock.Controller) *MockSAMLSettings {
	mock := &MockSAMLSettings{ctrl: ctrl}
	mock.recorder = &MockSAMLSettingsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSAMLSettings) EXPECT() *MockSAMLSettingsMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockSAMLSettings) Read(ctx context.Context) (*AdminSAMLSetting, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx)
	ret0, _ := ret[0].(*AdminSAMLSetting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockSAMLSettingsMockRecorder) Read(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockSAMLSettings)(nil).Read), ctx)
}

// RevokeIdpCert mocks base method.
func (m *MockSAMLSettings) RevokeIdpCert(ctx context.Context) (*AdminSAMLSetting, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeIdpCert", ctx)
	ret0, _ := ret[0].(*AdminSAMLSetting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevokeIdpCert indicates an expected call of RevokeIdpCert.
func (mr *MockSAMLSettingsMockRecorder) RevokeIdpCert(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeIdpCert", reflect.TypeOf((*MockSAMLSettings)(nil).RevokeIdpCert), ctx)
}

// Update mocks base method.
func (m *MockSAMLSettings) Update(ctx context.Context, options AdminSAMLSettingsUpdateOptions) (*AdminSAMLSetting, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, options)
	ret0, _ := ret[0].(*AdminSAMLSetting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockSAMLSettingsMockRecorder) Update(ctx, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSAMLSettings)(nil).Update), ctx, options)
}
