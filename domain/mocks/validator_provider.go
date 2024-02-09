// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "flamingo.me/form/domain"
	mock "github.com/stretchr/testify/mock"

	validator "github.com/go-playground/validator/v10"

	web "flamingo.me/flamingo/v3/framework/web"
)

// ValidatorProvider is an autogenerated mock type for the ValidatorProvider type
type ValidatorProvider struct {
	mock.Mock
}

type ValidatorProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *ValidatorProvider) EXPECT() *ValidatorProvider_Expecter {
	return &ValidatorProvider_Expecter{mock: &_m.Mock}
}

// ErrorsToValidationInfo provides a mock function with given fields: err
func (_m *ValidatorProvider) ErrorsToValidationInfo(err error) domain.ValidationInfo {
	ret := _m.Called(err)

	if len(ret) == 0 {
		panic("no return value specified for ErrorsToValidationInfo")
	}

	var r0 domain.ValidationInfo
	if rf, ok := ret.Get(0).(func(error) domain.ValidationInfo); ok {
		r0 = rf(err)
	} else {
		r0 = ret.Get(0).(domain.ValidationInfo)
	}

	return r0
}

// ValidatorProvider_ErrorsToValidationInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ErrorsToValidationInfo'
type ValidatorProvider_ErrorsToValidationInfo_Call struct {
	*mock.Call
}

// ErrorsToValidationInfo is a helper method to define mock.On call
//   - err error
func (_e *ValidatorProvider_Expecter) ErrorsToValidationInfo(err interface{}) *ValidatorProvider_ErrorsToValidationInfo_Call {
	return &ValidatorProvider_ErrorsToValidationInfo_Call{Call: _e.mock.On("ErrorsToValidationInfo", err)}
}

func (_c *ValidatorProvider_ErrorsToValidationInfo_Call) Run(run func(err error)) *ValidatorProvider_ErrorsToValidationInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error))
	})
	return _c
}

func (_c *ValidatorProvider_ErrorsToValidationInfo_Call) Return(_a0 domain.ValidationInfo) *ValidatorProvider_ErrorsToValidationInfo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ValidatorProvider_ErrorsToValidationInfo_Call) RunAndReturn(run func(error) domain.ValidationInfo) *ValidatorProvider_ErrorsToValidationInfo_Call {
	_c.Call.Return(run)
	return _c
}

// GetValidator provides a mock function with given fields:
func (_m *ValidatorProvider) GetValidator() *validator.Validate {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetValidator")
	}

	var r0 *validator.Validate
	if rf, ok := ret.Get(0).(func() *validator.Validate); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*validator.Validate)
		}
	}

	return r0
}

// ValidatorProvider_GetValidator_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetValidator'
type ValidatorProvider_GetValidator_Call struct {
	*mock.Call
}

// GetValidator is a helper method to define mock.On call
func (_e *ValidatorProvider_Expecter) GetValidator() *ValidatorProvider_GetValidator_Call {
	return &ValidatorProvider_GetValidator_Call{Call: _e.mock.On("GetValidator")}
}

func (_c *ValidatorProvider_GetValidator_Call) Run(run func()) *ValidatorProvider_GetValidator_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ValidatorProvider_GetValidator_Call) Return(_a0 *validator.Validate) *ValidatorProvider_GetValidator_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ValidatorProvider_GetValidator_Call) RunAndReturn(run func() *validator.Validate) *ValidatorProvider_GetValidator_Call {
	_c.Call.Return(run)
	return _c
}

// Validate provides a mock function with given fields: ctx, req, value
func (_m *ValidatorProvider) Validate(ctx context.Context, req *web.Request, value interface{}) domain.ValidationInfo {
	ret := _m.Called(ctx, req, value)

	if len(ret) == 0 {
		panic("no return value specified for Validate")
	}

	var r0 domain.ValidationInfo
	if rf, ok := ret.Get(0).(func(context.Context, *web.Request, interface{}) domain.ValidationInfo); ok {
		r0 = rf(ctx, req, value)
	} else {
		r0 = ret.Get(0).(domain.ValidationInfo)
	}

	return r0
}

// ValidatorProvider_Validate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Validate'
type ValidatorProvider_Validate_Call struct {
	*mock.Call
}

// Validate is a helper method to define mock.On call
//   - ctx context.Context
//   - req *web.Request
//   - value interface{}
func (_e *ValidatorProvider_Expecter) Validate(ctx interface{}, req interface{}, value interface{}) *ValidatorProvider_Validate_Call {
	return &ValidatorProvider_Validate_Call{Call: _e.mock.On("Validate", ctx, req, value)}
}

func (_c *ValidatorProvider_Validate_Call) Run(run func(ctx context.Context, req *web.Request, value interface{})) *ValidatorProvider_Validate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*web.Request), args[2].(interface{}))
	})
	return _c
}

func (_c *ValidatorProvider_Validate_Call) Return(_a0 domain.ValidationInfo) *ValidatorProvider_Validate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ValidatorProvider_Validate_Call) RunAndReturn(run func(context.Context, *web.Request, interface{}) domain.ValidationInfo) *ValidatorProvider_Validate_Call {
	_c.Call.Return(run)
	return _c
}

// NewValidatorProvider creates a new instance of ValidatorProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewValidatorProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *ValidatorProvider {
	mock := &ValidatorProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
