package oanda

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"

	"github.com/sirupsen/logrus"
)

func (executor *executorImpl) HTTPGet(urlSuffix string, query map[string]string) ([]byte, error) {
	url := executor.baseURL + urlSuffix
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+executor.personalAccessToken)
	if query != nil {
		q := req.URL.Query()
		for k, v := range query {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	reqBytes, err := httputil.DumpRequest(req, true)
	if err != nil {
		return nil, err
	}
	executor.logger.WithFields(logrus.Fields{"request": string(reqBytes)}).Debug()

	resp, err := executor.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return nil, err
	}
	executor.logger.WithFields(logrus.Fields{"response": string(respBytes)}).Debug()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("got HTTP status code %d; %s", resp.StatusCode, string(bs))
	}
	return bs, nil
}
