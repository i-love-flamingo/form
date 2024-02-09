package fake

import (
	"flamingo.me/form/domain/mocks"
	"reflect"
)

// NewFieldLevel is helper method to provide mocked instance of validator.FieldLevel interface
func NewFieldLevel(value interface{}, param string) *mocks.FieldLevel {
	reflected := reflect.ValueOf(value)

	fieldLevel := &mocks.FieldLevel{}

	fieldLevel.On("Field").Return(reflected).Maybe()
	fieldLevel.On("Param").Return(param).Maybe()

	return fieldLevel
}

// NewStructLevel is helper method to provide mocked instance of validator.StructLevel interface
func NewStructLevel(value interface{}) *mocks.StructLevel {
	reflected := reflect.ValueOf(value)

	structLevel := &mocks.StructLevel{}

	structLevel.On("Current").Return(reflected).Maybe()

	return structLevel
}
