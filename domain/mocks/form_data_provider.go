// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	web "flamingo.me/flamingo/v3/framework/web"
)

// FormDataProvider is an autogenerated mock type for the FormDataProvider type
type FormDataProvider struct {
	mock.Mock
}

// GetFormData provides a mock function with given fields: ctx, req
func (_m *FormDataProvider) GetFormData(ctx context.Context, req *web.Request) (interface{}, error) {
	ret := _m.Called(ctx, req)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(context.Context, *web.Request) interface{}); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *web.Request) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
