package actions

import (
	"math/rand"

	"github.com/thebrubaker/colony/actions/types"
	"github.com/thebrubaker/colony/need"
)

func OnTick(c *Context, a *Action) {
	c.Colonist.Needs.Increase(need.Exhaustion, float64(a.Type.EnergyCost())*c.TickElapsed)
	c.Colonist.Needs.Increase(need.Hunger, 0.05*c.TickElapsed)
	c.Colonist.Needs.Increase(need.Thirst, 0.05*c.TickElapsed)

	if i, ok := a.Type.(types.OnTick); ok {
		i.OnTick()
	}

	if i, ok := a.Type.(types.SimpleFulfillment); ok {
		SimpleFulfillment(c, i, a.TickProgress)
	}

	if i, ok := a.Type.(types.SimpleAgitate); ok {
		SimpleAgitate(c, i, a.TickProgress)
	}

	if i, ok := a.Type.(types.AddsToBag); ok {
		AddsToBag(c, i, a.TickProgress)
	}
}

func AddsToBag(c *Context, i types.AddsToBag, tickProgress float64) {
	storage, elem, count, odds := i.AddToBag()

	if rand.Float64() > odds {
		return
	}

	switch storage {
	case types.ColonistBag:
		c.Colonist.Bag.Add(elem, count)
	case types.Stockpile:
		return // TODO
	}
}

func SimpleFulfillment(c *Context, i types.SimpleFulfillment, tickProgress float64) {
	needType, total, ease := i.Satisfies()
	duration := float64(i.Duration())

	need, value := GetEasedValue(needType, total, ease, duration, tickProgress, c.TickElapsed)

	c.Colonist.Needs.Decrease(need, value)
}

func SimpleAgitate(c *Context, i types.SimpleAgitate, tickProgress float64) {
	needType, total, ease := i.Agitates()
	duration := float64(i.Duration())

	need, value := GetEasedValue(needType, total, ease, duration, tickProgress, c.TickElapsed)

	c.Colonist.Needs.Increase(need, value)
}
