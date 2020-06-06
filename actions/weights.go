package actions

import (
	"sort"

	"github.com/thebrubaker/colony/actions/types"
)

type ActionWeight struct {
	Type   types.Actionable
	Weight float64
}

func GetActionWeights(c *Context, currentType types.Actionable, actionTypes []types.Actionable) []*ActionWeight {
	var weights []*ActionWeight

	for _, actionType := range actionTypes {
		weights = append(weights, &ActionWeight{
			Type:   actionType,
			Weight: GetUtility(c, currentType, actionType),
		})
	}

	sort.Slice(weights, func(i, j int) bool {
		return weights[i].Weight > weights[j].Weight
	})

	return weights
}
