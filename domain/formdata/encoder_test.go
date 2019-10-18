package formdata

import (
	"github.com/stretchr/testify/suite"
	"net/url"
	"testing"
)

type (
	DefaultFormDataEncoderImplTestSuite struct {
		suite.Suite

		encoder *DefaultFormDataEncoderImpl
	}

	formDataEncoderTestData struct {
		Text   string    `form:"text" conform:"trim"`
		Number int       `form:"number"`
		Slice  []float64 `form:"slice"`
	}
)

func TestDefaultFormDataEncoderImplTestSuite(t *testing.T) {
	suite.Run(t, &DefaultFormDataEncoderImplTestSuite{})
}

func (t *DefaultFormDataEncoderImplTestSuite) SetupSuite() {
	t.encoder = &DefaultFormDataEncoderImpl{}
}

func (t *DefaultFormDataEncoderImplTestSuite) TestEncodeUnknownInterface_Error() {
	urlValues, err := t.encoder.encodeUnknownInterface(nil)

	t.Error(err)
	t.Equal(url.Values(nil), urlValues)
}

func (t *DefaultFormDataEncoderImplTestSuite) TestEncodeUnknownInterface_Empty() {
	urlValues, err := t.encoder.encodeUnknownInterface(url.Values{})

	t.NoError(err)
	t.Equal(url.Values{}, urlValues)
}

func (t *DefaultFormDataEncoderImplTestSuite) TestEncodeUnknownInterface_Struct() {
	urlValues, err := t.encoder.encodeUnknownInterface(formDataEncoderTestData{
		Number: 1,
		Text:   "string",
		Slice:  []float64{1, 2},
	})

	t.NoError(err)
	t.Equal(url.Values{
		"number": []string{"1"},
		"text":   []string{"string"},
		"slice":  []string{"1", "2"},
	}, urlValues)
}

func (t *DefaultFormDataEncoderImplTestSuite) TestEncode_Error() {
	urlValues, err := t.encoder.Encode(nil, nil)

	t.Error(err)
	t.Equal(url.Values(nil), urlValues)
}

func (t *DefaultFormDataEncoderImplTestSuite) TestEncode_Empty() {
	urlValues, err := t.encoder.Encode(nil, url.Values{})

	t.NoError(err)
	t.Equal(url.Values{}, urlValues)
}

func (t *DefaultFormDataEncoderImplTestSuite) TestEncode_Struct() {
	urlValues, err := t.encoder.Encode(nil, formDataEncoderTestData{
		Number: 1,
		Text:   "string",
		Slice:  []float64{1, 2},
	})

	t.NoError(err)
	t.Equal(url.Values{
		"number": []string{"1"},
		"text":   []string{"string"},
		"slice":  []string{"1", "2"},
	}, urlValues)
}

func (t *DefaultFormDataEncoderImplTestSuite) TestEncode_Map() {
	urlValues, err := t.encoder.Encode(nil, map[string]string{
		"number": "1",
		"text":   "string",
		"slice":  "1",
	})

	t.NoError(err)
	t.Equal(url.Values{
		"number": []string{"1"},
		"text":   []string{"string"},
		"slice":  []string{"1"},
	}, urlValues)
}
