package game

import (
	"github.com/thebrubaker/colony/actions"
	"github.com/thebrubaker/colony/colonist"
	"github.com/thebrubaker/colony/commands"
	"github.com/thebrubaker/colony/region"
)

type GameState struct {
	Ticker    *Ticker
	Region    *region.Region
	Colonists []*colonist.Colonist
	Actions   []*actions.Action
	Commands  []commands.Command
}

func CreateGameState() *GameState {
	t := &Ticker{
		Rate: PausedRate,
	}
	r := &region.Region{}
	c := []*colonist.Colonist{
		colonist.GenerateColonist("Joel"),
	}
	a := actions.InitActions(r, c)
	return &GameState{
		Ticker:    t,
		Region:    r,
		Colonists: c,
		Actions:   a,
	}
}

func (gs *GameState) update(timeElapsed float64) {
	tickElapsed := timeElapsed * float64(gs.Ticker.Rate)
	actions.Update(tickElapsed, gs.Region, gs.Colonists, gs.Actions)
	gs.Ticker.Count += tickElapsed
}
