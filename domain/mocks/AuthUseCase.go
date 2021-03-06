// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/questizen/core-system/domain"
	mock "github.com/stretchr/testify/mock"
)

// AuthUseCase is an autogenerated mock type for the AuthUseCase type
type AuthUseCase struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: _a0, _a1
func (_m *AuthUseCase) CreateUser(_a0 context.Context, _a1 *domain.AuthUser) (domain.User, error) {
	ret := _m.Called(_a0, _a1)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(context.Context, *domain.AuthUser) domain.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.AuthUser) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: ctx
func (_m *AuthUseCase) GetUser(ctx context.Context) ([]domain.User, error) {
	ret := _m.Called(ctx)

	var r0 []domain.User
	if rf, ok := ret.Get(0).(func(context.Context) []domain.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAuthUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthUseCase creates a new instance of AuthUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthUseCase(t mockConstructorTestingTNewAuthUseCase) *AuthUseCase {
	mock := &AuthUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
