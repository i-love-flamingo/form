// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	url "net/url"
)

// DefaultFormDataEncoder is an autogenerated mock type for the DefaultFormDataEncoder type
type DefaultFormDataEncoder struct {
	mock.Mock
}

type DefaultFormDataEncoder_Expecter struct {
	mock *mock.Mock
}

func (_m *DefaultFormDataEncoder) EXPECT() *DefaultFormDataEncoder_Expecter {
	return &DefaultFormDataEncoder_Expecter{mock: &_m.Mock}
}

// Encode provides a mock function with given fields: ctx, formData
func (_m *DefaultFormDataEncoder) Encode(ctx context.Context, formData interface{}) (url.Values, error) {
	ret := _m.Called(ctx, formData)

	if len(ret) == 0 {
		panic("no return value specified for Encode")
	}

	var r0 url.Values
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) (url.Values, error)); ok {
		return rf(ctx, formData)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) url.Values); ok {
		r0 = rf(ctx, formData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(url.Values)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, formData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DefaultFormDataEncoder_Encode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Encode'
type DefaultFormDataEncoder_Encode_Call struct {
	*mock.Call
}

// Encode is a helper method to define mock.On call
//   - ctx context.Context
//   - formData interface{}
func (_e *DefaultFormDataEncoder_Expecter) Encode(ctx interface{}, formData interface{}) *DefaultFormDataEncoder_Encode_Call {
	return &DefaultFormDataEncoder_Encode_Call{Call: _e.mock.On("Encode", ctx, formData)}
}

func (_c *DefaultFormDataEncoder_Encode_Call) Run(run func(ctx context.Context, formData interface{})) *DefaultFormDataEncoder_Encode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}))
	})
	return _c
}

func (_c *DefaultFormDataEncoder_Encode_Call) Return(_a0 url.Values, _a1 error) *DefaultFormDataEncoder_Encode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DefaultFormDataEncoder_Encode_Call) RunAndReturn(run func(context.Context, interface{}) (url.Values, error)) *DefaultFormDataEncoder_Encode_Call {
	_c.Call.Return(run)
	return _c
}

// NewDefaultFormDataEncoder creates a new instance of DefaultFormDataEncoder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDefaultFormDataEncoder(t interface {
	mock.TestingT
	Cleanup(func())
}) *DefaultFormDataEncoder {
	mock := &DefaultFormDataEncoder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
