package env

import "fmt"

type Environment int

const (
	UnknownEnvironment Environment = iota
	FxTradePractice
	FxTrade
)

func (env Environment) String() string {
	switch env {
	case FxTradePractice:
		return "FxTradePractice"
	case FxTrade:
		return "FxTrade"
	}
	return "unknown"
}

func (env Environment) URL() (string, error) {
	switch env {
	case FxTradePractice:
		return "https://api-fxpractice.oanda.com", nil
	case FxTrade:
		return "https://api-fxtrade.oanda.com", nil
	}
	return "", fmt.Errorf("unknown environment %s", env.String())
}
