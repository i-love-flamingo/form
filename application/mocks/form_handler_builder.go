// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	application "flamingo.me/form/application"
	domain "flamingo.me/form/domain"

	mock "github.com/stretchr/testify/mock"
)

// FormHandlerBuilder is an autogenerated mock type for the FormHandlerBuilder type
type FormHandlerBuilder struct {
	mock.Mock
}

type FormHandlerBuilder_Expecter struct {
	mock *mock.Mock
}

func (_m *FormHandlerBuilder) EXPECT() *FormHandlerBuilder_Expecter {
	return &FormHandlerBuilder_Expecter{mock: &_m.Mock}
}

// AddFormExtension provides a mock function with given fields: formExtension
func (_m *FormHandlerBuilder) AddFormExtension(formExtension domain.FormExtension) error {
	ret := _m.Called(formExtension)

	if len(ret) == 0 {
		panic("no return value specified for AddFormExtension")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.FormExtension) error); ok {
		r0 = rf(formExtension)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FormHandlerBuilder_AddFormExtension_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddFormExtension'
type FormHandlerBuilder_AddFormExtension_Call struct {
	*mock.Call
}

// AddFormExtension is a helper method to define mock.On call
//   - formExtension domain.FormExtension
func (_e *FormHandlerBuilder_Expecter) AddFormExtension(formExtension interface{}) *FormHandlerBuilder_AddFormExtension_Call {
	return &FormHandlerBuilder_AddFormExtension_Call{Call: _e.mock.On("AddFormExtension", formExtension)}
}

func (_c *FormHandlerBuilder_AddFormExtension_Call) Run(run func(formExtension domain.FormExtension)) *FormHandlerBuilder_AddFormExtension_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(domain.FormExtension))
	})
	return _c
}

func (_c *FormHandlerBuilder_AddFormExtension_Call) Return(_a0 error) *FormHandlerBuilder_AddFormExtension_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FormHandlerBuilder_AddFormExtension_Call) RunAndReturn(run func(domain.FormExtension) error) *FormHandlerBuilder_AddFormExtension_Call {
	_c.Call.Return(run)
	return _c
}

// AddNamedFormExtension provides a mock function with given fields: name
func (_m *FormHandlerBuilder) AddNamedFormExtension(name string) error {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for AddNamedFormExtension")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FormHandlerBuilder_AddNamedFormExtension_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddNamedFormExtension'
type FormHandlerBuilder_AddNamedFormExtension_Call struct {
	*mock.Call
}

// AddNamedFormExtension is a helper method to define mock.On call
//   - name string
func (_e *FormHandlerBuilder_Expecter) AddNamedFormExtension(name interface{}) *FormHandlerBuilder_AddNamedFormExtension_Call {
	return &FormHandlerBuilder_AddNamedFormExtension_Call{Call: _e.mock.On("AddNamedFormExtension", name)}
}

func (_c *FormHandlerBuilder_AddNamedFormExtension_Call) Run(run func(name string)) *FormHandlerBuilder_AddNamedFormExtension_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *FormHandlerBuilder_AddNamedFormExtension_Call) Return(_a0 error) *FormHandlerBuilder_AddNamedFormExtension_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FormHandlerBuilder_AddNamedFormExtension_Call) RunAndReturn(run func(string) error) *FormHandlerBuilder_AddNamedFormExtension_Call {
	_c.Call.Return(run)
	return _c
}

// Build provides a mock function with no fields
func (_m *FormHandlerBuilder) Build() domain.FormHandler {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Build")
	}

	var r0 domain.FormHandler
	if rf, ok := ret.Get(0).(func() domain.FormHandler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.FormHandler)
		}
	}

	return r0
}

// FormHandlerBuilder_Build_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Build'
type FormHandlerBuilder_Build_Call struct {
	*mock.Call
}

// Build is a helper method to define mock.On call
func (_e *FormHandlerBuilder_Expecter) Build() *FormHandlerBuilder_Build_Call {
	return &FormHandlerBuilder_Build_Call{Call: _e.mock.On("Build")}
}

func (_c *FormHandlerBuilder_Build_Call) Run(run func()) *FormHandlerBuilder_Build_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FormHandlerBuilder_Build_Call) Return(_a0 domain.FormHandler) *FormHandlerBuilder_Build_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FormHandlerBuilder_Build_Call) RunAndReturn(run func() domain.FormHandler) *FormHandlerBuilder_Build_Call {
	_c.Call.Return(run)
	return _c
}

