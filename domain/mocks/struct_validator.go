// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	validator "github.com/go-playground/validator/v10"
)

// StructValidator is an autogenerated mock type for the StructValidator type
type StructValidator struct {
	mock.Mock
}

type StructValidator_Expecter struct {
	mock *mock.Mock
}

func (_m *StructValidator) EXPECT() *StructValidator_Expecter {
	return &StructValidator_Expecter{mock: &_m.Mock}
}

// StructType provides a mock function with given fields:
func (_m *StructValidator) StructType() interface{} {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for StructType")
	}

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// StructValidator_StructType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StructType'
type StructValidator_StructType_Call struct {
	*mock.Call
}

// StructType is a helper method to define mock.On call
func (_e *StructValidator_Expecter) StructType() *StructValidator_StructType_Call {
	return &StructValidator_StructType_Call{Call: _e.mock.On("StructType")}
}

func (_c *StructValidator_StructType_Call) Run(run func()) *StructValidator_StructType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *StructValidator_StructType_Call) Return(_a0 interface{}) *StructValidator_StructType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StructValidator_StructType_Call) RunAndReturn(run func() interface{}) *StructValidator_StructType_Call {
	_c.Call.Return(run)
	return _c
}

// ValidateStruct provides a mock function with given fields: ctx, sl
func (_m *StructValidator) ValidateStruct(ctx context.Context, sl validator.StructLevel) {
	_m.Called(ctx, sl)
}

// StructValidator_ValidateStruct_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValidateStruct'
type StructValidator_ValidateStruct_Call struct {
	*mock.Call
}

// ValidateStruct is a helper method to define mock.On call
//   - ctx context.Context
//   - sl validator.StructLevel
func (_e *StructValidator_Expecter) ValidateStruct(ctx interface{}, sl interface{}) *StructValidator_ValidateStruct_Call {
	return &StructValidator_ValidateStruct_Call{Call: _e.mock.On("ValidateStruct", ctx, sl)}
}

func (_c *StructValidator_ValidateStruct_Call) Run(run func(ctx context.Context, sl validator.StructLevel)) *StructValidator_ValidateStruct_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(validator.StructLevel))
	})
	return _c
}

func (_c *StructValidator_ValidateStruct_Call) Return() *StructValidator_ValidateStruct_Call {
	_c.Call.Return()
	return _c
}

func (_c *StructValidator_ValidateStruct_Call) RunAndReturn(run func(context.Context, validator.StructLevel)) *StructValidator_ValidateStruct_Call {
	_c.Call.Return(run)
	return _c
}

// NewStructValidator creates a new instance of StructValidator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStructValidator(t interface {
	mock.TestingT
	Cleanup(func())
}) *StructValidator {
	mock := &StructValidator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
