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

	if i, ok := a.Type.(types.Gathers); ok {
		Gathers(c, i, c.TickElapsed)
	}

	if i, ok := a.Type.(types.SimpleSkillUp); ok {
		SimpleSkillUp(c, i, c.TickElapsed)
	}
}

func Gathers(c *Context, i types.Gathers, tickElapsed float64) {
	storage, elem, odds := i.Gather()

	// fmt.Println(rand.Float64(), tickElapsed, odds)

	if rand.Float64() > odds*tickElapsed {
		return
	}

	switch storage {
	case types.ColonistBag:
		c.Colonist.Bag.Add(elem, 1)
	case types.Stockpile:
		return // TODO
	}
}

func SimpleFulfillment(c *Context, i types.SimpleFulfillment, tickProgress float64) {
	needType, total, ease := i.Satisfies()
	duration := float64(i.Duration())

	value := GetEasedValue(total, ease, duration, tickProgress, c.TickElapsed)

	c.Colonist.Needs.Decrease(needType, value)
}

func SimpleAgitate(c *Context, i types.SimpleAgitate, tickProgress float64) {
	needType, total, ease := i.Agitates()
	duration := float64(i.Duration())

	value := GetEasedValue(total, ease, duration, tickProgress, c.TickElapsed)

	c.Colonist.Needs.Increase(needType, value)
}

func SimpleSkillUp(c *Context, i types.SimpleSkillUp, tickProgress float64) {
	skillType, total, ease := i.SkillUp()
	duration := float64(i.Duration())

	value := GetEasedValue(total, ease, duration, tickProgress, c.TickElapsed)

	c.Colonist.Skills.Increase(skillType, value)
}
