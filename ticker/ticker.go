package ticker

import (
	"time"
)

type TickRate string

const (
	PausedRate      TickRate = "pause"
	BaseTickRate    TickRate = "base"
	FastTickRate    TickRate = "fast"
	FastestTickRate TickRate = "fastest"
)

const (
	pausedRate      = 0
	baseTickRate    = 1
	fastTickRate    = 2
	fastestTickRate = 3
)

type Ticker struct {
	Rate           float64
	TickElapsed    float64
	TimeOfLastTick time.Time
	Count          float64
}

func (t *Ticker) SetTickRate(r TickRate) {
	switch r {
	case PausedRate:
		t.Rate = pausedRate
	case BaseTickRate:
		t.Rate = baseTickRate
	case FastTickRate:
		t.Rate = fastTickRate
	case FastestTickRate:
		t.Rate = fastestTickRate
	default:
		t.Rate = baseTickRate
	}
}

func CreateTick() *Ticker {
	return &Ticker{
		Rate:           baseTickRate,
		TickElapsed:    0,
		TimeOfLastTick: time.Now(),
		Count:          0,
	}
}

func (ticker *Ticker) OnTick(f func(t *Ticker)) {
	for {
		currentTime := time.Now()
		ticker.TickElapsed = currentTime.Sub(ticker.TimeOfLastTick).Seconds() * ticker.Rate
		f(ticker)
		ticker.TimeOfLastTick = currentTime
		ticker.Count += ticker.TickElapsed
	}
}

func OnTickMilliseconds(t time.Duration, f func()) {
	for range time.Tick(t * time.Millisecond) {
		f()
	}
}
