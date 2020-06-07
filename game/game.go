package game

import (
	"github.com/thebrubaker/colony/actions"
	"github.com/thebrubaker/colony/colonist"
	"github.com/thebrubaker/colony/region"
	"github.com/thebrubaker/colony/ticker"
)

type Game struct {
	Name      string
	Ticker    *ticker.Ticker
	Region    *region.Region
	Colonists []*colonist.Colonist
	Actions   []*actions.Action
}

func CreateGame(name string) *Game {
	t := ticker.CreateTick()

	r := &region.Region{}

	c := []*colonist.Colonist{
		colonist.GenerateColonist(name),
	}

	a := actions.InitActions(r, c)

	return &Game{
		Name:      name,
		Ticker:    t,
		Region:    r,
		Colonists: c,
		Actions:   a,
	}
}

func (g *Game) SetSpeed(r ticker.TickRate) {
	g.Ticker.SetTickRate(r)
}

func (g *Game) Start() {
	g.Ticker.OnTick(func(t *ticker.Ticker) {
		actions.Update(t.TickElapsed, g.Region, g.Colonists, g.Actions)
	})
}
