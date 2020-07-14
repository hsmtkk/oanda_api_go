package oanda

func (executor *executorImpl) Instruments() (string, error) {
	query := map[string]string{"accountId": executor.accountID}
	bs, err := executor.HTTPGet("/v1/instruments", query)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
