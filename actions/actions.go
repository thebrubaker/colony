package actions

import (
	"math/rand"

	"github.com/thebrubaker/colony/actions/types"
	"github.com/thebrubaker/colony/colonist"
	"github.com/thebrubaker/colony/region"
)

var ColonistActions = []types.Actionable{
	types.DrinkGroundWater,
	types.GatherWildBerries,
	types.GatherWood,
	types.BasicRelax,
	types.BasicRest,
	types.HaulItems,
}

type Context struct {
	Region         *region.Region
	Colonists      []*colonist.Colonist
	ActiveColonist *colonist.Colonist
	TickElapsed    float64
}
type Action struct {
	Type           types.Actionable
	TickProgress   float64
	TickDuration   types.TickDuration
	ActionComplete bool
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
			Type: types.Thinking,
		}
		statuses := types.Thinking.Status()
		randomIndex := rand.Intn(len(statuses))
		colonist.SetStatus(statuses[randomIndex])
	}

	return actions
}
