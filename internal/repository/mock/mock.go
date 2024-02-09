// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/repository/interfaces.go
//
// Generated by this command:
//
//	mockgen -source=./internal/repository/interfaces.go -destination=./internal/repository/mock/mock.go
//

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	entity "demo/bank-linking-listener/internal/service/entity"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserRepository) Create(user entity.User) (*entity.User, entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(entity.Error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserRepositoryMockRecorder) Create(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepository)(nil).Create), user)
}

// GetByEmail mocks base method.
func (m *MockUserRepository) GetByEmail(email string) (*entity.User, entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", email)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(entity.Error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail.
func (mr *MockUserRepositoryMockRecorder) GetByEmail(email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockUserRepository)(nil).GetByEmail), email)
}