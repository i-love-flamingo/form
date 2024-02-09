// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "flamingo.me/form/domain"
	mock "github.com/stretchr/testify/mock"

	web "flamingo.me/flamingo/v3/framework/web"
)

// FormDataValidator is an autogenerated mock type for the FormDataValidator type
type FormDataValidator struct {
	mock.Mock
}

type FormDataValidator_Expecter struct {
	mock *mock.Mock
}

func (_m *FormDataValidator) EXPECT() *FormDataValidator_Expecter {
	return &FormDataValidator_Expecter{mock: &_m.Mock}
}

// Validate provides a mock function with given fields: ctx, req, validatorProvider, formData
func (_m *FormDataValidator) Validate(ctx context.Context, req *web.Request, validatorProvider domain.ValidatorProvider, formData interface{}) (*domain.ValidationInfo, error) {
	ret := _m.Called(ctx, req, validatorProvider, formData)

	if len(ret) == 0 {
		panic("no return value specified for Validate")
	}

	var r0 *domain.ValidationInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *web.Request, domain.ValidatorProvider, interface{}) (*domain.ValidationInfo, error)); ok {
		return rf(ctx, req, validatorProvider, formData)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *web.Request, domain.ValidatorProvider, interface{}) *domain.ValidationInfo); ok {
		r0 = rf(ctx, req, validatorProvider, formData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.ValidationInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *web.Request, domain.ValidatorProvider, interface{}) error); ok {
		r1 = rf(ctx, req, validatorProvider, formData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FormDataValidator_Validate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Validate'
type FormDataValidator_Validate_Call struct {
	*mock.Call
}

// Validate is a helper method to define mock.On call
//   - ctx context.Context
//   - req *web.Request
//   - validatorProvider domain.ValidatorProvider
//   - formData interface{}
func (_e *FormDataValidator_Expecter) Validate(ctx interface{}, req interface{}, validatorProvider interface{}, formData interface{}) *FormDataValidator_Validate_Call {
	return &FormDataValidator_Validate_Call{Call: _e.mock.On("Validate", ctx, req, validatorProvider, formData)}
}

func (_c *FormDataValidator_Validate_Call) Run(run func(ctx context.Context, req *web.Request, validatorProvider domain.ValidatorProvider, formData interface{})) *FormDataValidator_Validate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*web.Request), args[2].(domain.ValidatorProvider), args[3].(interface{}))
	})
	return _c
}

func (_c *FormDataValidator_Validate_Call) Return(_a0 *domain.ValidationInfo, _a1 error) *FormDataValidator_Validate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FormDataValidator_Validate_Call) RunAndReturn(run func(context.Context, *web.Request, domain.ValidatorProvider, interface{}) (*domain.ValidationInfo, error)) *FormDataValidator_Validate_Call {
	_c.Call.Return(run)
	return _c
}

// NewFormDataValidator creates a new instance of FormDataValidator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFormDataValidator(t interface {
	mock.TestingT
	Cleanup(func())
}) *FormDataValidator {
	mock := &FormDataValidator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
