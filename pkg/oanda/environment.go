package oanda

type Environment int

const (
	UnknownEnvironment Environment = iota
	Sandbox
	FxTradePractice
	FxTrade
)

func (env Environment) String() string {
	switch env {
	case Sandbox:
		return "Sandbox"
	case FxTradePractice:
		return "FxTradePractice"
	case FxTrade:
		return "FxTrade"
	}
	return "unknown"
}
