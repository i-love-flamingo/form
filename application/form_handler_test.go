package application

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"testing"

	"flamingo.me/flamingo/v3/framework/flamingo"
	"flamingo.me/flamingo/v3/framework/web"
	"flamingo.me/form/domain"
	"flamingo.me/form/domain/mocks"
	"github.com/stretchr/testify/suite"
)

type (
	FormHandlerImplTestSuite struct {
		suite.Suite

		handler *formHandlerImpl

		defaultProvider   *mocks.DefaultFormDataProvider
		defaultDecoder    *mocks.DefaultFormDataDecoder
		defaultValidator  *mocks.DefaultFormDataValidator
		provider          *mocks.FormDataProvider
		decoder           *mocks.FormDataDecoder
		validator         *mocks.FormDataValidator
		firstExtension    *mocks.CompleteFormService
		secondExtension   *mocks.FormDataProvider
		thirdExtension    *mocks.FormDataDecoder
		fourthExtension   *mocks.FormDataValidator
		validatorProvider *mocks.ValidatorProvider
		logger            *flamingo.NullLogger

		context context.Context
		request *web.Request
	}
)

func TestFormHandlerImplTestSuite(t *testing.T) {
	suite.Run(t, &FormHandlerImplTestSuite{})
}

func (t *FormHandlerImplTestSuite) SetupSuite() {
	t.context = context.Background()
}

func (t *FormHandlerImplTestSuite) SetupTest() {
	t.defaultProvider = &mocks.DefaultFormDataProvider{}
	t.defaultDecoder = &mocks.DefaultFormDataDecoder{}
	t.defaultValidator = &mocks.DefaultFormDataValidator{}
	t.provider = &mocks.FormDataProvider{}
	t.decoder = &mocks.FormDataDecoder{}
	t.validator = &mocks.FormDataValidator{}
	t.firstExtension = &mocks.CompleteFormService{}
	t.secondExtension = &mocks.FormDataProvider{}
	t.thirdExtension = &mocks.FormDataDecoder{}
	t.fourthExtension = &mocks.FormDataValidator{}
	t.validatorProvider = &mocks.ValidatorProvider{}
	t.logger = &flamingo.NullLogger{}

	t.handler = &formHandlerImpl{
		defaultFormDataProvider:  t.defaultProvider,
		defaultFormDataDecoder:   t.defaultDecoder,
		defaultFormDataValidator: t.defaultValidator,
		formDataProvider:         t.provider,
		formDataDecoder:          t.decoder,
		formDataValidator:        t.validator,
		formExtensions: map[string]domain.FormExtension{
			"first":  t.firstExtension,
			"second": t.secondExtension,
			"third":  t.thirdExtension,
			"fourth": t.fourthExtension,
		},
		validatorProvider: t.validatorProvider,
		logger:            t.logger,
	}

	t.request = web.CreateRequest(&http.Request{}, nil)
}

func (t *FormHandlerImplTestSuite) TearDownTest() {
	t.defaultProvider.AssertExpectations(t.T())
	t.defaultDecoder.AssertExpectations(t.T())
	t.defaultValidator.AssertExpectations(t.T())
	t.provider.AssertExpectations(t.T())
	t.decoder.AssertExpectations(t.T())
	t.validator.AssertExpectations(t.T())
	t.firstExtension.AssertExpectations(t.T())
	t.secondExtension.AssertExpectations(t.T())
	t.thirdExtension.AssertExpectations(t.T())
	t.fourthExtension.AssertExpectations(t.T())
	t.validatorProvider.AssertExpectations(t.T())
}

func (t *FormHandlerImplTestSuite) TestHandleUnsubmittedForm_Error() {
	t.provider.On("GetFormData", t.context, t.request).Return(nil, errors.New("error")).Once()

	t.firstExtension.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Once()
	t.secondExtension.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Once()
	t.defaultProvider.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Twice()

	result, err := t.handler.HandleUnsubmittedForm(t.context, t.request)
	t.Equal(domain.NewFormErrorWithParent(errors.New("error")), err)
	t.Nil(result)
}

