package actions

import (
	"testing"

	"github.com/thebrubaker/colony/actions/types"
	"github.com/thebrubaker/colony/colonist"
	"github.com/thebrubaker/colony/need"
	"github.com/thebrubaker/colony/region"
)

func TestUtility(t *testing.T) {
	tickElapsed := 0.0

	region := &region.Region{}

	colonist := colonist.GenerateColonist("Test")

	colonist.Needs.Set(need.Hunger, 100)
	colonist.Needs.Set(need.Thirst, 60)
	colonist.Needs.Set(need.Stress, 0)

	context := &Context{
		Region:      region,
		Colonist:    colonist,
		TickElapsed: tickElapsed,
	}

	PickBerries := &types.PickBerries{}
	DrinkFromStream := &types.DrinkFromStream{}
	WatchClouds := &types.WatchClouds{}

	weights := GetActionWeights(context, nil, []types.Actionable{
		PickBerries,
		DrinkFromStream,
		WatchClouds,
	})

	if weights[0].Type != PickBerries {
		t.Error("failed to sort PickBerries as highest weight", weights[0].Type, PickBerries)
	}

	if weights[0].Weight != float64(types.Survival)+20 {
		t.Error("failed to sort hunger as highest weight", weights[0].Weight, float64(types.Survival)+20)
	}

	if weights[1].Type != DrinkFromStream {
		t.Error("failed to sort DrinkFromStream as highest weight", weights[0].Type, DrinkFromStream)
	}

	if weights[1].Weight != float64(types.Survival) {
		t.Error("failed to sort thirst as second highest weight", weights[1].Weight, float64(types.Survival))
	}

	if weights[2].Type != WatchClouds {
		t.Error("failed to sort WatchClouds as highest weight", weights[0].Type, WatchClouds)
	}

	if weights[2].Weight != 0 {
		t.Error("failed to sort stress as third highest weight", weights[2].Weight, 0)
	}
}
