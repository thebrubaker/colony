package ticker

import (
	"time"
)

type TickRate string

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
	case "pause":
		t.Rate = pausedRate
	case "base":
		t.Rate = baseTickRate
	case "fast":
		t.Rate = fastTickRate
	case "fastest":
		t.Rate = fastestTickRate
	default:
		t.Rate = baseTickRate
	}
}

func OnTick(f func(t *Ticker)) {
	ticker := &Ticker{
		Rate:           baseTickRate,
		TickElapsed:    0,
		TimeOfLastTick: time.Now(),
		Count:          0,
	}

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
