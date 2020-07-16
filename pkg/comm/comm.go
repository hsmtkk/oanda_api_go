package comm

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hsmtkk/oanda_api_go/pkg/env"
	"github.com/sirupsen/logrus"
)

type Communicator interface {
	Get(path string, query map[string]string) ([]byte, error)
}

type commImpl struct {
	client  *http.Client
	baseURL string
	token   string
	logger  *logrus.Logger
}

func New(env env.Environment, token string, logger *logrus.Logger) (Communicator, error) {
	url, err := env.URL()
	if err != nil {
		return nil, err
	}
	return &commImpl{client: http.DefaultClient, baseURL: url, token: token, logger: logger}, nil
}

func NewForTest(client *http.Client, baseURL string, token string, logger *logrus.Logger) Communicator {
	return &commImpl{client: client, baseURL: baseURL, logger: logger}
}

func (comm *commImpl) Get(path string, query map[string]string) ([]byte, error) {
	url := comm.baseURL + path
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+comm.token)
	resp, err := comm.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("got HTTP status code %d", resp.StatusCode)
	}
	return respBytes, nil
}
