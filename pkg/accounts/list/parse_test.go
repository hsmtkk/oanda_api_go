package list_test

import (
	"io/ioutil"
	"testing"

	"github.com/hsmtkk/oanda_api_go/pkg/accounts/list"
	"github.com/hsmtkk/oanda_api_go/test"
)

func TestParse(t *testing.T) {
	respBody, err := ioutil.ReadFile("list.json")
	test.AssertNil(t, err)
	want := []string{"012"}
	got, err := list.Parse(respBody)
	test.AssertNil(t, err)
	test.AssertEqualStrings(t, want, got)
}
