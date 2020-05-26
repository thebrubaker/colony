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
	Inventory      Stock
}

func (gameState *GameState) Update() {
	colonist := gameState.ActiveColonist

	colonist.DetermineAction([]*ActionType{
		SearchForWater,
		WatchClouds,
		StandIdle,
		DrinkWater,
	})

	colonist.OnTick()
}

func (gameState *GameState) Render() {
	colonist := gameState.ActiveColonist
	data := struct {
		Tick      string
		Name      string
		Status    string
		Stress    string
		Thirst    string
		Inventory Stock
	}{
		fmt.Sprintf("%d", uint(gameState.Ticker.Count)),
		colonist.Name,
		colonist.CurrentAction.Type.Status,
		fmt.Sprintf("%f", colonist.Stress),
		fmt.Sprintf("%f", colonist.Thirst),
		gameState.Inventory,
	}
	output, _ := json.MarshalIndent(data, "", "    ")
	fmt.Printf("\033c%s\n", string(output))
}

type Stock map[string]uint

func InitGame() *GameState {
	game := &GameState{
		Ticker: &Ticker{
			LastTick: time.Now(),
			Elapsed:  0,
			Count:    0,
		},
		Inventory: Stock{
			"water": 0,
		},
	}

	colonist := Colonist{
		Name:   "Artokun",
		Game:   game,
		Thirst: 40,
		Stress: 0,
	}

	colonist.SetActionType(WakingUp)

	game.ActiveColonist = &colonist

	return game
}
