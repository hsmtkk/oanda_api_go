package oanda

import "fmt"

func EnvURL(env Environment) (string, error) {
	switch env {
	case Sandbox:
		return "http://api-sandbox.oanda.com", nil
	case FxTradePractice:
		return "https://api-fxpractice.oanda.com", nil
	case FxTrade:
		return "https://api-fxtrade.oanda.com", nil
	}
	return "", fmt.Errorf("unknown environment %s", env.String())
}
