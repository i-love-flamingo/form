// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	url "net/url"

	web "flamingo.me/flamingo/v3/framework/web"
)

// DefaultFormDataDecoder is an autogenerated mock type for the DefaultFormDataDecoder type
type DefaultFormDataDecoder struct {
	mock.Mock
}

// Decode provides a mock function with given fields: ctx, req, values, formData
func (_m *DefaultFormDataDecoder) Decode(ctx context.Context, req *web.Request, values url.Values, formData interface{}) (interface{}, error) {
	ret := _m.Called(ctx, req, values, formData)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(context.Context, *web.Request, url.Values, interface{}) interface{}); ok {
		r0 = rf(ctx, req, values, formData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *web.Request, url.Values, interface{}) error); ok {
		r1 = rf(ctx, req, values, formData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
