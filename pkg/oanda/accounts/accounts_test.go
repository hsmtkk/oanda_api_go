package accounts_test

import (
	"testing"

	"github.com/hsmtkk/oanda_api_go/pkg/oanda/accounts"
	"github.com/hsmtkk/oanda_api_go/test"
)

func TestParse(t *testing.T) {
	input := []byte(`{"account":{"guaranteedStopLossOrderMode":"DISABLED","hedgingEnabled":false,"id":"101-009-15693939-001","createdTime":"2020-07-14T03:55:01.325486361Z","currency":"JPY","createdByUserID":15693939,"alias":"Primary","marginRate":"0.04","lastTransactionID":"3","balance":"3000000.0000","openTradeCount":0,"openPositionCount":0,"pendingOrderCount":0,"pl":"0.0000","resettablePL":"0.0000","resettablePLTime":"0","financing":"0.0000","commission":"0.0000","dividendAdjustment":"0","guaranteedExecutionFees":"0.0000","orders":[],"positions":[],"trades":[],"unrealizedPL":"0.0000","NAV":"3000000.0000","marginUsed":"0.0000","marginAvailable":"3000000.0000","positionValue":"0.0000","marginCloseoutUnrealizedPL":"0.0000","marginCloseoutNAV":"3000000.0000","marginCloseoutMarginUsed":"0.0000","marginCloseoutPositionValue":"0.0000","marginCloseoutPercent":"0.00000","withdrawalLimit":"3000000.0000"},"lastTransactionID":"3"}`)
	want := accounts.Response{ID: "101-009-15693939-001", Currency: "JPY"}
	got, err := accounts.Parse(input)
	test.AssertNil(t, err)
	assertEqualResponse(t, want, got)
}

func assertEqualResponse(t *testing.T, want, got accounts.Response) {
	t.Helper()
	if want.ID != got.ID || want.Currency != got.Currency {
		t.Errorf("want %v, got %v", want, got)
	}
}
