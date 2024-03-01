// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	authz "github.com/chainloop-dev/chainloop/app/controlplane/internal/authz"

	mock "github.com/stretchr/testify/mock"
)

// Enforcer is an autogenerated mock type for the Enforcer type
type Enforcer struct {
	mock.Mock
}

// Enforce provides a mock function with given fields: sub, p
func (_m *Enforcer) Enforce(sub string, p *authz.Policy) (bool, error) {
	ret := _m.Called(sub, p)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *authz.Policy) (bool, error)); ok {
		return rf(sub, p)
	}
	if rf, ok := ret.Get(0).(func(string, *authz.Policy) bool); ok {
		r0 = rf(sub, p)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, *authz.Policy) error); ok {
		r1 = rf(sub, p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewEnforcer interface {
	mock.TestingT
	Cleanup(func())
}

// NewEnforcer creates a new instance of Enforcer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEnforcer(t mockConstructorTestingTNewEnforcer) *Enforcer {
	mock := &Enforcer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
