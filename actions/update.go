package actions

import (
	"github.com/thebrubaker/colony/colonist"
	"github.com/thebrubaker/colony/region"
)

func Update(tickElapsed float64, region *region.Region, colonists []*colonist.Colonist, actions []*Action) {
	for _, colonist := range colonists {
		for i, a := range actions {
			if a.Colonist == colonist {
				context := &Context{
					Region:      region,
					Colonist:    colonist,
					TickElapsed: tickElapsed,
				}

				actions[i] = context.DetermineAction(a)
			}
		}
	}
}
