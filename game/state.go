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
	actions   []*actions.Action
}

func CreateGameState() *GameState {
	t := &Ticker{
		Rate: BaseTickRate,
	}
	r := &region.Region{}
	c := []*colonist.Colonist{
		colonist.NewColonist("Joel"),
	}
	a := actions.InitActions(r, c)
	return &GameState{
		Ticker:    t,
		Region:    r,
		Colonists: c,
		actions:   a,
	}
}

func (gs *GameState) update(timeElapsed float64) {
	if gs.Ticker.Rate == PausedRate {
		return
	}
	tickElapsed := timeElapsed * float64(gs.Ticker.Rate)
	actions.Update(tickElapsed, gs.Region, gs.Colonists, gs.actions)
	gs.Ticker.Count += tickElapsed
}
