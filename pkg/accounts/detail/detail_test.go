package detail_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hsmtkk/oanda_api_go/pkg/accounts/detail"
	"github.com/hsmtkk/oanda_api_go/test"
)

func testList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		respBytes, err := ioutil.ReadFile("detail.json")
		test.AssertNil(t, err)
		w.Write(respBytes)
	}))
	defer ts.Close()

	getter := detail.NewForTest(ts.Client(), ts.URL)
	want := detail.Detail{ID: "101-009-15693939-001", Currency: "JPY"}
	got, err := getter.Detail("101-009-15693939-001")
	test.AssertNil(t, err)
	assertEqualDetail(t, want, got)
}

func assertEqualDetail(t *testing.T, want, got detail.Detail) {
	if want.ID != got.ID || want.Currency != got.Currency {
		t.Errorf("want %v, got %v", want, got)
	}
}
