package detail_test

import (
	"io/ioutil"
	"testing"

	"github.com/hsmtkk/oanda_api_go/pkg/accounts/detail"
	"github.com/hsmtkk/oanda_api_go/test"
)

func TestParse(t *testing.T) {
	respBody, err := ioutil.ReadFile("detail.json")
	test.AssertNil(t, err)
	want := detail.Detail{ID: "101-009-15693939-001", Currency: "JPY"}
	got, err := detail.Parse(respBody)
	test.AssertNil(t, err)
	assertEqualDetail(t, want, got)
}