func (t *FormHandlerImplTestSuite) TestHandleUnsubmittedForm_Success() {
	t.provider.On("GetFormData", t.context, t.request).Return(struct {
		FirstField string `form:"firstField" validate:"required"`
	}{}, nil).Once()

	t.firstExtension.On("GetFormData", t.context, t.request).Return(struct {
		SecondField string `form:"secondField" validate:"required"`
	}{}, nil).Twice()
	t.secondExtension.On("GetFormData", t.context, t.request).Return(struct {
		ThirdField string `form:"thirdField" validate:"required"`
	}{}, nil).Twice()
	t.defaultProvider.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Times(4)

	result, err := t.handler.HandleUnsubmittedForm(t.context, t.request)
	t.NoError(err)

	form := domain.NewForm(false, map[string][]domain.ValidationRule{
		"firstField": {
			{
				Name: "required",
			},
		},
		"secondField": {
			{
				Name: "required",
			},
		},
		"thirdField": {
			{
				Name: "required",
			},
		},
	})
	form.Data = struct {
		FirstField string `form:"firstField" validate:"required"`
	}{}
	form.FormExtensionsData = map[string]interface{}{
		"first": struct {
			SecondField string `form:"secondField" validate:"required"`
		}{},
		"second": struct {
			ThirdField string `form:"thirdField" validate:"required"`
		}{},
		"third":  map[string]int{},
		"fourth": map[string]int{},
	}

	t.Equal(&form, result)
}

func (t *FormHandlerImplTestSuite) TestExtractValidationRules_NotStruct() {
	t.Equal(map[string][]domain.ValidationRule{}, t.handler.extractValidationRules(nil))
	t.Equal(map[string][]domain.ValidationRule{}, t.handler.extractValidationRules("string"))
	t.Equal(map[string][]domain.ValidationRule{}, t.handler.extractValidationRules(1))
	t.Equal(map[string][]domain.ValidationRule{}, t.handler.extractValidationRules(map[string]interface{}{}))
}

func (t *FormHandlerImplTestSuite) TestExtractValidationRules_Struct() {
	t.Equal(map[string][]domain.ValidationRule{
		"zero": {
			{
				Name: "required",
			},
		},
		"first": {
			{
				Name: "required",
			},
			{
				Name:  "gte",
				Value: "10",
			},
		},
		"second": {
			{
				Name:  "gte",
				Value: "10",
			},
		},
		"Sixth": {
			{
				Name: "required",
			},
			{
				Name:  "gte",
				Value: "10",
			},
		},
		"Seventh": {
			{
				Name: "required",
			},
		},
		"subStruct.first": {
			{
				Name: "required",
			},
			{
				Name:  "gte",
				Value: "10",
			},
		},
		"subStruct.second": {
			{
				Name:  "gte",
				Value: "10",
			},
		},
		"subStruct.Sixth": {
			{
				Name: "required",
			},
			{
				Name:  "gte",
				Value: "10",
			},
		},
		"StructWithoutName.first": {
			{
				Name: "required",
			},
			{
				Name:  "gte",
				Value: "10",
			},
		},
		"StructWithoutName.second": {
			{
				Name:  "gte",
				Value: "10",
			},
		},
		"StructWithoutName.Sixth": {
			{
				Name: "required",
			},
			{
				Name:  "gte",
				Value: "10",
			},
		},
		"referenceStruct.first": {
			{
				Name: "required",
			},
			{
				Name:  "gte",
				Value: "10",
			},
		},
		"referenceStruct.second": {
			{
				Name:  "gte",
				Value: "10",
			},
		},
		"referenceStruct.Sixth": {
			{
				Name: "required",
			},
			{
				Name:  "gte",
				Value: "10",
			},
		},
		"ReferenceWithoutName.first": {
			{
				Name: "required",
			},
			{
				Name:  "gte",
				Value: "10",
			},
		},
		"ReferenceWithoutName.second": {
			{
				Name:  "gte",
				Value: "10",
			},
		},
		"ReferenceWithoutName.Sixth": {
			{
				Name: "required",
			},
			{
				Name:  "gte",
				Value: "10",
			},
		},
	}, t.handler.extractValidationRules(struct {
		Zero           interface{} `form:"zero" validate:"required"`
		First          string      `form:"first" validate:"required,gte=10"`
		Second         string      `form:"second" validate:"omitempty,gte=10"`
		Third          string      `form:"-" validate:"required,gte=10"`
		Fourth         string      `form:"fourth" validate:""`
		Fifth          string      `form:"fifth"`
		Sixth          string      `validate:"required,gte=10"`
		Seventh        *string     `validate:"required"`
		StructWithName struct {
			First  string `form:"first" validate:"required,gte=10"`
			Second string `form:"second" validate:"omitempty,gte=10"`
			Third  string `form:"-" validate:"required,gte=10"`
			Fourth string `form:"fourth" validate:""`
			Fifth  string `form:"fifth"`
			Sixth  string `validate:"required,gte=10"`
		} `form:"subStruct"`
		StructWithoutName struct {
			First  string `form:"first" validate:"required,gte=10"`
			Second string `form:"second" validate:"omitempty,gte=10"`
			Third  string `form:"-" validate:"required,gte=10"`
			Fourth string `form:"fourth" validate:""`
			Fifth  string `form:"fifth"`
			Sixth  string `validate:"required,gte=10"`
		}
		ReferenceWithName *struct {
			First  string `form:"first" validate:"required,gte=10"`
			Second string `form:"second" validate:"omitempty,gte=10"`
			Third  string `form:"-" validate:"required,gte=10"`
			Fourth string `form:"fourth" validate:""`
			Fifth  string `form:"fifth"`
			Sixth  string `validate:"required,gte=10"`
		} `form:"referenceStruct"`
		ReferenceWithoutName *struct {
			First  string `form:"first" validate:"required,gte=10"`
			Second string `form:"second" validate:"omitempty,gte=10"`
			Third  string `form:"-" validate:"required,gte=10"`
			Fourth string `form:"fourth" validate:""`
			Fifth  string `form:"fifth"`
			Sixth  string `validate:"required,gte=10"`
		}
	}{}))
}

