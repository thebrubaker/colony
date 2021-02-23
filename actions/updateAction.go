package actions

import (
	"math/rand"

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

func (ctx *Context) UpdateAction(action *Action) {
	ctx.ActiveColonist.Needs.Increase(colonist.Exhaustion, float64(action.Type.TakesEffort())*ctx.TickElapsed)
	ctx.ActiveColonist.Needs.Increase(colonist.Hunger, 0.05*ctx.TickElapsed)
	ctx.ActiveColonist.Needs.Increase(colonist.Thirst, 0.05*ctx.TickElapsed)

	ctx.ActiveColonist.Desires.Decrease(colonist.Fulfillment, 0.05*ctx.TickElapsed)
	ctx.ActiveColonist.Desires.Decrease(colonist.Belonging, 0.05*ctx.TickElapsed)
	ctx.ActiveColonist.Desires.Decrease(colonist.Esteem, 0.05*ctx.TickElapsed)

	if agitateNeeds := action.Type.AgitatesNeeds(); agitateNeeds != nil {
		for _, need := range agitateNeeds {
			ctx.AgitateNeed(need, action.Type, action.TickProgress)
		}
	}

	if satisfyNeeds := action.Type.SatisfiesNeeds(); satisfyNeeds != nil {
		for _, need := range satisfyNeeds {
			ctx.SatisfyNeed(need, action.Type, action.TickProgress)
		}
	}

	if satisfyDesires := action.Type.SatisfiesDesires(); satisfyDesires != nil {
		for _, desire := range satisfyDesires {
			ctx.SatisfyDesire(desire, action.Type, action.TickProgress)
		}
	}

	if produces := action.Type.ProducesResources(); produces != nil {
		for _, produce := range produces {
			ctx.ProduceResource(produce, action.TickProgress)
		}
	}

	if skills := action.Type.ImprovesSkills(); skills != nil {
		for _, skill := range skills {
			ctx.ImproveSkill(skill)
		}
	}

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

	action.TickProgress = action.TickProgress + ctx.TickElapsed
}

func (ctx *Context) ProduceResource(produce types.ProduceResource, tickElapsed float64) {
	if rand.Float64() > produce.ChancePerTick*ctx.TickElapsed {
		return
	}

	ctx.ActiveColonist.Bag.Add(produce.Resource, 1)
}

func (ctx *Context) ImproveSkill(improve types.ImproveSkill) {
	ctx.ActiveColonist.Skills.Increase(improve.Skill, improve.AmountPerTick*ctx.TickElapsed)
}

func (ctx *Context) SatisfyNeed(need types.SatisfyNeed, action types.Actionable, actionProgress float64) {
	value := ctx.GetEasedValue(need.Total, need.Ease, float64(action.HasDuration()), actionProgress)

	ctx.ActiveColonist.Needs.Decrease(need.NeedType, value)
}

func (ctx *Context) SatisfyDesire(desire types.SatisfyDesire, action types.Actionable, actionProgress float64) {
	value := ctx.GetEasedValue(desire.Total, desire.Ease, float64(action.HasDuration()), actionProgress)

	ctx.ActiveColonist.Desires.Increase(desire.DesireType, value)
}

func (ctx *Context) AgitateNeed(need types.AgitateNeed, action types.Actionable, actionProgress float64) {
	value := ctx.GetEasedValue(need.Total, need.Ease, float64(action.HasDuration()), actionProgress)

	ctx.ActiveColonist.Needs.Increase(need.NeedType, value)
}

func (ctx *Context) RemoveResources(resources []types.ConsumeResource) {
	for _, resourceQuantity := range resources {
		ctx.ActiveColonist.Bag.Remove(resourceQuantity.Resource, resourceQuantity.Amount)
	}
}

func CheckBagSpace(b *colonist.Bag, quantity uint) bool {
	return b.GetAvailableSpace() >= quantity
}

func (ctx *Context) GetEasedValue(total float64, ease func(float64) float64, duration float64, tickProgress float64) float64 {
	if ease == nil {
		return (total / duration) * ctx.TickElapsed
	}

	previousTick := ease(tickProgress / duration)
	currentTick := ease((tickProgress + ctx.TickElapsed) / duration)

	return total * (currentTick - previousTick)
}
