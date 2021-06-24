package validators

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"flamingo.me/form/domain/mocks"
)

type (
	MaximumAgeValidatorTestSuite struct {
		suite.Suite

		validator *MaximumAgeValidator
	}
)

func TestMaximumAgeValidatorTestSuite(t *testing.T) {
	suite.Run(t, &MaximumAgeValidatorTestSuite{})
}

func (t *MaximumAgeValidatorTestSuite) SetupTest() {
	t.validator = &MaximumAgeValidator{}
	t.validator.Inject(&struct {
		DateFormat string `inject:"config:form.validator.dateFormat"`
	}{
		DateFormat: "2006-01-02",
	})
}

func (t *MaximumAgeValidatorTestSuite) TestValidatorName() {
	t.Equal("maximumage", t.validator.ValidatorName())
}

func (t *MaximumAgeValidatorTestSuite) TestValidateField() {
	now := time.Now()
	age149Years364Days := time.Date(now.Year()-150, now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	ageExactly150Years := time.Date(now.Year()-150, now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	age150Years364Days := time.Date(now.Year()-151, now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	ageExactly151Years := time.Date(now.Year()-151, now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	testCases := []struct {
		Name   string
		Date   string
		Result bool
	}{
		{
			Name:   "empty date is allowed",
			Date:   "",
			Result: true,
		},
		{
			Name:   "empty date with whitespace is allowed",
			Date:   " ",
			Result: true,
		},
		{
			Name:   "invalid date format is allowed",
			Date:   "wrong",
			Result: true,
		},
		{
			Name:   "age of 149 years and 364 days is allowed",
			Date:   age149Years364Days.Format("2006-01-02"),
			Result: true,
		},
		{
			Name:   "age of 150 years is allowed",
			Date:   ageExactly150Years.Format("2006-01-02"),
			Result: true,
		},
		{
			Name:   "age of 150 years and 364 days is allowed",
			Date:   age150Years364Days.Format("2006-01-02"),
			Result: true,
		},
		{
			Name:   "age of 151 years is blocked",
			Date:   ageExactly151Years.Format("2006-01-02"),
			Result: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func() {
			fieldLevel := &mocks.FieldLevel{}
			fieldLevel.On("Field").Return(reflect.ValueOf(testCase.Date)).Once()
			fieldLevel.On("Param").Return("150").Once()
			t.Equal(testCase.Result, t.validator.ValidateField(nil, fieldLevel))
			fieldLevel.AssertExpectations(t.T())
		})
	}
}