// Must provides a mock function with given fields: err
func (_m *FormHandlerBuilder) Must(err error) application.FormHandlerBuilder {
	ret := _m.Called(err)

	if len(ret) == 0 {
		panic("no return value specified for Must")
	}

	var r0 application.FormHandlerBuilder
	if rf, ok := ret.Get(0).(func(error) application.FormHandlerBuilder); ok {
		r0 = rf(err)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(application.FormHandlerBuilder)
		}
	}

	return r0
}

// FormHandlerBuilder_Must_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Must'
type FormHandlerBuilder_Must_Call struct {
	*mock.Call
}

// Must is a helper method to define mock.On call
//   - err error
func (_e *FormHandlerBuilder_Expecter) Must(err interface{}) *FormHandlerBuilder_Must_Call {
	return &FormHandlerBuilder_Must_Call{Call: _e.mock.On("Must", err)}
}

func (_c *FormHandlerBuilder_Must_Call) Run(run func(err error)) *FormHandlerBuilder_Must_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error))
	})
	return _c
}

func (_c *FormHandlerBuilder_Must_Call) Return(_a0 application.FormHandlerBuilder) *FormHandlerBuilder_Must_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FormHandlerBuilder_Must_Call) RunAndReturn(run func(error) application.FormHandlerBuilder) *FormHandlerBuilder_Must_Call {
	_c.Call.Return(run)
	return _c
}

// SetFormDataDecoder provides a mock function with given fields: formDataDecoder
func (_m *FormHandlerBuilder) SetFormDataDecoder(formDataDecoder domain.FormDataDecoder) application.FormHandlerBuilder {
	ret := _m.Called(formDataDecoder)

	if len(ret) == 0 {
		panic("no return value specified for SetFormDataDecoder")
	}

	var r0 application.FormHandlerBuilder
	if rf, ok := ret.Get(0).(func(domain.FormDataDecoder) application.FormHandlerBuilder); ok {
		r0 = rf(formDataDecoder)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(application.FormHandlerBuilder)
		}
	}

	return r0
}

// FormHandlerBuilder_SetFormDataDecoder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetFormDataDecoder'
type FormHandlerBuilder_SetFormDataDecoder_Call struct {
	*mock.Call
}

// SetFormDataDecoder is a helper method to define mock.On call
//   - formDataDecoder domain.FormDataDecoder
func (_e *FormHandlerBuilder_Expecter) SetFormDataDecoder(formDataDecoder interface{}) *FormHandlerBuilder_SetFormDataDecoder_Call {
	return &FormHandlerBuilder_SetFormDataDecoder_Call{Call: _e.mock.On("SetFormDataDecoder", formDataDecoder)}
}

func (_c *FormHandlerBuilder_SetFormDataDecoder_Call) Run(run func(formDataDecoder domain.FormDataDecoder)) *FormHandlerBuilder_SetFormDataDecoder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(domain.FormDataDecoder))
	})
	return _c
}

func (_c *FormHandlerBuilder_SetFormDataDecoder_Call) Return(_a0 application.FormHandlerBuilder) *FormHandlerBuilder_SetFormDataDecoder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FormHandlerBuilder_SetFormDataDecoder_Call) RunAndReturn(run func(domain.FormDataDecoder) application.FormHandlerBuilder) *FormHandlerBuilder_SetFormDataDecoder_Call {
	_c.Call.Return(run)
	return _c
}

// SetFormDataProvider provides a mock function with given fields: formDataProvider
func (_m *FormHandlerBuilder) SetFormDataProvider(formDataProvider domain.FormDataProvider) application.FormHandlerBuilder {
	ret := _m.Called(formDataProvider)

	if len(ret) == 0 {
		panic("no return value specified for SetFormDataProvider")
	}

	var r0 application.FormHandlerBuilder
	if rf, ok := ret.Get(0).(func(domain.FormDataProvider) application.FormHandlerBuilder); ok {
		r0 = rf(formDataProvider)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(application.FormHandlerBuilder)
		}
	}

	return r0
}

// FormHandlerBuilder_SetFormDataProvider_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetFormDataProvider'
type FormHandlerBuilder_SetFormDataProvider_Call struct {
	*mock.Call
}

// SetFormDataProvider is a helper method to define mock.On call
//   - formDataProvider domain.FormDataProvider
func (_e *FormHandlerBuilder_Expecter) SetFormDataProvider(formDataProvider interface{}) *FormHandlerBuilder_SetFormDataProvider_Call {
	return &FormHandlerBuilder_SetFormDataProvider_Call{Call: _e.mock.On("SetFormDataProvider", formDataProvider)}
}

func (_c *FormHandlerBuilder_SetFormDataProvider_Call) Run(run func(formDataProvider domain.FormDataProvider)) *FormHandlerBuilder_SetFormDataProvider_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(domain.FormDataProvider))
	})
	return _c
}

