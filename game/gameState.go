package game

import (
	"encoding/json"
	"fmt"
	"time"
)

type Ticker struct {
	LastTick time.Time
	Elapsed  float64
	Count    float64
}

type GameState struct {
	ActiveColonist *Colonist
	Ticker         *Ticker
	Inventory      *Inventory
}

func (gameState *GameState) Update(ticker *Ticker) {
	gameState.ActiveColonist.ProcessActions(ActionTypes, gameState).OnTick(ticker)
}

func (gameState *GameState) Render() {
	data := ToJsonGame(gameState)

	output, _ := json.MarshalIndent(data, "", "    ")

	fmt.Printf("\033c%s\n", string(output))
}
