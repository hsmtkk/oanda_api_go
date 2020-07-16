package detail

import "encoding/json"

type schema struct {
	Account Detail `json:"account"`
}

type Detail struct {
	ID       string `json:"id"`
	Currency string `json:"currency"`
}

func Parse(respBody []byte) (Detail, error) {
	sch := schema{}
	if err := json.Unmarshal(respBody, &sch); err != nil {
		return Detail{}, err
	}
	return sch.Account, nil
}
