package controllers

import (
	"log"
	"time"

	"github.com/rs/xid"
	"github.com/thebrubaker/colony/game"
)

type GameController struct {
	streams *StreamController
	games   map[game.GameKey]*game.Game
	actionc chan func()
	quitc   chan struct{}
}

func NewGameController(streams *StreamController) *GameController {
	gc := &GameController{
		streams: streams,
		games:   make(map[game.GameKey]*game.Game),
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
			for key, game := range gc.games {
				gc.streams.Broadcast(key, game.Render())
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

func (gc *GameController) CreateGame() game.GameKey {
	c := make(chan game.GameKey)
	gc.actionc <- func() {
		key := NewGameKey()
		log.Printf("creating game %s", key)
		gc.games[key] = game.CreateGame()
		c <- key
	}
	return <-c
}

func (gc *GameController) SendCommand(key game.GameKey, commandType string) bool {
	c := make(chan bool)
	gc.actionc <- func() {
		log.Printf("command %s sent to game %s", commandType, key)
		g := gc.games[game.GameKey(key)]
		g.AddCommand(commandType)
		c <- true
	}
	return <-c
}

func NewGameKey() game.GameKey {
	return game.GameKey(xid.New().String())
}
