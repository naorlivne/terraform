// Code generated by MockGen. DO NOT EDIT.
// Source: ssh_key.go

// Package tfe is a generated GoMock package.
package tfe

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSSHKeys is a mock of SSHKeys interface.
type MockSSHKeys struct {
	ctrl     *gomock.Controller
	recorder *MockSSHKeysMockRecorder
}

// MockSSHKeysMockRecorder is the mock recorder for MockSSHKeys.
type MockSSHKeysMockRecorder struct {
	mock *MockSSHKeys
}

// NewMockSSHKeys creates a new mock instance.
func NewMockSSHKeys(ctrl *gomock.Controller) *MockSSHKeys {
	mock := &MockSSHKeys{ctrl: ctrl}
	mock.recorder = &MockSSHKeysMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSSHKeys) EXPECT() *MockSSHKeysMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSSHKeys) Create(ctx context.Context, organization string, options SSHKeyCreateOptions) (*SSHKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, organization, options)
	ret0, _ := ret[0].(*SSHKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockSSHKeysMockRecorder) Create(ctx, organization, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSSHKeys)(nil).Create), ctx, organization, options)
}

// Delete mocks base method.
func (m *MockSSHKeys) Delete(ctx context.Context, sshKeyID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, sshKeyID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockSSHKeysMockRecorder) Delete(ctx, sshKeyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSSHKeys)(nil).Delete), ctx, sshKeyID)
}

// List mocks base method.
func (m *MockSSHKeys) List(ctx context.Context, organization string, options SSHKeyListOptions) (*SSHKeyList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, organization, options)
	ret0, _ := ret[0].(*SSHKeyList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockSSHKeysMockRecorder) List(ctx, organization, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSSHKeys)(nil).List), ctx, organization, options)
}

// Read mocks base method.
func (m *MockSSHKeys) Read(ctx context.Context, sshKeyID string) (*SSHKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, sshKeyID)
	ret0, _ := ret[0].(*SSHKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockSSHKeysMockRecorder) Read(ctx, sshKeyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockSSHKeys)(nil).Read), ctx, sshKeyID)
}

// Update mocks base method.
func (m *MockSSHKeys) Update(ctx context.Context, sshKeyID string, options SSHKeyUpdateOptions) (*SSHKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, sshKeyID, options)
	ret0, _ := ret[0].(*SSHKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockSSHKeysMockRecorder) Update(ctx, sshKeyID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSSHKeys)(nil).Update), ctx, sshKeyID, options)
}