func (t *FormHandlerImplTestSuite) TestCollectFormExtensionValidationRules() {
	t.firstExtension.On("GetFormData", t.context, t.request).Return(struct {
		FirstFirstField  string `form:"firstFirstField" validate:"required,min=10"`
		FirstSecondField string `form:"firstSecondField" validate:"required"`
	}{}, nil).Once()
	t.secondExtension.On("GetFormData", t.context, t.request).Return(struct {
		SecondFirstField  string `form:"secondFirstField" validate:"required,min=10"`
		SecondSecondField string `form:"secondSecondField" validate:"required"`
	}{}, nil).Once()
	t.defaultProvider.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Twice()

	result, err := t.handler.collectFormExtensionValidationRules(t.context, t.request)
	t.NoError(err)

	t.Equal(map[string][]domain.ValidationRule{
		"firstFirstField": {
			{
				Name: "required",
			},
			{
				Name:  "min",
				Value: "10",
			},
		},
		"firstSecondField": {
			{
				Name: "required",
			},
		},
		"secondFirstField": {
			{
				Name: "required",
			},
			{
				Name:  "min",
				Value: "10",
			},
		},
		"secondSecondField": {
			{
				Name: "required",
			},
		},
	}, result)
}

func (t *FormHandlerImplTestSuite) TestGetUrlValues_PostError() {
	t.request.Request().Method = http.MethodPost

	values, err := t.handler.getURLValues(t.request, http.MethodPost)
	t.Error(err)
	t.Nil(values)
}

func (t *FormHandlerImplTestSuite) TestGetUrlValues_PostSuccess() {
	t.request.Request().Method = http.MethodPost
	t.request.Request().PostForm = url.Values{
		"first":  []string{"first"},
		"second": []string{"second"},
	}

	values, err := t.handler.getURLValues(t.request, http.MethodPost)
	t.NoError(err)
	t.Equal(&url.Values{
		"first":  []string{"first"},
		"second": []string{"second"},
	}, values)
}

func (t *FormHandlerImplTestSuite) TestGetUrlValues_GetSuccess() {
	t.request.Request().Method = http.MethodGet
	t.request.Request().URL = &url.URL{
		RawQuery: url.Values{
			"first":  []string{"first"},
			"second": []string{"second"},
		}.Encode(),
	}

	values, err := t.handler.getURLValues(t.request, http.MethodGet)
	t.NoError(err)
	t.Equal(&url.Values{
		"first":  []string{"first"},
		"second": []string{"second"},
	}, values)
}

