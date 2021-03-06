package detail

import (
	"net/http"
	"os"

	"github.com/hsmtkk/oanda_api_go/pkg/comm"
	"github.com/hsmtkk/oanda_api_go/pkg/env"
	"github.com/sirupsen/logrus"
)

type Getter interface {
	Detail(accountID string) (Detail, error)
}

func New(env env.Environment, token string, logger *logrus.Logger) (Getter, error) {
	communicator, err := comm.New(env, token, logger)
	if err != nil {
		return nil, err
	}
	return &getterImpl{communicator: communicator}, nil
}

func NewForTest(client *http.Client, url string) Getter {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetOutput(os.Stdout)
	logger.SetReportCaller(true)
	communicator := comm.NewForTest(client, url, "test", logger)
	return &getterImpl{communicator: communicator}
}

type getterImpl struct {
	communicator comm.Communicator
}

func (getter *getterImpl) Detail(accountID string) (Detail, error) {
	respBytes, err := getter.communicator.Get("/v3/accounts/"+accountID, nil)
	if err != nil {
		return Detail{}, err
	}
	return Parse(respBytes)
}
