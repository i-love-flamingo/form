package domain

import "fmt"

// Form as struct for storing form processing results
type Form struct {
	// Data  the form Data Struct (Forms DTO)
	Data interface{}
	// FormExtensionsData the additional form Data Structs (Forms DTO) fetched from form extensions
	FormExtensionsData map[string]interface{}
	// ValidationInfo for the form
	ValidationInfo ValidationInfo
	// submitted  flag if form was submitted and this is the result page
	submitted bool
	// validationRules contains map with validation rules for all validatable fields
	validationRules map[string][]ValidationRule
}

// FormError is used as wrapper for storing form error messages
type FormError struct {
	details string
	parent  error
}

// NewForm returns new instance of Form struct
func NewForm(submitted bool, validationRules map[string][]ValidationRule) Form {
	return Form{
		submitted:       submitted,
		validationRules: validationRules,
	}
}

// IsValidAndSubmitted defines if form data is valid and form is submitted
func (f Form) IsValidAndSubmitted() bool {
	return f.IsValid() && f.IsSubmitted()
}

// IsValid defines if form data is valid
func (f Form) IsValid() bool {
	return f.ValidationInfo.IsValid()
}

// IsSubmitted defines if form is submitted
func (f Form) IsSubmitted() bool {
	return f.submitted
}

// HasErrorForField method which defines if there is any field validations error for specific field
func (f Form) HasErrorForField(name string) bool {
	return f.ValidationInfo.HasErrorsForField(name)
}

// HasAnyFieldErrors method which defines if there is any field validations error for any field
func (f Form) HasAnyFieldErrors() bool {
	return f.ValidationInfo.HasAnyFieldErrors()
}

// HasGeneralErrors method which defines if there is any general validations error
func (f Form) HasGeneralErrors() bool {
	return f.ValidationInfo.HasGeneralErrors()
}

// GetGeneralErrors method which returns list of all general validation errors
func (f Form) GetGeneralErrors() []Error {
	return f.ValidationInfo.generalErrors
}

// GetErrorsForField method which returns list of all general validation errors for specific field
func (f Form) GetErrorsForField(name string) []Error {
	return f.ValidationInfo.GetErrorsForField(name)
}

// GetValidationRulesForField adds option to extract validation rules for desired field in templates
func (f Form) GetValidationRulesForField(name string) []ValidationRule {
	return f.validationRules[name]
}

// GetValidationRules returns all available validation rules
func (f Form) GetValidationRules() map[string][]ValidationRule {
	return f.validationRules
}

// NewFormError returns new instance of error interface by defining string content of error
func NewFormError(details string) FormError {
	return FormError{
		details: details,
	}
}

// NewFormErrorf returns new instance of error interface by defining formatted string content of error with arguments
func NewFormErrorf(details string, args ...interface{}) FormError {
	return NewFormError(fmt.Sprintf(details, args...))
}

// NewFormErrorWithParent returns new instance of error interface by defining parent error
func NewFormErrorWithParent(err error) FormError {
	return FormError{
		details: err.Error(),
		parent:  err,
	}
}

// Error represents implementation for required method so FormError can fulfil error interface
func (e FormError) Error() string {
	return fmt.Sprintf("FormError: %s", e.details)
}

// Parent returns parent error wrapped by FormError
func (e FormError) Parent() error {
	return e.parent
}