// Code generated by MockGen. DO NOT EDIT.
// Source: shortener.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	models "ozon-fintech/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockServices is a mock of Services interface.
type MockServices struct {
	ctrl     *gomock.Controller
	recorder *MockServicesMockRecorder
}

// MockServicesMockRecorder is the mock recorder for MockServices.
type MockServicesMockRecorder struct {
	mock *MockServices
}

// NewMockServices creates a new mock instance.
func NewMockServices(ctrl *gomock.Controller) *MockServices {
	mock := &MockServices{ctrl: ctrl}
	mock.recorder = &MockServicesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServices) EXPECT() *MockServicesMockRecorder {
	return m.recorder
}

// GetFullURL mocks base method.
func (m *MockServices) GetFullURL(shortURL string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFullURL", shortURL)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFullURL indicates an expected call of GetFullURL.
func (mr *MockServicesMockRecorder) GetFullURL(shortURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFullURL", reflect.TypeOf((*MockServices)(nil).GetFullURL), shortURL)
}

// LoadShortURL mocks base method.
func (m *MockServices) LoadShortURL(link models.Link) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadShortURL", link)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadShortURL indicates an expected call of LoadShortURL.
func (mr *MockServicesMockRecorder) LoadShortURL(link interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadShortURL", reflect.TypeOf((*MockServices)(nil).LoadShortURL), link)
}