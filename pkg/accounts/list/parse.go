package list

import "encoding/json"

type schema struct {
	Accounts []account `json:"accounts"`
}

type account struct {
	ID string `json:"id"`
}

func Parse(respBody []byte) ([]string, error) {
	sch := schema{}
	if err := json.Unmarshal(respBody, &sch); err != nil {
		return nil, err
	}
	ids := []string{}
	for _, ac := range sch.Accounts {
		ids = append(ids, ac.ID)
	}
	return ids, nil
}
