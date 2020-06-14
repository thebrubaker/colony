package actions

import (
	"sort"

	"github.com/thebrubaker/colony/actions/types"
)

type ActionWeight struct {
	Type   types.Actionable
	Weight float64
}

func GetActionWeights(c *Context, currentAction types.Actionable, availableActions []types.Actionable) []*ActionWeight {
	var weights []*ActionWeight

	for _, actionType := range availableActions {
		weights = append(weights, &ActionWeight{
			Type:   actionType,
			Weight: GetUtility(c, currentAction, actionType),
		})
	}

	sort.Slice(weights, func(i, j int) bool {
		return weights[i].Weight > weights[j].Weight
	})

	return weights
}
