package actions

import (
	"log"
	"math/rand"

	"github.com/thebrubaker/colony/actions/types"
	"github.com/thebrubaker/colony/colonist"
	"github.com/thebrubaker/colony/stackable"
)

func (ctx *Context) StartAction(action types.Actionable) {
	if resources := action.ConsumesResources(); resources != nil {
		ctx.RemoveResources(resources)
	}
	statuses := action.Status()
	randomIndex := rand.Intn(len(statuses))
	ctx.ActiveColonist.SetStatus(statuses[randomIndex])
}

func (ctx *Context) EndAction(action types.Actionable) {

}

func (ctx *Context) UpdateAction(action *Action) {
	ctx.ActiveColonist.Needs.Increase(colonist.Exhaustion, float64(action.Type.TakesEffort())*ctx.TickElapsed)
	ctx.ActiveColonist.Needs.Increase(colonist.Hunger, 0.05*ctx.TickElapsed)
	ctx.ActiveColonist.Needs.Increase(colonist.Thirst, 0.05*ctx.TickElapsed)
	ctx.ActiveColonist.Needs.Increase(colonist.Stress, 0.05*ctx.TickElapsed)

	ctx.ActiveColonist.Desires.Decrease(colonist.Fulfillment, 0.05*ctx.TickElapsed)
	ctx.ActiveColonist.Desires.Decrease(colonist.Belonging, 0.05*ctx.TickElapsed)
	ctx.ActiveColonist.Desires.Decrease(colonist.Esteem, 0.05*ctx.TickElapsed)

	if agitateNeeds := action.Type.AgitatesNeeds(); agitateNeeds != nil {
		for _, need := range agitateNeeds {
			ctx.AgitateNeed(need, float64(action.TickDuration), action.TickProgress)
		}
	}

	if satisfyNeeds := action.Type.SatisfiesNeeds(); satisfyNeeds != nil {
		for _, need := range satisfyNeeds {
			ctx.SatisfyNeed(need, float64(action.TickDuration), action.TickProgress)
		}
	}

	if satisfyDesires := action.Type.SatisfiesDesires(); satisfyDesires != nil {
		for _, desire := range satisfyDesires {
			ctx.SatisfyDesire(desire, float64(action.TickDuration), action.TickProgress)
		}
	}

	if produces := action.Type.ProducesResources(); produces != nil {
		for _, produce := range produces {
			ctx.ProduceResource(produce)
		}
	}

	if skills := action.Type.ImprovesSkills(); skills != nil {
		for _, skill := range skills {
			ctx.ImproveSkill(skill)
		}
	}

	if action.Type == types.HaulItems {
		ctx.UnloadBagUntilEmpty(float64(action.TickDuration))
	}

	if ctx.shouldEndEarly(action) {
		action.ActionComplete = true
	}

	action.TickProgress = action.TickProgress + ctx.TickElapsed
}

func (ctx *Context) ProduceResource(produce types.ProduceResource) {
	if ctx.ActiveColonist.Bag.IsFull() {
		return
	}

	skillScore := ctx.ActiveColonist.Skills[produce.Skill]
	chancePerTick := produce.ChancePerTick + skillScore/100

	if rand.Float64() > chancePerTick*ctx.TickElapsed {
		return
	}

	ctx.ActiveColonist.Bag.Add(produce.Resource, 1)
}

func (ctx *Context) ImproveSkill(improve types.ImproveSkill) {
	ctx.ActiveColonist.Skills.Increase(improve.Skill, improve.AmountPerTick*ctx.TickElapsed)
}

func (ctx *Context) SatisfyNeed(need types.SatisfyNeed, duration float64, actionProgress float64) {
	value := ctx.GetEasedValue(need.Total, need.Ease, float64(duration), actionProgress)

	ctx.ActiveColonist.Needs.Decrease(need.NeedType, value)
}

func (ctx *Context) SatisfyDesire(desire types.SatisfyDesire, duration float64, actionProgress float64) {
	value := ctx.GetEasedValue(desire.Total, desire.Ease, float64(duration), actionProgress)

	ctx.ActiveColonist.Desires.Increase(desire.DesireType, value)
}

func (ctx *Context) AgitateNeed(need types.AgitateNeed, duration float64, actionProgress float64) {
	value := ctx.GetEasedValue(need.Total, need.Ease, float64(duration), actionProgress)

	ctx.ActiveColonist.Needs.Increase(need.NeedType, value)
}

func (ctx *Context) RemoveResources(resources []types.ConsumeResource) {
	for _, resourceQuantity := range resources {
		ctx.ActiveColonist.Bag.Remove(resourceQuantity.Resource, resourceQuantity.Amount)
	}
}

func (ctx *Context) UnloadBagUntilEmpty(tickDuration float64) {
	if ctx.ActiveColonist.Bag.IsEmpty() {
		return
	}

	// roll the dice per tick, return if no go
	if rand.Float64() > 2*ctx.TickElapsed {
		return
	}

	// take first item, remove it, add to stockpile
	item := ctx.ActiveColonist.Bag.Items[0]
	if i, ok := item.(stackable.Stackable); ok {
		item = i.GetItem()
	}

	if err := ctx.ActiveColonist.Bag.Remove(item, 1); err != nil {
		log.Fatal(err)
	}
	if err := ctx.Region.Stockpile.Add(item, 1); err != nil {
		log.Fatal(err)
	}
}

func (ctx *Context) GetEasedValue(total float64, ease func(float64) float64, duration float64, tickProgress float64) float64 {
	if ease == nil {
		return (total / duration) * ctx.TickElapsed
	}

	previousTick := ease(tickProgress / duration)
	currentTick := ease((tickProgress + ctx.TickElapsed) / duration)

	return total * (currentTick - previousTick)
}
