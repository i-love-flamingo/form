// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	url "net/url"
)

// FormDataEncoder is an autogenerated mock type for the FormDataEncoder type
type FormDataEncoder struct {
	mock.Mock
}

// Encode provides a mock function with given fields: ctx, formData
func (_m *FormDataEncoder) Encode(ctx context.Context, formData interface{}) (url.Values, error) {
	ret := _m.Called(ctx, formData)

	var r0 url.Values
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) url.Values); ok {
		r0 = rf(ctx, formData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(url.Values)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, formData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