func (_c *FormHandlerBuilder_SetFormDataProvider_Call) Return(_a0 application.FormHandlerBuilder) *FormHandlerBuilder_SetFormDataProvider_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FormHandlerBuilder_SetFormDataProvider_Call) RunAndReturn(run func(domain.FormDataProvider) application.FormHandlerBuilder) *FormHandlerBuilder_SetFormDataProvider_Call {
	_c.Call.Return(run)
	return _c
}

// SetFormDataValidator provides a mock function with given fields: formDataValidator
func (_m *FormHandlerBuilder) SetFormDataValidator(formDataValidator domain.FormDataValidator) application.FormHandlerBuilder {
	ret := _m.Called(formDataValidator)

	if len(ret) == 0 {
		panic("no return value specified for SetFormDataValidator")
	}

	var r0 application.FormHandlerBuilder
	if rf, ok := ret.Get(0).(func(domain.FormDataValidator) application.FormHandlerBuilder); ok {
		r0 = rf(formDataValidator)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(application.FormHandlerBuilder)
		}
	}

	return r0
}

// FormHandlerBuilder_SetFormDataValidator_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetFormDataValidator'
type FormHandlerBuilder_SetFormDataValidator_Call struct {
	*mock.Call
}

// SetFormDataValidator is a helper method to define mock.On call
//   - formDataValidator domain.FormDataValidator
func (_e *FormHandlerBuilder_Expecter) SetFormDataValidator(formDataValidator interface{}) *FormHandlerBuilder_SetFormDataValidator_Call {
	return &FormHandlerBuilder_SetFormDataValidator_Call{Call: _e.mock.On("SetFormDataValidator", formDataValidator)}
}

func (_c *FormHandlerBuilder_SetFormDataValidator_Call) Run(run func(formDataValidator domain.FormDataValidator)) *FormHandlerBuilder_SetFormDataValidator_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(domain.FormDataValidator))
	})
	return _c
}

func (_c *FormHandlerBuilder_SetFormDataValidator_Call) Return(_a0 application.FormHandlerBuilder) *FormHandlerBuilder_SetFormDataValidator_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FormHandlerBuilder_SetFormDataValidator_Call) RunAndReturn(run func(domain.FormDataValidator) application.FormHandlerBuilder) *FormHandlerBuilder_SetFormDataValidator_Call {
	_c.Call.Return(run)
	return _c
}

// SetFormService provides a mock function with given fields: formService
func (_m *FormHandlerBuilder) SetFormService(formService domain.FormService) error {
	ret := _m.Called(formService)

	if len(ret) == 0 {
		panic("no return value specified for SetFormService")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.FormService) error); ok {
		r0 = rf(formService)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FormHandlerBuilder_SetFormService_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetFormService'
type FormHandlerBuilder_SetFormService_Call struct {
	*mock.Call
}

// SetFormService is a helper method to define mock.On call
//   - formService domain.FormService
func (_e *FormHandlerBuilder_Expecter) SetFormService(formService interface{}) *FormHandlerBuilder_SetFormService_Call {
	return &FormHandlerBuilder_SetFormService_Call{Call: _e.mock.On("SetFormService", formService)}
}

func (_c *FormHandlerBuilder_SetFormService_Call) Run(run func(formService domain.FormService)) *FormHandlerBuilder_SetFormService_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(domain.FormService))
	})
	return _c
}

func (_c *FormHandlerBuilder_SetFormService_Call) Return(_a0 error) *FormHandlerBuilder_SetFormService_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FormHandlerBuilder_SetFormService_Call) RunAndReturn(run func(domain.FormService) error) *FormHandlerBuilder_SetFormService_Call {
	_c.Call.Return(run)
	return _c
}

// SetNamedFormDataDecoder provides a mock function with given fields: name
func (_m *FormHandlerBuilder) SetNamedFormDataDecoder(name string) error {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for SetNamedFormDataDecoder")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FormHandlerBuilder_SetNamedFormDataDecoder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetNamedFormDataDecoder'
type FormHandlerBuilder_SetNamedFormDataDecoder_Call struct {
	*mock.Call
}

// SetNamedFormDataDecoder is a helper method to define mock.On call
//   - name string
func (_e *FormHandlerBuilder_Expecter) SetNamedFormDataDecoder(name interface{}) *FormHandlerBuilder_SetNamedFormDataDecoder_Call {
	return &FormHandlerBuilder_SetNamedFormDataDecoder_Call{Call: _e.mock.On("SetNamedFormDataDecoder", name)}
}

func (_c *FormHandlerBuilder_SetNamedFormDataDecoder_Call) Run(run func(name string)) *FormHandlerBuilder_SetNamedFormDataDecoder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *FormHandlerBuilder_SetNamedFormDataDecoder_Call) Return(_a0 error) *FormHandlerBuilder_SetNamedFormDataDecoder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FormHandlerBuilder_SetNamedFormDataDecoder_Call) RunAndReturn(run func(string) error) *FormHandlerBuilder_SetNamedFormDataDecoder_Call {
	_c.Call.Return(run)
	return _c
}

