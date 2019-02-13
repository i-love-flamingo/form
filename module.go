package form

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/config"
	"flamingo.me/form/application"
	"flamingo.me/form/domain"
	"flamingo.me/form/domain/formdata"
	"flamingo.me/form/domain/validators"
)

type (
	// Module is struct for defining form2 module dependencies
	Module struct {
		CustomRegex config.Map `inject:"config:form.validator.customRegex"`
	}
)

// Configure is main method for handling module dependencies via dingo injector
func (m *Module) Configure(injector *dingo.Injector) {
	for name, value := range m.CustomRegex {
		regex, ok := value.(string)
		if !ok {
			panic("wrong value passed as validation regex")
		}
		regexValidator := validators.NewRegexValidator(name, regex)
		injector.BindMulti(new(domain.FieldValidator)).ToInstance(regexValidator)
	}
	injector.BindMulti(new(domain.FieldValidator)).To(validators.DateFormatValidator{})
	injector.BindMulti(new(domain.FieldValidator)).To(validators.MinimumAgeValidator{})
	injector.BindMulti(new(domain.FieldValidator)).To(validators.MaximumAgeValidator{})

	injector.Bind(new(domain.ValidatorProvider)).To(application.ValidatorProviderImpl{}).AsEagerSingleton()

	injector.Bind(new(domain.DefaultFormDataProvider)).To(formdata.DefaultFormDataProviderImpl{})
	injector.Bind(new(domain.DefaultFormDataDecoder)).To(formdata.DefaultFormDataDecoderImpl{})
	injector.Bind(new(domain.DefaultFormDataValidator)).To(formdata.DefaultFormDataValidatorImpl{})
	injector.Bind(new(application.FormHandlerFactory)).To(application.FormHandlerFactoryImpl{}).AsEagerSingleton()
}

// DefaultConfig is method which is responsible for setting up default module configuration
func (m *Module) DefaultConfig() config.Map {
	return config.Map{
		"form.validator": config.Map{
			"dateFormat":  "2006-01-02",
			"customRegex": config.Map{},
		},
	}
}
