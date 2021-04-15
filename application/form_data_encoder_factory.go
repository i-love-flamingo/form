package application

import "flamingo.me/form/domain"

type (
	// FormDataEncoderFactory as interface for simple creation of form encoder instance
	FormDataEncoderFactory interface {
		CreateWithFormService(formService domain.FormService) domain.FormDataEncoder
		CreateByNamedEncoder(name string) domain.FormDataEncoder
	}

	// FormDataEncoderFactoryImpl as actual implementation of FormDataEncoderFactory interface
	FormDataEncoderFactoryImpl struct {
		namedFormServices      map[string]domain.FormService
		namedFormDataEncoders  map[string]domain.FormDataEncoder
		defaultFormDataEncoder domain.DefaultFormDataEncoder
	}
)

// Inject is method used to set all dependencies as local variables
func (f *FormDataEncoderFactoryImpl) Inject(
	namedFormDataEncoders map[string]domain.FormDataEncoder,
	namedFormServices map[string]domain.FormService,
	defaultFormDataEncoder domain.DefaultFormDataEncoder,

) {
	f.namedFormServices = namedFormServices
	f.namedFormDataEncoders = namedFormDataEncoders
	f.defaultFormDataEncoder = defaultFormDataEncoder
}

// CreateWithFormService - factory method
func (f *FormDataEncoderFactoryImpl) CreateWithFormService(formService domain.FormService) domain.FormDataEncoder {
	if encoder, ok := formService.(domain.FormDataEncoder); ok {
		return encoder
	}
	return f.defaultFormDataEncoder
}

// CreateByNamedEncoder - factory method
func (f *FormDataEncoderFactoryImpl) CreateByNamedEncoder(name string) domain.FormDataEncoder {
	if encoder, ok := f.namedFormDataEncoders[name]; ok {
		return encoder
	}
	if service, ok := f.namedFormServices[name]; ok {
		return f.CreateWithFormService(service)
	}
	return f.defaultFormDataEncoder
}
