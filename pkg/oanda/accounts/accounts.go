package accounts

import "encoding/json"

type Response struct {
	ID       string
	Currency string
}

type responseSchema struct {
	Account struct {
		ID       string `json:"id"`
		Currency string `json:"currency"`
	} `json:"account"`
}

func Parse(jsBytes []byte) (Response, error) {
	sch := responseSchema{}
	if err := json.Unmarshal(jsBytes, &sch); err != nil {
		return Response{}, err
	}
	return Response{ID: sch.Account.ID, Currency: sch.Account.Currency}, nil
}
