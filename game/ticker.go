package game

type TickRate float64

const (
	PausedRate      TickRate = 0
	BaseTickRate    TickRate = 1
	FastTickRate    TickRate = 2
	FastestTickRate TickRate = 3
)

type Ticker struct {
	Rate  TickRate
	Count float64
}
