package actions

import (
	"github.com/thebrubaker/colony/actions/types"
	"github.com/thebrubaker/colony/colonist"
	"github.com/thebrubaker/colony/region"
)

type Context struct {
	Region         *region.Region
	Colonists      []*colonist.Colonist
	ActiveColonist *colonist.Colonist
	TickElapsed    float64
}
type Action struct {
	Type         types.Actionable
	TickProgress float64
}

type Actions map[string]*Action

func (a Actions) Update(ctx *Context) {
	for _, colonist := range ctx.Colonists {
		if activeAction, ok := a[colonist.Key]; ok {
			ctx.ActiveColonist = colonist
			action := ctx.DetermineAction(activeAction)
			ctx.UpdateAction(action)
			a[colonist.Key] = action
		}
	}
}

func CreateActions(colonists []*colonist.Colonist) Actions {
	actions := make(Actions)

	for _, colonist := range colonists {
		actions[colonist.Key] = &Action{
			Type: types.CryoSleepWakeup,
		}
		colonist.Status = types.CryoSleepWakeup.Status()[0]
	}

	return actions
}