func (t *FormHandlerImplTestSuite) TestProcessExtension_ProviderError() {
	t.secondExtension.On("GetFormData", t.context, t.request).Return(nil, errors.New("error")).Once()

	err := t.handler.processExtension(t.context, t.request, url.Values{}, "second", t.secondExtension, &domain.Form{})
	t.Error(err)
}

func (t *FormHandlerImplTestSuite) TestProcessExtension_ProviderSuccess() {
	t.secondExtension.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Once()
	t.defaultDecoder.On("Decode", t.context, t.request, url.Values{}, map[string]int{}).Return(map[string]int{}, nil).Once()
	t.defaultValidator.On("Validate", t.context, t.request, t.validatorProvider, map[string]int{}).Return(&domain.ValidationInfo{}, nil).Once()

	form := domain.NewForm(true, nil)
	err := t.handler.processExtension(t.context, t.request, url.Values{}, "second", t.secondExtension, &form)
	t.NoError(err)
}

func (t *FormHandlerImplTestSuite) TestProcessExtension_DecoderError() {
	t.defaultProvider.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Once()
	t.thirdExtension.On("Decode", t.context, t.request, url.Values{}, map[string]int{}).Return(nil, errors.New("error")).Once()

	form := domain.NewForm(true, nil)
	err := t.handler.processExtension(t.context, t.request, url.Values{}, "third", t.thirdExtension, &form)
	t.Error(err)
}

func (t *FormHandlerImplTestSuite) TestProcessExtension_DecoderSuccess() {
	t.defaultProvider.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Once()
	t.thirdExtension.On("Decode", t.context, t.request, url.Values{}, map[string]int{}).Return(map[string]int{}, nil).Once()
	t.defaultValidator.On("Validate", t.context, t.request, t.validatorProvider, map[string]int{}).Return(&domain.ValidationInfo{}, nil).Once()

	form := domain.NewForm(true, nil)
	err := t.handler.processExtension(t.context, t.request, url.Values{}, "third", t.thirdExtension, &form)
	t.NoError(err)
}

func (t *FormHandlerImplTestSuite) TestProcessExtension_ValidatorError() {
	t.defaultProvider.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Once()
	t.defaultDecoder.On("Decode", t.context, t.request, url.Values{}, map[string]int{}).Return(map[string]int{}, nil).Once()
	t.fourthExtension.On("Validate", t.context, t.request, t.validatorProvider, map[string]int{}).Return(nil, errors.New("error")).Once()

	form := domain.NewForm(true, nil)
	err := t.handler.processExtension(t.context, t.request, url.Values{}, "fourth", t.fourthExtension, &form)
	t.Error(err)
}

func (t *FormHandlerImplTestSuite) TestProcessExtension_ValidatorSuccess() {
	t.defaultProvider.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Once()
	t.defaultDecoder.On("Decode", t.context, t.request, url.Values{}, map[string]int{}).Return(map[string]int{}, nil).Once()
	t.fourthExtension.On("Validate", t.context, t.request, t.validatorProvider, map[string]int{}).Return(&domain.ValidationInfo{}, nil).Once()

	form := domain.NewForm(true, nil)
	err := t.handler.processExtension(t.context, t.request, url.Values{}, "fourth", t.fourthExtension, &form)
	t.NoError(err)
}

func (t *FormHandlerImplTestSuite) TestHandleSubmittedForm_GetFormDataError() {
	t.provider.On("GetFormData", t.context, t.request).Return(nil, errors.New("error")).Once()

	t.firstExtension.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Once()
	t.secondExtension.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Once()
	t.defaultProvider.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Twice()

	result, err := t.handler.HandleSubmittedForm(t.context, t.request)
	t.Equal(domain.NewFormErrorWithParent(errors.New("error")), err)
	t.Nil(result)
}

func (t *FormHandlerImplTestSuite) TestHandleSubmittedForm_PostValueProcessingError() {
	t.provider.On("GetFormData", t.context, t.request).Return(map[string]string{}, nil).Once()

	t.firstExtension.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Once()
	t.secondExtension.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Once()
	t.defaultProvider.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Twice()

	t.request.Request().Method = http.MethodPost

	result, err := t.handler.HandleSubmittedForm(t.context, t.request)
	t.Equal(domain.NewFormErrorWithParent(errors.New("missing form body")), err)
	t.Nil(result)
}

