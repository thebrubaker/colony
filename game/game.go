package game

import (
	"time"
)

type Game struct {
	state   *GameState
	actionc chan func()
	quitc   chan struct{}
}

func CreateGame() *Game {
	g := &Game{
		state:   CreateGameState(),
		actionc: make(chan func()),
		quitc:   make(chan struct{}),
	}
	go g.loop()
	return g
}

func (g *Game) loop() {
	previousTime := time.Now()

	for {
		select {
		case <-g.quitc:
			return
		case f := <-g.actionc:
			f()
		case <-time.Tick(10 * time.Millisecond):
			currentTime := time.Now()
			tickElapsed := currentTime.Sub(previousTime).Seconds()
			g.state.update(tickElapsed)
			previousTime = currentTime
		}
	}
}

func (g *Game) Stop() {
	close(g.quitc)
}

func (g *Game) AddCommand(commandType string) bool {
	c := make(chan bool)
	g.actionc <- func() {
		c <- true
	}
	return <-c
}

func (g *Game) SetTickRate(rate TickRate) bool {
	c := make(chan bool)
	g.actionc <- func() {
		g.state.Ticker.Rate = rate

		c <- true
	}
	return <-c
}

func (g *Game) GetTickRate(rate TickRate) TickRate {
	c := make(chan TickRate)
	g.actionc <- func() {
		c <- g.state.Ticker.Rate
	}
	return <-c
}

func (g *Game) Render() GameState {
	c := make(chan GameState)
	g.actionc <- func() {
		c <- *g.state
	}
	return <-c
}
