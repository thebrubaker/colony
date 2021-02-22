package actions

import (
	"github.com/thebrubaker/colony/actions/types"
	"github.com/thebrubaker/colony/colonist"
)

func (ctx *Context) StartAction(action types.Actionable) {
	if resources := action.ConsumesResources(); resources != nil {
		ctx.RemoveResources(resources)
	}
	statuses := action.Status()
	ctx.ActiveColonist.SetStatus(statuses[0])
}

func (ctx *Context) EndAction(action types.Actionable) {

}

func (ctx *Context) UpdateAction(action types.Actionable) {
	ctx.ActiveColonist.Needs.Increase(colonist.Exhaustion, float64(action.TakesEffort())*ctx.TickElapsed)
	ctx.ActiveColonist.Needs.Increase(colonist.Hunger, 0.05*ctx.TickElapsed)
	ctx.ActiveColonist.Needs.Increase(colonist.Thirst, 0.05*ctx.TickElapsed)

	// if i, ok := action.(types.SimpleFulfillment); ok {
	// 	SimpleFulfillment(c, i, a.TickProgress)
	// }

	// if i, ok := action.(types.SimpleAgitate); ok {
	// 	SimpleAgitate(c, i, a.TickProgress)
	// }

	// if i, ok := action.(types.Gathers); ok {
	// 	Gathers(c, i, ctx.ActiveColonist.TickElapsed)
	// }

	// if i, ok := action.(types.SimpleSkillUp); ok {
	// 	SimpleSkillUp(c, i, ctx.ActiveColonist.TickElapsed)
	// }
}

// func Gathers(c *Context, i types.Gathers, tickElapsed float64) {
// 	storage, elem, odds := i.Gather()

// 	// fmt.Println(rand.Float64(), tickElapsed, odds)

// 	if rand.Float64() > odds*tickElapsed {
// 		return
// 	}

// 	switch storage {
// 	case types.ColonistBag:
// 		ctx.ActiveColonist.Colonist.Bag.Add(elem, 1)
// 	case types.Stockpile:
// 		return // TODO
// 	}
// }

// func SimpleFulfillment(c *Context, i types.SimpleFulfillment, tickProgress float64) {
// 	needType, total, ease := i.Satisfies()
// 	duration := float64(i.Duration())

// 	value := GetEasedValue(total, ease, duration, tickProgress, ctx.ActiveColonist.TickElapsed)

// 	ctx.ActiveColonist.Colonist.Needs.Decrease(needType, value)
// }

// func SimpleAgitate(c *Context, i types.SimpleAgitate, tickProgress float64) {
// 	needType, total, ease := i.Agitates()
// 	duration := float64(i.Duration())

// 	value := GetEasedValue(total, ease, duration, tickProgress, ctx.ActiveColonist.TickElapsed)

// 	ctx.ActiveColonist.Colonist.Needs.Increase(needType, value)
// }

// func SimpleSkillUp(c *Context, i types.SimpleSkillUp, tickProgress float64) {
// 	skillType, total, ease := i.SkillUp()
// 	duration := float64(i.Duration())

// 	value := GetEasedValue(total, ease, duration, tickProgress, ctx.ActiveColonist.TickElapsed)

// 	ctx.ActiveColonist.Colonist.Skills.Increase(skillType, value)
// }

func (ctx *Context) RemoveResources(resources []types.ConsumeResource) {
	for _, resourceQuantity := range resources {
		ctx.ActiveColonist.Bag.Remove(resourceQuantity.Resource, resourceQuantity.Amount)
	}
}

func CheckBagSpace(b *colonist.Bag, quantity uint) bool {
	return b.GetAvailableSpace() >= quantity
}