func (t *FormHandlerImplTestSuite) TestHandleSubmittedForm_DecodeError() {
	t.provider.On("GetFormData", t.context, t.request).Return(map[string]string{}, nil).Once()

	t.firstExtension.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Once()
	t.secondExtension.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Once()
	t.defaultProvider.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Twice()

	t.request.Request().Method = http.MethodPost
	t.request.Request().PostForm = url.Values{
		"first":  []string{"first"},
		"second": []string{"second"},
	}

	t.decoder.On("Decode", t.context, t.request, url.Values{
		"first":  []string{"first"},
		"second": []string{"second"},
	}, map[string]string{}).Return(nil, errors.New("error")).Once()

	result, err := t.handler.HandleSubmittedForm(t.context, t.request)
	t.Equal(domain.NewFormErrorWithParent(errors.New("error")), err)
	t.Nil(result)
}

func (t *FormHandlerImplTestSuite) TestHandleSubmittedForm_ValidateError() {
	t.provider.On("GetFormData", t.context, t.request).Return(map[string]string{}, nil).Once()

	t.firstExtension.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Once()
	t.secondExtension.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Once()
	t.defaultProvider.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Twice()

	t.request.Request().Method = http.MethodPost
	t.request.Request().PostForm = url.Values{
		"first":  []string{"first"},
		"second": []string{"second"},
	}

	t.decoder.On("Decode", t.context, t.request, url.Values{
		"first":  []string{"first"},
		"second": []string{"second"},
	}, map[string]string{}).Return(map[string]string{
		"first":  "first",
		"second": "second",
	}, nil).Once()

	t.validator.On("Validate", t.context, t.request, t.validatorProvider, map[string]string{
		"first":  "first",
		"second": "second",
	}).Return(nil, errors.New("error")).Once()

	result, err := t.handler.HandleSubmittedForm(t.context, t.request)
	t.Equal(domain.NewFormErrorWithParent(errors.New("error")), err)
	t.Nil(result)
}

func (t *FormHandlerImplTestSuite) TestHandleSubmittedForm_FormExtensionError() {
	t.request.Request().Method = http.MethodPost
	t.request.Request().PostForm = url.Values{
		"first":  []string{"first"},
		"second": []string{"second"},
	}

	t.firstExtension.On("GetFormData", t.context, t.request).Return(nil, errors.New("error")).Maybe()
	t.secondExtension.On("GetFormData", t.context, t.request).Return(nil, errors.New("error")).Maybe()
	t.defaultProvider.On("GetFormData", t.context, t.request).Return(nil, errors.New("error")).Maybe()

	result, err := t.handler.HandleSubmittedForm(t.context, t.request)
	t.Equal(domain.NewFormErrorWithParent(errors.New("error")), err)
	t.Nil(result)
}

