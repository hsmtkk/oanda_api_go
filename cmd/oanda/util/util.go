package util

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	accountIDKey = "OANDA_ACCOUNT_ID"
	tokenKey     = "OANDA_TOKEN"
)

func GetToken() (string, error) {
	return getEnvVar(tokenKey)
}

func GetLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	if viper.GetBool("verbose") {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.ErrorLevel)
	}
	return logger
}

func getEnvVar(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("environment variable %s is not defined", key)
	}
	return val, nil
}
