// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "flamingo.me/form/domain"
	mock "github.com/stretchr/testify/mock"

	web "flamingo.me/flamingo/v3/framework/web"
)

// FormHandler is an autogenerated mock type for the FormHandler type
type FormHandler struct {
	mock.Mock
}

type FormHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *FormHandler) EXPECT() *FormHandler_Expecter {
	return &FormHandler_Expecter{mock: &_m.Mock}
}

// HandleForm provides a mock function with given fields: ctx, req
func (_m *FormHandler) HandleForm(ctx context.Context, req *web.Request) (*domain.Form, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for HandleForm")
	}

	var r0 *domain.Form
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *web.Request) (*domain.Form, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *web.Request) *domain.Form); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Form)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *web.Request) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FormHandler_HandleForm_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HandleForm'
type FormHandler_HandleForm_Call struct {
	*mock.Call
}

// HandleForm is a helper method to define mock.On call
//   - ctx context.Context
//   - req *web.Request
func (_e *FormHandler_Expecter) HandleForm(ctx interface{}, req interface{}) *FormHandler_HandleForm_Call {
	return &FormHandler_HandleForm_Call{Call: _e.mock.On("HandleForm", ctx, req)}
}

func (_c *FormHandler_HandleForm_Call) Run(run func(ctx context.Context, req *web.Request)) *FormHandler_HandleForm_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*web.Request))
	})
	return _c
}

func (_c *FormHandler_HandleForm_Call) Return(_a0 *domain.Form, _a1 error) *FormHandler_HandleForm_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FormHandler_HandleForm_Call) RunAndReturn(run func(context.Context, *web.Request) (*domain.Form, error)) *FormHandler_HandleForm_Call {
	_c.Call.Return(run)
	return _c
}

// HandleSubmittedForm provides a mock function with given fields: ctx, req
func (_m *FormHandler) HandleSubmittedForm(ctx context.Context, req *web.Request) (*domain.Form, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for HandleSubmittedForm")
	}

	var r0 *domain.Form
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *web.Request) (*domain.Form, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *web.Request) *domain.Form); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Form)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *web.Request) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FormHandler_HandleSubmittedForm_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HandleSubmittedForm'
type FormHandler_HandleSubmittedForm_Call struct {
	*mock.Call
}

// HandleSubmittedForm is a helper method to define mock.On call
//   - ctx context.Context
//   - req *web.Request
func (_e *FormHandler_Expecter) HandleSubmittedForm(ctx interface{}, req interface{}) *FormHandler_HandleSubmittedForm_Call {
	return &FormHandler_HandleSubmittedForm_Call{Call: _e.mock.On("HandleSubmittedForm", ctx, req)}
}

func (_c *FormHandler_HandleSubmittedForm_Call) Run(run func(ctx context.Context, req *web.Request)) *FormHandler_HandleSubmittedForm_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*web.Request))
	})
	return _c
}

func (_c *FormHandler_HandleSubmittedForm_Call) Return(_a0 *domain.Form, _a1 error) *FormHandler_HandleSubmittedForm_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FormHandler_HandleSubmittedForm_Call) RunAndReturn(run func(context.Context, *web.Request) (*domain.Form, error)) *FormHandler_HandleSubmittedForm_Call {
	_c.Call.Return(run)
	return _c
}

// HandleSubmittedGETForm provides a mock function with given fields: ctx, req
func (_m *FormHandler) HandleSubmittedGETForm(ctx context.Context, req *web.Request) (*domain.Form, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for HandleSubmittedGETForm")
	}

	var r0 *domain.Form
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *web.Request) (*domain.Form, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *web.Request) *domain.Form); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Form)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *web.Request) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FormHandler_HandleSubmittedGETForm_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HandleSubmittedGETForm'
type FormHandler_HandleSubmittedGETForm_Call struct {
	*mock.Call
}

// HandleSubmittedGETForm is a helper method to define mock.On call
//   - ctx context.Context
//   - req *web.Request
func (_e *FormHandler_Expecter) HandleSubmittedGETForm(ctx interface{}, req interface{}) *FormHandler_HandleSubmittedGETForm_Call {
	return &FormHandler_HandleSubmittedGETForm_Call{Call: _e.mock.On("HandleSubmittedGETForm", ctx, req)}
}

func (_c *FormHandler_HandleSubmittedGETForm_Call) Run(run func(ctx context.Context, req *web.Request)) *FormHandler_HandleSubmittedGETForm_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*web.Request))
	})
	return _c
}

func (_c *FormHandler_HandleSubmittedGETForm_Call) Return(_a0 *domain.Form, _a1 error) *FormHandler_HandleSubmittedGETForm_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FormHandler_HandleSubmittedGETForm_Call) RunAndReturn(run func(context.Context, *web.Request) (*domain.Form, error)) *FormHandler_HandleSubmittedGETForm_Call {
	_c.Call.Return(run)
	return _c
}

// HandleUnsubmittedForm provides a mock function with given fields: ctx, req
func (_m *FormHandler) HandleUnsubmittedForm(ctx context.Context, req *web.Request) (*domain.Form, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for HandleUnsubmittedForm")
	}

	var r0 *domain.Form
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *web.Request) (*domain.Form, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *web.Request) *domain.Form); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Form)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *web.Request) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FormHandler_HandleUnsubmittedForm_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HandleUnsubmittedForm'
type FormHandler_HandleUnsubmittedForm_Call struct {
	*mock.Call
}

// HandleUnsubmittedForm is a helper method to define mock.On call
//   - ctx context.Context
//   - req *web.Request
func (_e *FormHandler_Expecter) HandleUnsubmittedForm(ctx interface{}, req interface{}) *FormHandler_HandleUnsubmittedForm_Call {
	return &FormHandler_HandleUnsubmittedForm_Call{Call: _e.mock.On("HandleUnsubmittedForm", ctx, req)}
}

func (_c *FormHandler_HandleUnsubmittedForm_Call) Run(run func(ctx context.Context, req *web.Request)) *FormHandler_HandleUnsubmittedForm_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*web.Request))
	})
	return _c
}

func (_c *FormHandler_HandleUnsubmittedForm_Call) Return(_a0 *domain.Form, _a1 error) *FormHandler_HandleUnsubmittedForm_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FormHandler_HandleUnsubmittedForm_Call) RunAndReturn(run func(context.Context, *web.Request) (*domain.Form, error)) *FormHandler_HandleUnsubmittedForm_Call {
	_c.Call.Return(run)
	return _c
}

// NewFormHandler creates a new instance of FormHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFormHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *FormHandler {
	mock := &FormHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
