// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/grafana/grafana/pkg/services/live/features (interfaces: LiveMessageStore)

// Package features is a generated GoMock package.
package features

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	models "github.com/grafana/grafana/pkg/models"
)

// MockLiveMessageStore is a mock of LiveMessageStore interface.
type MockLiveMessageStore struct {
	ctrl     *gomock.Controller
	recorder *MockLiveMessageStoreMockRecorder
}

// MockLiveMessageStoreMockRecorder is the mock recorder for MockLiveMessageStore.
type MockLiveMessageStoreMockRecorder struct {
	mock *MockLiveMessageStore
}

// NewMockLiveMessageStore creates a new mock instance.
func NewMockLiveMessageStore(ctrl *gomock.Controller) *MockLiveMessageStore {
	mock := &MockLiveMessageStore{ctrl: ctrl}
	mock.recorder = &MockLiveMessageStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLiveMessageStore) EXPECT() *MockLiveMessageStoreMockRecorder {
	return m.recorder
}

// GetLiveMessage mocks base method.
func (m *MockLiveMessageStore) GetLiveMessage(arg0 *models.GetLiveMessageQuery) (models.LiveMessage, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLiveMessage", arg0)
	ret0, _ := ret[0].(models.LiveMessage)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetLiveMessage indicates an expected call of GetLiveMessage.
func (mr *MockLiveMessageStoreMockRecorder) GetLiveMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLiveMessage", reflect.TypeOf((*MockLiveMessageStore)(nil).GetLiveMessage), arg0)
}

// SaveLiveMessage mocks base method.
func (m *MockLiveMessageStore) SaveLiveMessage(arg0 *models.SaveLiveMessageQuery) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveLiveMessage", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveLiveMessage indicates an expected call of SaveLiveMessage.
func (mr *MockLiveMessageStoreMockRecorder) SaveLiveMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveLiveMessage", reflect.TypeOf((*MockLiveMessageStore)(nil).SaveLiveMessage), arg0)
}
