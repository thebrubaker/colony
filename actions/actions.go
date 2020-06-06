package actions

import (
	"github.com/thebrubaker/colony/actions/types"
	"github.com/thebrubaker/colony/colonist"
	"github.com/thebrubaker/colony/region"
)

type Action struct {
	Colonist     *colonist.Colonist
	Type         types.Actionable
	TickProgress float64
}

func InitActions(region *region.Region, colonists []*colonist.Colonist) []*Action {
	var actions []*Action

	for _, colonist := range colonists {
		context := &Context{
			Region:      region,
			Colonist:    colonist,
			TickElapsed: 0,
		}

		actions = append(actions, context.CreateStartingAction(colonist))
	}

	return actions
}

func (c *Context) CreateStartingAction(colonist *colonist.Colonist) *Action {
	action := &Action{
		Type:     &types.StartingAction{},
		Colonist: colonist,
	}

	c.OnStart(action)

	return action
}
