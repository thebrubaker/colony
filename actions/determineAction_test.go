package actions

import (
	"testing"

	"github.com/thebrubaker/colony/actions/types"
	"github.com/thebrubaker/colony/colonist"
	"github.com/thebrubaker/colony/region"
)

func TestChooseWeightedAction(t *testing.T) {
	tickElapsed := 0.0

	region := &region.Region{}

	c := colonist.NewColonist("Test")

	c.Needs[colonist.Thirst] = 90
	c.Needs[colonist.Hunger] = 90
	c.Needs[colonist.Stress] = 90
	c.Needs[colonist.Exhaustion] = 90

	ctx := &Context{
		Region:         region,
		Colonists:      []*colonist.Colonist{c},
		ActiveColonist: c,
		TickElapsed:    tickElapsed,
	}

	choices := ctx.GetWeightedChoices(nil, []types.Actionable{
		types.DrinkGroundWater,
		types.GatherWildBerries,
		types.BasicRelax,
		types.GatherWood,
	})

	// t.Error(spew.Sdump(choices))

	if choices[0].Weight < choices[1].Weight {
		t.Error("failed to give drinking water more weight than eating food")
	}

	if choices[1].Weight < choices[2].Weight {
		t.Error("failed to give eating food more weight than wandering")
	}

	if choices[2].Weight < choices[3].Weight {
		t.Error("failed to give wandering more weight than working")
	}
}
