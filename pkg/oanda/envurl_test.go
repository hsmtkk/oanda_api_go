package oanda_test

import "testing"

func TestEnvURL(t *testing.T) {
	want := "https://api-fxpractice.oanda.com"
	got, err := oanda.EnvURL(oanda.FxTradePractice)
	test.AssertNil(t, err)
	test.AssertEqualString(t, want, got)
}
