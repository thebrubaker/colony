package game

import (
	"encoding/json"
	"fmt"
	"time"
)

type Ticker struct {
	LastTick time.Time `json:last_tick`
	Elapsed  float64   `json:elapsed`
	Count    float64   `json:count`
}

type GameState struct {
	ActiveColonist *Colonist  `json:active_colonist`
	Ticker         *Ticker    `json:ticker`
	Inventory      *Inventory `json:inventory`
}

func (gameState *GameState) Update(ticker *Ticker) {
	gameState.ActiveColonist.ProcessActions(ActionTypes, gameState).OnTick(ticker)
}

func (gameState *GameState) Render() string {
	data := ToJsonGame(gameState)

	output, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		fmt.Println(err)
	}

	return string(output)
}
