package actions

import "github.com/thebrubaker/colony/actions/types"

func GetUtility(c *Context, currentType types.Actionable, a types.Actionable) float64 {
	if i, ok := a.(types.GetUtility); ok {
		return i.GetUtility()
	}

	if i, ok := a.(types.Weighted); ok {
		return float64(a.Priority()) + i.Weight()*25
	}

	if i, ok := a.(types.SimpleNeed); ok {
		need := c.Colonist.Needs.Get(i.Need())
		if need < float64(i.Priority()) {
			return 0
		}
		return float64(a.Priority()) + getSimpleWeight(i, c.Colonist.Needs.Get(i.Need()))*25
	}

	return float64(a.Priority()) + 10
}

func getSimpleWeight(i types.SimpleNeed, need float64) float64 {
	priority := float64(i.Priority())

	w := (need - priority) / (100 - priority)

	if e, ok := i.(types.EasedWeight); ok {
		return e.EaseWeight(w)
	}

	return w
}
