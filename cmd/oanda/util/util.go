package util

import (
	"fmt"
	"os"

	"github.com/hsmtkk/oanda_api_go/pkg/oanda"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	accountIDKey = "OANDA_ACCOUNT_ID"
	tokenKey     = "OANDA_TOKEN"
)

func NewExecutor() (oanda.Executor, error) {
	accountID, err := loadEnvVar(accountIDKey)
	if err != nil {
		return nil, err
	}
	token, err := loadEnvVar(tokenKey)
	if err != nil {
		return nil, err
	}
	logger := logrus.New()
	if viper.GetBool("verbose") {
		logger.SetLevel(logrus.DebugLevel)
		logger.SetOutput(os.Stdout)
	}
	return oanda.New(oanda.FxTradePractice, accountID, token, logger)
}

func loadEnvVar(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("environment variable %s is not defined", key)
	}
	return val, nil
}
