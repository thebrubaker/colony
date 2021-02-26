package game

import (
	"log"
	"math"

	"github.com/thebrubaker/colony/keys"
)

type GameController struct {
	games   map[keys.GameKey]*Game
	actionc chan func()
	quitc   chan struct{}
}

func NewController() *GameController {
	gc := &GameController{
		games:   make(map[keys.GameKey]*Game),
		actionc: make(chan func()),
		quitc:   make(chan struct{}),
	}
	go gc.loop()
	return gc
}

func (gc *GameController) loop() {
	for {
		select {
		case <-gc.quitc:
			return
		case f := <-gc.actionc:
			f()
		}
	}
}

func (gc *GameController) Stop() {
	for key, g := range gc.games {
		log.Printf("stopping game %s", key)
		g.Stop()
		delete(gc.games, key)
	}
	close(gc.quitc)
}

func (gc *GameController) CreateGame() keys.GameKey {
	c := make(chan keys.GameKey)
	gc.actionc <- func() {
		key := keys.NewGameKey()
		log.Printf("creating game %s", key)
		gc.games[key] = CreateGame()
		c <- key
	}
	return <-c
}

func (gc *GameController) SendCommand(key keys.GameKey, commandType string) bool {
	c := make(chan bool)
	gc.actionc <- func() {
		log.Printf("command %s sent to game %s", commandType, key)
		g := gc.games[keys.GameKey(key)]
		g.AddCommand(commandType)
		c <- true
	}
	return <-c
}

func (gc *GameController) SetSpeed(key keys.GameKey, r TickRate) bool {
	c := make(chan bool)
	gc.actionc <- func() {
		g := gc.games[key]
		g.SetTickRate(r)
		c <- true
	}
	return <-c
}

func (gc *GameController) IncreaseSpeed(key keys.GameKey) bool {
	c := make(chan bool)
	gc.actionc <- func() {
		g := gc.games[key]
		rate := g.state.Ticker.Rate
		newRate := math.Min(float64(FastestTickRate), float64(rate)+1)
		g.SetTickRate(TickRate(newRate))
		c <- true
	}
	return <-c
}

func (gc *GameController) DecreaseSpeed(key keys.GameKey) bool {
	c := make(chan bool)
	gc.actionc <- func() {
		g := gc.games[key]
		rate := g.state.Ticker.Rate
		newRate := math.Max(float64(BaseTickRate), float64(rate)-1)
		g.SetTickRate(TickRate(newRate))
		c <- true
	}
	return <-c
}

func (gc *GameController) Render(key keys.GameKey) GameState {
	c := make(chan GameState)
	gc.actionc <- func() {
		c <- gc.games[key].Render()
	}
	return <-c
}
