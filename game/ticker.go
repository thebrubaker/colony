package game

type tickRate float64

const (
	PausedRate      = 0
	BaseTickRate    = 1
	FastTickRate    = 2
	FastestTickRate = 3
)

type Ticker struct {
	Rate  tickRate
	Count float64
}
