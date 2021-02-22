package game

import (
	"github.com/thebrubaker/colony/actions"
	"github.com/thebrubaker/colony/colonist"
	"github.com/thebrubaker/colony/region"
)

type GameState struct {
	Ticker    *Ticker
	Region    *region.Region
	Colonists []*colonist.Colonist
	Actions   actions.Actions
}

// Create a new GameState
func CreateGameState() *GameState {
	colonists := []*colonist.Colonist{
		colonist.NewColonist("John Wallis"),
	}
	return &GameState{
		Ticker: &Ticker{
			Rate: BaseTickRate,
		},
		Region:    &region.Region{},
		Colonists: colonists,
		Actions:   actions.CreateActions(colonists),
	}
}

func (gs *GameState) update(timeElapsed float64) {
	if gs.Ticker.Rate == PausedRate {
		return
	}
	tickElapsed := timeElapsed * float64(gs.Ticker.Rate)
	gs.Actions.Update(&actions.Context{
		Region:      gs.Region,
		Colonists:   gs.Colonists,
		TickElapsed: tickElapsed,
	})
	gs.Ticker.Count += tickElapsed
}