func (t *FormHandlerImplTestSuite) TestHandleSubmittedForm_Success() {
	t.provider.On("GetFormData", t.context, t.request).Return(map[string]string{}, nil).Once()

	t.request.Request().Method = http.MethodPost
	t.request.Request().PostForm = url.Values{
		"first":  []string{"first"},
		"second": []string{"second"},
	}

	t.decoder.On("Decode", t.context, t.request, url.Values{
		"first":  []string{"first"},
		"second": []string{"second"},
	}, map[string]string{}).Return(map[string]string{
		"first":  "first",
		"second": "second",
	}, nil).Once()

	t.validator.On("Validate", t.context, t.request, t.validatorProvider, map[string]string{
		"first":  "first",
		"second": "second",
	}).Return(&domain.ValidationInfo{}, nil).Once()

	t.firstExtension.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Twice()
	t.firstExtension.On("Decode", t.context, t.request, url.Values{
		"first":  []string{"first"},
		"second": []string{"second"},
	}, map[string]int{}).Return(map[string]int{}, nil).Once()
	t.firstExtension.On("Validate", t.context, t.request, t.validatorProvider, map[string]int{}).Return(&domain.ValidationInfo{}, nil).Once()

	t.secondExtension.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Twice()
	t.defaultDecoder.On("Decode", t.context, t.request, url.Values{
		"first":  []string{"first"},
		"second": []string{"second"},
	}, map[string]int{}).Return(map[string]int{}, nil).Once()
	t.defaultValidator.On("Validate", t.context, t.request, t.validatorProvider, map[string]int{}).Return(&domain.ValidationInfo{}, nil).Once()

	t.defaultProvider.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Twice()
	t.thirdExtension.On("Decode", t.context, t.request, url.Values{
		"first":  []string{"first"},
		"second": []string{"second"},
	}, map[string]int{}).Return(map[string]int{}, nil).Once()
	t.defaultValidator.On("Validate", t.context, t.request, t.validatorProvider, map[string]int{}).Return(&domain.ValidationInfo{}, nil).Once()

	t.defaultProvider.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Twice()
	t.defaultDecoder.On("Decode", t.context, t.request, url.Values{
		"first":  []string{"first"},
		"second": []string{"second"},
	}, map[string]int{}).Return(map[string]int{}, nil).Once()
	t.fourthExtension.On("Validate", t.context, t.request, t.validatorProvider, map[string]int{}).Return(&domain.ValidationInfo{}, nil).Once()

	result, err := t.handler.HandleSubmittedForm(t.context, t.request)
	t.NoError(err)

	form := domain.NewForm(true, map[string][]domain.ValidationRule{})
	form.Data = map[string]string{
		"first":  "first",
		"second": "second",
	}
	form.FormExtensionsData = map[string]interface{}{
		"first":  map[string]int{},
		"second": map[string]int{},
		"third":  map[string]int{},
		"fourth": map[string]int{},
	}

	t.Equal(&form, result)
}

func (t *FormHandlerImplTestSuite) TestHandleForm_Unsubmitted() {
	t.provider.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Once()

	t.firstExtension.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Twice()
	t.secondExtension.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Twice()
	t.defaultProvider.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Times(4)

	result, err := t.handler.HandleUnsubmittedForm(t.context, t.request)
	t.NoError(err)

	form := domain.NewForm(false, map[string][]domain.ValidationRule{})
	form.Data = map[string]int{}
	form.FormExtensionsData = map[string]interface{}{
		"first":  map[string]int{},
		"second": map[string]int{},
		"third":  map[string]int{},
		"fourth": map[string]int{},
	}

	t.Equal(&form, result)
}

func (t *FormHandlerImplTestSuite) TestHandleForm_Submitted() {
	t.provider.On("GetFormData", t.context, t.request).Return(map[string]string{}, nil).Once()

	t.request.Request().Method = http.MethodPost
	t.request.Request().PostForm = url.Values{
		"first":  []string{"first"},
		"second": []string{"second"},
	}

	t.decoder.On("Decode", t.context, t.request, url.Values{
		"first":  []string{"first"},
		"second": []string{"second"},
	}, map[string]string{}).Return(map[string]string{
		"first":  "first",
		"second": "second",
	}, nil).Once()

	t.validator.On("Validate", t.context, t.request, t.validatorProvider, map[string]string{
		"first":  "first",
		"second": "second",
	}).Return(&domain.ValidationInfo{}, nil).Once()

	t.firstExtension.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Twice()
	t.firstExtension.On("Decode", t.context, t.request, url.Values{
		"first":  []string{"first"},
		"second": []string{"second"},
	}, map[string]int{}).Return(map[string]int{}, nil).Once()
	t.firstExtension.On("Validate", t.context, t.request, t.validatorProvider, map[string]int{}).Return(&domain.ValidationInfo{}, nil).Once()

	t.secondExtension.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Twice()
	t.defaultDecoder.On("Decode", t.context, t.request, url.Values{
		"first":  []string{"first"},
		"second": []string{"second"},
	}, map[string]int{}).Return(map[string]int{}, nil).Once()
	t.defaultValidator.On("Validate", t.context, t.request, t.validatorProvider, map[string]int{}).Return(&domain.ValidationInfo{}, nil).Once()

	t.defaultProvider.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Twice()
	t.thirdExtension.On("Decode", t.context, t.request, url.Values{
		"first":  []string{"first"},
		"second": []string{"second"},
	}, map[string]int{}).Return(map[string]int{}, nil).Once()
	t.defaultValidator.On("Validate", t.context, t.request, t.validatorProvider, map[string]int{}).Return(&domain.ValidationInfo{}, nil).Once()

	t.defaultProvider.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Twice()
	t.defaultDecoder.On("Decode", t.context, t.request, url.Values{
		"first":  []string{"first"},
		"second": []string{"second"},
	}, map[string]int{}).Return(map[string]int{}, nil).Once()
	t.fourthExtension.On("Validate", t.context, t.request, t.validatorProvider, map[string]int{}).Return(&domain.ValidationInfo{}, nil).Once()

	result, err := t.handler.HandleForm(t.context, t.request)
	t.NoError(err)

	form := domain.NewForm(true, map[string][]domain.ValidationRule{})
	form.Data = map[string]string{
		"first":  "first",
		"second": "second",
	}
	form.FormExtensionsData = map[string]interface{}{
		"first":  map[string]int{},
		"second": map[string]int{},
		"third":  map[string]int{},
		"fourth": map[string]int{},
	}

	t.Equal(&form, result)
}

