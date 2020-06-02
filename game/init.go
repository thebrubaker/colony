// package game

// import (
// 	"time"
// )

// func InitGame() *GameState {
// 	gameState := &GameState{
// 		Ticker:    &Ticker{time.Now(), 0, 0},
// 		Stockpile: &Inventory{},
// 	}

// 	gameState.Colonists = []*Colonist{
// 		MakeColonist("Blackthorne", gameState),
// 	}

// 	return gameState
// }

// func MakeColonist(name string, gameState *GameState) *Colonist {
// 	colonist := &Colonist{
// 		Name:       name,
// 		Exhaustion: &Attribute{0},
// 		Hunger:     &Attribute{60},
// 		Thirst:     &Attribute{60},
// 		Stress:     &Attribute{0},
// 		Equipment:  &Inventory{},
// 		GameState:  gameState,
// 	}

// 	colonist.SetActionType(WakingUp, gameState)

// 	return colonist
// }
