package validators

import (
	"context"
	"strings"
	"time"

	"flamingo.me/form/domain"

	validator "github.com/go-playground/validator/v10"
)

type (
	// DateFormatValidator defines date format validator which validates date format depending on application's configuration
	//
	// Data struct {
	//	 Date string `validate:"dateformat"`
	// }
	//
	DateFormatValidator struct {
		dateFormat string
	}
)

var _ domain.FieldValidator = &DateFormatValidator{}

// Inject is method used to set all dependencies as local variables
func (v *DateFormatValidator) Inject(cfg *struct {
	DateFormat string `inject:"config:form.validator.dateFormat"`
}) {
	v.dateFormat = cfg.DateFormat
}

// ValidatorName defines tag name of date format validator
func (v *DateFormatValidator) ValidatorName() string {
	return "dateformat"
}

// ValidateField validates string for right date format. Valid if string is empty or in right date format.
func (v *DateFormatValidator) ValidateField(_ context.Context, fl validator.FieldLevel) bool {
	converted, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	if len(strings.TrimSpace(converted)) == 0 {
		return true
	}

	_, err := time.Parse(v.dateFormat, converted)
	return err == nil
}
