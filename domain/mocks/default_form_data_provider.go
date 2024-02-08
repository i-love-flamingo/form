// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	web "flamingo.me/flamingo/v3/framework/web"
)

// DefaultFormDataProvider is an autogenerated mock type for the DefaultFormDataProvider type
type DefaultFormDataProvider struct {
	mock.Mock
}

type DefaultFormDataProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *DefaultFormDataProvider) EXPECT() *DefaultFormDataProvider_Expecter {
	return &DefaultFormDataProvider_Expecter{mock: &_m.Mock}
}

// GetFormData provides a mock function with given fields: ctx, req
func (_m *DefaultFormDataProvider) GetFormData(ctx context.Context, req *web.Request) (interface{}, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for GetFormData")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *web.Request) (interface{}, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *web.Request) interface{}); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *web.Request) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DefaultFormDataProvider_GetFormData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFormData'
type DefaultFormDataProvider_GetFormData_Call struct {
	*mock.Call
}

// GetFormData is a helper method to define mock.On call
//   - ctx context.Context
//   - req *web.Request
func (_e *DefaultFormDataProvider_Expecter) GetFormData(ctx interface{}, req interface{}) *DefaultFormDataProvider_GetFormData_Call {
	return &DefaultFormDataProvider_GetFormData_Call{Call: _e.mock.On("GetFormData", ctx, req)}
}

func (_c *DefaultFormDataProvider_GetFormData_Call) Run(run func(ctx context.Context, req *web.Request)) *DefaultFormDataProvider_GetFormData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*web.Request))
	})
	return _c
}

func (_c *DefaultFormDataProvider_GetFormData_Call) Return(_a0 interface{}, _a1 error) *DefaultFormDataProvider_GetFormData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DefaultFormDataProvider_GetFormData_Call) RunAndReturn(run func(context.Context, *web.Request) (interface{}, error)) *DefaultFormDataProvider_GetFormData_Call {
	_c.Call.Return(run)
	return _c
}

// NewDefaultFormDataProvider creates a new instance of DefaultFormDataProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDefaultFormDataProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *DefaultFormDataProvider {
	mock := &DefaultFormDataProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
