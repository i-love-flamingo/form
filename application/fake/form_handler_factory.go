package fake

import (
	"flamingo.me/form/application"
	"flamingo.me/form/domain"
)

type (
	// FormHandlerFactoryImpl defines faked implementation of FormHandlerFactory interface used for unit testing
	FormHandlerFactoryImpl struct {
		formHandler domain.FormHandler
	}
)

// New returns faked implementation of FormHandlerFactory interface which should deliver mocked domain.FormHandler instance
func New(formHandler domain.FormHandler) application.FormHandlerFactory {
	return &FormHandlerFactoryImpl{
		formHandler: formHandler,
	}
}

// CreateSimpleFormHandler returns mocked instance of domain.FormHandler interface
func (f *FormHandlerFactoryImpl) CreateSimpleFormHandler() domain.FormHandler {
	return f.formHandler
}

// CreateFormHandlerWithFormService returns mocked instance of domain.FormHandler interface
func (f *FormHandlerFactoryImpl) CreateFormHandlerWithFormService(domain.FormService, ...string) domain.FormHandler {
	return f.formHandler
}

// CreateFormHandlerWithFormServices returns mocked instance of domain.FormHandler interface
func (f *FormHandlerFactoryImpl) CreateFormHandlerWithFormServices(domain.FormDataProvider, domain.FormDataDecoder, domain.FormDataValidator, ...string) domain.FormHandler {
	return f.formHandler
}

// GetFormHandlerBuilder returns faked instance of FormHandlerBuilder interface
func (f *FormHandlerFactoryImpl) GetFormHandlerBuilder() application.FormHandlerBuilder {
	return &formHandlerBuilderImpl{
		formHandler: f.formHandler,
	}
}
