package env_test

import (
	"testing"

	"github.com/hsmtkk/oanda_api_go/pkg/env"
	"github.com/hsmtkk/oanda_api_go/test"
)

func TestEnvURL(t *testing.T) {
	want := "https://api-fxpractice.oanda.com"
	e := env.FxTradePractice
	got, err := e.URL()
	test.AssertNil(t, err)
	test.AssertEqualString(t, want, got)
}
