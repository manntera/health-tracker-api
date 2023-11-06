// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -source interface.go -destination=./mock/interfaceMock.go
//

// Package mock_userRepository is a generated GoMock package.
package mock_userRepository

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	userRepository "manntera.com/health-tracker-api/pkg/repository/userRepository"
)

// MockIUserRepository is a mock of IUserRepository interface.
type MockIUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepositoryMockRecorder
}

// MockIUserRepositoryMockRecorder is the mock recorder for MockIUserRepository.
type MockIUserRepositoryMockRecorder struct {
	mock *MockIUserRepository
}

// NewMockIUserRepository creates a new mock instance.
func NewMockIUserRepository(ctrl *gomock.Controller) *MockIUserRepository {
	mock := &MockIUserRepository{ctrl: ctrl}
	mock.recorder = &MockIUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRepository) EXPECT() *MockIUserRepositoryMockRecorder {
	return m.recorder
}

// AddData mocks base method.
func (m *MockIUserRepository) AddData(ctx context.Context, userData *userRepository.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddData", ctx, userData)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddData indicates an expected call of AddData.
func (mr *MockIUserRepositoryMockRecorder) AddData(ctx, userData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddData", reflect.TypeOf((*MockIUserRepository)(nil).AddData), ctx, userData)
}

// DeleteData mocks base method.
func (m *MockIUserRepository) DeleteData(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteData", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteData indicates an expected call of DeleteData.
func (mr *MockIUserRepositoryMockRecorder) DeleteData(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteData", reflect.TypeOf((*MockIUserRepository)(nil).DeleteData), ctx, id)
}

// GetData mocks base method.
func (m *MockIUserRepository) GetData(ctx context.Context, id string) (*userRepository.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetData", ctx, id)
	ret0, _ := ret[0].(*userRepository.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetData indicates an expected call of GetData.
func (mr *MockIUserRepositoryMockRecorder) GetData(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetData", reflect.TypeOf((*MockIUserRepository)(nil).GetData), ctx, id)
}
