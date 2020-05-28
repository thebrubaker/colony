package game

import "time"

func InitGame() *GameState {
	gameState := &GameState{
		Ticker:    &Ticker{time.Now(), 0, 0},
		Inventory: &Inventory{},
	}

	colonist := Colonist{
		Name:      "Caloke",
		Energy:    &Attribute{100},
		Hunger:    &Attribute{40},
		Thirst:    &Attribute{40},
		Stress:    &Attribute{0},
		Inventory: &Inventory{},
		GameState: gameState,
	}

	colonist.SetActionType(WakingUp, gameState)

	gameState.ActiveColonist = &colonist

	return gameState
}
