package game

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/thebrubaker/colony/keys"
)

type Broadcaster interface {
	Broadcast(keys.GameKey, GameState)
}

type GameController struct {
	streams Broadcaster
	games   map[keys.GameKey]*Game
	actionc chan func()
	quitc   chan struct{}
}

func NewController(streams Broadcaster) *GameController {
	gc := &GameController{
		streams: streams,
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
		case <-time.Tick(33 * time.Millisecond):
			for key, g := range gc.games {
				data, err := json.MarshalIndent(g.Render(), "", "    ")
				if err != nil {
					return
				}
				fmt.Println(string(data))

				gc.streams.Broadcast(key, g.Render())
			}
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
		log.Printf("set game %s speed to %f", key, r)
		g := gc.games[keys.GameKey(key)]
		g.SetTickRate(r)
		c <- true
	}
	return <-c
}
