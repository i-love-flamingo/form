// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	domain "flamingo.me/form/domain"
	mock "github.com/stretchr/testify/mock"
)

// FormDataEncoderFactory is an autogenerated mock type for the FormDataEncoderFactory type
type FormDataEncoderFactory struct {
	mock.Mock
}

type FormDataEncoderFactory_Expecter struct {
	mock *mock.Mock
}

func (_m *FormDataEncoderFactory) EXPECT() *FormDataEncoderFactory_Expecter {
	return &FormDataEncoderFactory_Expecter{mock: &_m.Mock}
}

// CreateByNamedEncoder provides a mock function with given fields: name
func (_m *FormDataEncoderFactory) CreateByNamedEncoder(name string) domain.FormDataEncoder {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for CreateByNamedEncoder")
	}

	var r0 domain.FormDataEncoder
	if rf, ok := ret.Get(0).(func(string) domain.FormDataEncoder); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.FormDataEncoder)
		}
	}

	return r0
}

// FormDataEncoderFactory_CreateByNamedEncoder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateByNamedEncoder'
type FormDataEncoderFactory_CreateByNamedEncoder_Call struct {
	*mock.Call
}

// CreateByNamedEncoder is a helper method to define mock.On call
//   - name string
func (_e *FormDataEncoderFactory_Expecter) CreateByNamedEncoder(name interface{}) *FormDataEncoderFactory_CreateByNamedEncoder_Call {
	return &FormDataEncoderFactory_CreateByNamedEncoder_Call{Call: _e.mock.On("CreateByNamedEncoder", name)}
}

func (_c *FormDataEncoderFactory_CreateByNamedEncoder_Call) Run(run func(name string)) *FormDataEncoderFactory_CreateByNamedEncoder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *FormDataEncoderFactory_CreateByNamedEncoder_Call) Return(_a0 domain.FormDataEncoder) *FormDataEncoderFactory_CreateByNamedEncoder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FormDataEncoderFactory_CreateByNamedEncoder_Call) RunAndReturn(run func(string) domain.FormDataEncoder) *FormDataEncoderFactory_CreateByNamedEncoder_Call {
	_c.Call.Return(run)
	return _c
}

// CreateWithFormService provides a mock function with given fields: formService
func (_m *FormDataEncoderFactory) CreateWithFormService(formService domain.FormService) domain.FormDataEncoder {
	ret := _m.Called(formService)

	if len(ret) == 0 {
		panic("no return value specified for CreateWithFormService")
	}

	var r0 domain.FormDataEncoder
	if rf, ok := ret.Get(0).(func(domain.FormService) domain.FormDataEncoder); ok {
		r0 = rf(formService)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.FormDataEncoder)
		}
	}

	return r0
}

// FormDataEncoderFactory_CreateWithFormService_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateWithFormService'
type FormDataEncoderFactory_CreateWithFormService_Call struct {
	*mock.Call
}

// CreateWithFormService is a helper method to define mock.On call
//   - formService domain.FormService
func (_e *FormDataEncoderFactory_Expecter) CreateWithFormService(formService interface{}) *FormDataEncoderFactory_CreateWithFormService_Call {
	return &FormDataEncoderFactory_CreateWithFormService_Call{Call: _e.mock.On("CreateWithFormService", formService)}
}

func (_c *FormDataEncoderFactory_CreateWithFormService_Call) Run(run func(formService domain.FormService)) *FormDataEncoderFactory_CreateWithFormService_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(domain.FormService))
	})
	return _c
}

func (_c *FormDataEncoderFactory_CreateWithFormService_Call) Return(_a0 domain.FormDataEncoder) *FormDataEncoderFactory_CreateWithFormService_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FormDataEncoderFactory_CreateWithFormService_Call) RunAndReturn(run func(domain.FormService) domain.FormDataEncoder) *FormDataEncoderFactory_CreateWithFormService_Call {
	_c.Call.Return(run)
	return _c
}

// NewFormDataEncoderFactory creates a new instance of FormDataEncoderFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFormDataEncoderFactory(t interface {
	mock.TestingT
	Cleanup(func())
}) *FormDataEncoderFactory {
	mock := &FormDataEncoderFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
