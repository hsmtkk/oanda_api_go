package oanda

import (
	"net/http"

	"github.com/hsmtkk/oanda_api_go/pkg/oanda/accounts"
	"github.com/sirupsen/logrus"
)

type Executor interface {
	Accounts() (accounts.Response, error)
	Instruments() (string, error)
}

type executorImpl struct {
	client              *http.Client
	baseURL             string
	accountID           string
	personalAccessToken string
	logger              *logrus.Logger
}

func New(env Environment, accountID, personalAccessToken string, logger *logrus.Logger) (Executor, error) {
	client := http.DefaultClient
	baseURL, err := EnvURL(env)
	if err != nil {
		return nil, err
	}
	return &executorImpl{client: client, baseURL: baseURL, accountID: accountID, personalAccessToken: personalAccessToken, logger: logger}, nil
}

func (executor *executorImpl) Accounts() (accounts.Response, error) {
	bs, err := executor.HTTPGet("/v3/accounts/"+executor.accountID, nil)
	if err != nil {
		return accounts.Response{}, err
	}
	return accounts.Parse(bs)
}