func (t *FormHandlerImplTestSuite) TestHandleSubmittedGETForm() {
	t.provider.On("GetFormData", t.context, t.request).Return(map[string]string{}, nil).Once()

	t.request.Request().Method = http.MethodGet
	t.request.Request().URL = &url.URL{
		RawQuery: url.Values{
			"first":  []string{"first"},
			"second": []string{"second"},
		}.Encode(),
	}

	t.decoder.On("Decode", t.context, t.request, url.Values{
		"first":  []string{"first"},
		"second": []string{"second"},
	}, map[string]string{}).Return(map[string]string{
		"first":  "first",
		"second": "second",
	}, nil).Once()

	t.validator.On("Validate", t.context, t.request, t.validatorProvider, map[string]string{
		"first":  "first",
		"second": "second",
	}).Return(&domain.ValidationInfo{}, nil).Once()

	t.firstExtension.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Twice()
	t.firstExtension.On("Decode", t.context, t.request, url.Values{
		"first":  []string{"first"},
		"second": []string{"second"},
	}, map[string]int{}).Return(map[string]int{}, nil).Once()
	t.firstExtension.On("Validate", t.context, t.request, t.validatorProvider, map[string]int{}).Return(&domain.ValidationInfo{}, nil).Once()

	t.secondExtension.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Twice()
	t.defaultDecoder.On("Decode", t.context, t.request, url.Values{
		"first":  []string{"first"},
		"second": []string{"second"},
	}, map[string]int{}).Return(map[string]int{}, nil).Once()
	t.defaultValidator.On("Validate", t.context, t.request, t.validatorProvider, map[string]int{}).Return(&domain.ValidationInfo{}, nil).Once()

	t.defaultProvider.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Twice()
	t.thirdExtension.On("Decode", t.context, t.request, url.Values{
		"first":  []string{"first"},
		"second": []string{"second"},
	}, map[string]int{}).Return(map[string]int{}, nil).Once()
	t.defaultValidator.On("Validate", t.context, t.request, t.validatorProvider, map[string]int{}).Return(&domain.ValidationInfo{}, nil).Once()

	t.defaultProvider.On("GetFormData", t.context, t.request).Return(map[string]int{}, nil).Twice()
	t.defaultDecoder.On("Decode", t.context, t.request, url.Values{
		"first":  []string{"first"},
		"second": []string{"second"},
	}, map[string]int{}).Return(map[string]int{}, nil).Once()
	t.fourthExtension.On("Validate", t.context, t.request, t.validatorProvider, map[string]int{}).Return(&domain.ValidationInfo{}, nil).Once()

	result, err := t.handler.HandleSubmittedGETForm(t.context, t.request)
	t.NoError(err)

	form := domain.NewForm(true, map[string][]domain.ValidationRule{})
	form.Data = map[string]string{
		"first":  "first",
		"second": "second",
	}
	form.FormExtensionsData = map[string]interface{}{
		"first":  map[string]int{},
		"second": map[string]int{},
		"third":  map[string]int{},
		"fourth": map[string]int{},
	}

	t.Equal(&form, result)
}
