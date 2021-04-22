package formdata

import (
	"context"
	"net/url"

	"github.com/go-playground/form"

	"flamingo.me/form/domain"
)

type (
	// DefaultFormDataEncoderImpl represents implementation of default domain.FormDataEncoder.
	DefaultFormDataEncoderImpl struct{}
)

var _ domain.DefaultFormDataEncoder = &DefaultFormDataEncoderImpl{}

// Encode performs default form data encoding, depending if passed form data is instance of map[string]string or any other interface.
func (p *DefaultFormDataEncoderImpl) Encode(_ context.Context, formData interface{}) (url.Values, error) {
	if data, ok := formData.(map[string]string); ok {
		urlValues := make(url.Values)
		for k, v := range data {
			urlValues[k] = []string{v}
		}
		return urlValues, nil
	}

	return p.encodeUnknownInterface(formData)
}

func (p *DefaultFormDataEncoderImpl) encodeUnknownInterface(formData interface{}) (url.Values, error) {
	encoder := form.NewEncoder()
	return encoder.Encode(formData)
}
