package form_test

import (
	"testing"

	"flamingo.me/dingo"
	"flamingo.me/form"
)

func TestModule_Configure(t *testing.T) {
	if err := dingo.TryModule(new(form.Module)); err != nil {
		t.Error(err)
	}
}
