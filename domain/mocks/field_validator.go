// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	validator "gopkg.in/go-playground/validator.v9"
)

// FieldValidator is an autogenerated mock type for the FieldValidator type
type FieldValidator struct {
	mock.Mock
}

type FieldValidator_Expecter struct {
	mock *mock.Mock
}

func (_m *FieldValidator) EXPECT() *FieldValidator_Expecter {
	return &FieldValidator_Expecter{mock: &_m.Mock}
}

// ValidateField provides a mock function with given fields: ctx, fl
func (_m *FieldValidator) ValidateField(ctx context.Context, fl validator.FieldLevel) bool {
	ret := _m.Called(ctx, fl)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, validator.FieldLevel) bool); ok {
		r0 = rf(ctx, fl)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// FieldValidator_ValidateField_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValidateField'
type FieldValidator_ValidateField_Call struct {
	*mock.Call
}

// ValidateField is a helper method to define mock.On call
//   - ctx context.Context
//   - fl validator.FieldLevel
func (_e *FieldValidator_Expecter) ValidateField(ctx interface{}, fl interface{}) *FieldValidator_ValidateField_Call {
	return &FieldValidator_ValidateField_Call{Call: _e.mock.On("ValidateField", ctx, fl)}
}

func (_c *FieldValidator_ValidateField_Call) Run(run func(ctx context.Context, fl validator.FieldLevel)) *FieldValidator_ValidateField_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(validator.FieldLevel))
	})
	return _c
}

func (_c *FieldValidator_ValidateField_Call) Return(_a0 bool) *FieldValidator_ValidateField_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FieldValidator_ValidateField_Call) RunAndReturn(run func(context.Context, validator.FieldLevel) bool) *FieldValidator_ValidateField_Call {
	_c.Call.Return(run)
	return _c
}

// ValidatorName provides a mock function with given fields:
func (_m *FieldValidator) ValidatorName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// FieldValidator_ValidatorName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValidatorName'
type FieldValidator_ValidatorName_Call struct {
	*mock.Call
}

// ValidatorName is a helper method to define mock.On call
func (_e *FieldValidator_Expecter) ValidatorName() *FieldValidator_ValidatorName_Call {
	return &FieldValidator_ValidatorName_Call{Call: _e.mock.On("ValidatorName")}
}

func (_c *FieldValidator_ValidatorName_Call) Run(run func()) *FieldValidator_ValidatorName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FieldValidator_ValidatorName_Call) Return(_a0 string) *FieldValidator_ValidatorName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FieldValidator_ValidatorName_Call) RunAndReturn(run func() string) *FieldValidator_ValidatorName_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewFieldValidator interface {
	mock.TestingT
	Cleanup(func())
}

// NewFieldValidator creates a new instance of FieldValidator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFieldValidator(t mockConstructorTestingTNewFieldValidator) *FieldValidator {
	mock := &FieldValidator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
