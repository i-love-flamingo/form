// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// FormService is an autogenerated mock type for the FormService type
type FormService struct {
	mock.Mock
}

type FormService_Expecter struct {
	mock *mock.Mock
}

func (_m *FormService) EXPECT() *FormService_Expecter {
	return &FormService_Expecter{mock: &_m.Mock}
}

// NewFormService creates a new instance of FormService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFormService(t interface {
	mock.TestingT
	Cleanup(func())
}) *FormService {
	mock := &FormService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
