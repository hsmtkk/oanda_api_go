package list_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hsmtkk/oanda_api_go/pkg/accounts/list"
	"github.com/hsmtkk/oanda_api_go/test"
)

func TestList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		respBytes, err := ioutil.ReadFile("list.json")
		test.AssertNil(t, err)
		w.Write(respBytes)
	}))
	defer ts.Close()

	getter := list.NewForTest(ts.Client(), ts.URL)
	want := []string{"012"}
	got, err := getter.List()
	test.AssertNil(t, err)
	test.AssertEqualStrings(t, want, got)
}
