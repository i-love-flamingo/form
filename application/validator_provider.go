package application

import (
	"context"
	"strings"

	"github.com/go-playground/validator/v10"

	"flamingo.me/flamingo/v3/framework/web"
	"flamingo.me/form/domain"
)

type (
	// ValidatorProviderImpl as struct which implements interface ValidatorProvider
	ValidatorProviderImpl struct {
		validate *validator.Validate
	}
)

var _ domain.ValidatorProvider = &ValidatorProviderImpl{}

// Inject initialize instance of validator.Validate struct
func (p *ValidatorProviderImpl) Inject(fieldValidators []domain.FieldValidator, structValidators []domain.StructValidator) {
	validate := validator.New()
	p.attachFieldValidators(validate, fieldValidators)
	p.attachStructValidators(validate, structValidators)
	p.validate = validate
}

// Validate method which validates any struct and returns domain.ValidationInfo as a result of validation
func (p *ValidatorProviderImpl) Validate(ctx context.Context, req *web.Request, value interface{}) domain.ValidationInfo {
	reqCtx := web.ContextWithRequest(ctx, req)
	validate := p.GetValidator()
	err := validate.StructCtx(reqCtx, value)

	return p.ErrorsToValidationInfo(err)
}

// GetValidator method which returns instance of validator.Validate struct with all injected field and struct validations
func (p *ValidatorProviderImpl) GetValidator() *validator.Validate {
	return p.validate
}

// ErrorsToValidationInfo method which transforms errors into domain.ValidationInfo
func (p *ValidatorProviderImpl) ErrorsToValidationInfo(err error) domain.ValidationInfo {
	validationInfo := domain.ValidationInfo{}

	if err == nil {
		return validationInfo
	}

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, err := range validationErrors {
			fieldName := p.getRelativeFieldNameFromValidationError(err)
			validationInfo.AddFieldError(fieldName, "formError."+fieldName+"."+err.Tag(), err.Field()+" "+err.Tag())
		}
	} else {
		validationInfo.AddGeneralError("formError.invalidValidation", err.Error())
	}

	return validationInfo
}

// attachFieldValidators method which attach all injected instances of FieldValidator interface into validator.Validate instance
func (p *ValidatorProviderImpl) attachFieldValidators(validate *validator.Validate, fieldValidators []domain.FieldValidator) {
	for _, fieldValidator := range fieldValidators {
		validate.RegisterValidationCtx(fieldValidator.ValidatorName(), fieldValidator.ValidateField)
	}
}

// attachFieldValidators method which attach all injected instances of StructValidator interface into validator.Validate instance
func (p *ValidatorProviderImpl) attachStructValidators(validate *validator.Validate, structValidators []domain.StructValidator) {
	for _, structValidator := range structValidators {
		validate.RegisterStructValidationCtx(structValidator.ValidateStruct, structValidator.StructType())
	}
}

// getRelativeFieldNameFromValidationError method which extracts relative field name depending on it's full namespace
func (p *ValidatorProviderImpl) getRelativeFieldNameFromValidationError(err validator.FieldError) string {
	namespace := err.Namespace()

	//first part of namespace is not required to have the relative path:
	fieldName := namespace[(strings.Index(namespace, ".") + 1):]

	// initialize array of namespace parts
	parts := strings.Split(fieldName, ".")
	result := make([]string, len(parts))

	for i, part := range parts {
		result[i] = strings.ToLower(part[0:1]) + part[1:]
	}

	return strings.Join(result, ".")
}
