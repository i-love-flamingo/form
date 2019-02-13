package formdata

import (
	"testing"

	"flamingo.me/flamingo/v3/framework/web"
	"flamingo.me/form/domain"
	"flamingo.me/form/domain/mocks"
	"github.com/stretchr/testify/suite"
)

type (
	DefaultFormDataValidatorImplTestSuite struct {
		suite.Suite

		validator *DefaultFormDataValidatorImpl

		validatorProvider *mocks.ValidatorProvider
	}
)

func TestDefaultFormDataValidatorImplTestSuite(t *testing.T) {
	suite.Run(t, &DefaultFormDataValidatorImplTestSuite{})
}

func (t *DefaultFormDataValidatorImplTestSuite) SetupSuite() {
	t.validator = &DefaultFormDataValidatorImpl{}
}

func (t *DefaultFormDataValidatorImplTestSuite) SetupTest() {
	t.validatorProvider = &mocks.ValidatorProvider{}
}

func (t *DefaultFormDataValidatorImplTestSuite) TearDownTest() {
	t.validatorProvider.AssertExpectations(t.T())
	t.validatorProvider = nil
}

func (t *DefaultFormDataValidatorImplTestSuite) TestGetFormData() {
	t.validatorProvider.On("Validate", nil, (*web.Request)(nil), "something").Return(domain.ValidationInfo{}).Once()

	result, err := t.validator.Validate(nil, nil, t.validatorProvider, "something")

	t.NoError(err)
	t.Equal(&domain.ValidationInfo{}, result)
}