// SetNamedFormDataProvider provides a mock function with given fields: name
func (_m *FormHandlerBuilder) SetNamedFormDataProvider(name string) error {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for SetNamedFormDataProvider")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FormHandlerBuilder_SetNamedFormDataProvider_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetNamedFormDataProvider'
type FormHandlerBuilder_SetNamedFormDataProvider_Call struct {
	*mock.Call
}

// SetNamedFormDataProvider is a helper method to define mock.On call
//   - name string
func (_e *FormHandlerBuilder_Expecter) SetNamedFormDataProvider(name interface{}) *FormHandlerBuilder_SetNamedFormDataProvider_Call {
	return &FormHandlerBuilder_SetNamedFormDataProvider_Call{Call: _e.mock.On("SetNamedFormDataProvider", name)}
}

func (_c *FormHandlerBuilder_SetNamedFormDataProvider_Call) Run(run func(name string)) *FormHandlerBuilder_SetNamedFormDataProvider_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *FormHandlerBuilder_SetNamedFormDataProvider_Call) Return(_a0 error) *FormHandlerBuilder_SetNamedFormDataProvider_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FormHandlerBuilder_SetNamedFormDataProvider_Call) RunAndReturn(run func(string) error) *FormHandlerBuilder_SetNamedFormDataProvider_Call {
	_c.Call.Return(run)
	return _c
}

// SetNamedFormDataValidator provides a mock function with given fields: name
func (_m *FormHandlerBuilder) SetNamedFormDataValidator(name string) error {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for SetNamedFormDataValidator")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FormHandlerBuilder_SetNamedFormDataValidator_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetNamedFormDataValidator'
type FormHandlerBuilder_SetNamedFormDataValidator_Call struct {
	*mock.Call
}

// SetNamedFormDataValidator is a helper method to define mock.On call
//   - name string
func (_e *FormHandlerBuilder_Expecter) SetNamedFormDataValidator(name interface{}) *FormHandlerBuilder_SetNamedFormDataValidator_Call {
	return &FormHandlerBuilder_SetNamedFormDataValidator_Call{Call: _e.mock.On("SetNamedFormDataValidator", name)}
}

func (_c *FormHandlerBuilder_SetNamedFormDataValidator_Call) Run(run func(name string)) *FormHandlerBuilder_SetNamedFormDataValidator_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *FormHandlerBuilder_SetNamedFormDataValidator_Call) Return(_a0 error) *FormHandlerBuilder_SetNamedFormDataValidator_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FormHandlerBuilder_SetNamedFormDataValidator_Call) RunAndReturn(run func(string) error) *FormHandlerBuilder_SetNamedFormDataValidator_Call {
	_c.Call.Return(run)
	return _c
}

// SetNamedFormService provides a mock function with given fields: name
func (_m *FormHandlerBuilder) SetNamedFormService(name string) error {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for SetNamedFormService")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FormHandlerBuilder_SetNamedFormService_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetNamedFormService'
type FormHandlerBuilder_SetNamedFormService_Call struct {
	*mock.Call
}

// SetNamedFormService is a helper method to define mock.On call
//   - name string
func (_e *FormHandlerBuilder_Expecter) SetNamedFormService(name interface{}) *FormHandlerBuilder_SetNamedFormService_Call {
	return &FormHandlerBuilder_SetNamedFormService_Call{Call: _e.mock.On("SetNamedFormService", name)}
}

func (_c *FormHandlerBuilder_SetNamedFormService_Call) Run(run func(name string)) *FormHandlerBuilder_SetNamedFormService_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *FormHandlerBuilder_SetNamedFormService_Call) Return(_a0 error) *FormHandlerBuilder_SetNamedFormService_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FormHandlerBuilder_SetNamedFormService_Call) RunAndReturn(run func(string) error) *FormHandlerBuilder_SetNamedFormService_Call {
	_c.Call.Return(run)
	return _c
}

// NewFormHandlerBuilder creates a new instance of FormHandlerBuilder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFormHandlerBuilder(t interface {
	mock.TestingT
	Cleanup(func())
}) *FormHandlerBuilder {
	mock := &FormHandlerBuilder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
