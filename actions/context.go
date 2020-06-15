package actions

import (
	"github.com/thebrubaker/colony/actions/types"
	"github.com/thebrubaker/colony/colonist"
	"github.com/thebrubaker/colony/need"
	"github.com/thebrubaker/colony/region"
)

type Context struct {
	Region      *region.Region
	Colonist    *colonist.Colonist
	TickElapsed float64
}

func (c *Context) DetermineAction(currentAction *Action) *Action {
	if currentAction != nil && currentAction.TickProgress < float64(currentAction.Type.Duration()) {
		return c.ContinueAction(currentAction)
	}

	return c.NextAction(currentAction)
}

func (c *Context) ContinueAction(a *Action) *Action {
	OnTick(c, a)

	a.TickProgress = a.TickProgress + c.TickElapsed

	return a
}

func (c *Context) NextAction(a *Action) *Action {
	c.OnEnd(a)

	nextAction := c.SelectNextAction(a.Type)

	c.OnStart(nextAction)

	return nextAction
}

func (c *Context) SelectNextAction(currentType types.Actionable) *Action {
	weights := GetActionWeights(c, currentType, types.InitTypes())

	return &Action{
		colonist: c.Colonist,
		Type:     weights[0].Type,
	}
}

func (c *Context) OnEnd(a *Action) {
	if i, ok := a.Type.(types.OnEnd); ok {
		i.OnEnd()
	}
}

func (c *Context) OnStart(a *Action) {
	if i, ok := a.Type.(types.OnStart); ok {
		i.OnStart()
	}

	c.Colonist.SetStatus(a.Type.Status())
}

func GetEasedValue(needType need.NeedType, total float64, ease func(float64) float64, duration float64, tickProgress float64, tickElapsed float64) (need.NeedType, float64) {
	if ease == nil {
		return needType, (total / duration) * tickElapsed
	}

	previousTick := ease(tickProgress / duration)
	currentTick := ease((tickProgress + tickElapsed) / duration)

	return needType, total * (currentTick - previousTick)
}
