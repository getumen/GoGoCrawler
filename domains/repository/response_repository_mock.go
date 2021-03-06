// Code generated by MockGen. DO NOT EDIT.
// Source: response_repository.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	models "github.com/getumen/gogo_crawler/domains/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockResponseRepository is a mock of ResponseRepository interface
type MockResponseRepository struct {
	ctrl     *gomock.Controller
	recorder *MockResponseRepositoryMockRecorder
}

// MockResponseRepositoryMockRecorder is the mock recorder for MockResponseRepository
type MockResponseRepositoryMockRecorder struct {
	mock *MockResponseRepository
}

// NewMockResponseRepository creates a new mock instance
func NewMockResponseRepository(ctrl *gomock.Controller) *MockResponseRepository {
	mock := &MockResponseRepository{ctrl: ctrl}
	mock.recorder = &MockResponseRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResponseRepository) EXPECT() *MockResponseRepositoryMockRecorder {
	return m.recorder
}

// Save mocks base method
func (m *MockResponseRepository) Save(ctx context.Context, response *models.Response) error {
	ret := m.ctrl.Call(m, "Save", ctx, response)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockResponseRepositoryMockRecorder) Save(ctx, response interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockResponseRepository)(nil).Save), ctx, response)
}

// IsExist mocks base method
func (m *MockResponseRepository) IsExist(ctx context.Context, response *models.Response) (bool, error) {
	ret := m.ctrl.Call(m, "IsExist", ctx, response)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsExist indicates an expected call of IsExist
func (mr *MockResponseRepositoryMockRecorder) IsExist(ctx, response interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExist", reflect.TypeOf((*MockResponseRepository)(nil).IsExist), ctx, response)
}
