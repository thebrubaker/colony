package types

import (
	"github.com/fogleman/ease"
	"github.com/thebrubaker/colony/colonist"
	"github.com/thebrubaker/colony/resources"
)

var DrinkGroundWater = &SimpleAction{
	status: []string{
		"drinking water from a nearby stream",
		"drinking a small amount of collected rain water",
	},
	effort:      Easy,
	duration:    Slow,
	utilityNeed: colonist.Thirst,
	satisfiesNeeds: []SatisfyNeed{
		{colonist.Thirst, 65, ease.InOutCubic},
	},
	agitatesNeeds: []AgitateNeed{
		{colonist.Stress, 10, nil},
	},
}

var GatherWildBerries = &SimpleAction{
	status: []string{
		"gathering wild berries",
		"gathering berries from a nearby forest",
		"gathering ripe berries from wild bushes",
	},
	effort:      Easy,
	duration:    Slow,
	utilityNeed: colonist.Hunger,
	satisfiesNeeds: []SatisfyNeed{
		{colonist.Hunger, 35, ease.InOutCubic},
		{colonist.Stress, 10, ease.InOutCubic},
	},
	producesResources: []ProduceResource{
		{resources.Berries, 0.75},
	},
	improvesSkills: []ImproveSkill{
		{colonist.Gathering, 0.1},
	},
}

var GatherWood = &SimpleAction{
	status: []string{
		"gathering dried out sticks and logs from the ground",
		"gathering large branches from a nearby forest",
		"collecting large sticks and branches from nearby",
	},
	effort:        Demanding,
	duration:      Slow,
	utilityDesire: colonist.Fulfillment,
	satisfiesDesires: []SatisfyDesire{
		{colonist.Fulfillment, 35, ease.InOutCubic},
	},
	producesResources: []ProduceResource{
		{resources.Wood, 0.5},
	},
	improvesSkills: []ImproveSkill{
		{colonist.Woodcutting, 0.05},
	},
}

var BasicRelax = &SimpleAction{
	status: []string{
		"sitting on the ground lost in thought",
		"watching clouds as they pass by in the distance",
		"wandering around aimlessly",
	},
	effort:      Painless,
	duration:    Moderate,
	utilityNeed: colonist.Stress,
	satisfiesNeeds: []SatisfyNeed{
		{colonist.Stress, 30, ease.InOutCubic},
	},
}

var BasicRest = &SimpleAction{
	status: []string{
		"trying to sleep on a pile of leaves gathered from the ground",
		"huddled in a protected corner attempting to sleep",
		"trying to sleep using worn rags and torn clothing as a blanket",
	},
	effort:      Easy,
	duration:    Slow,
	utilityNeed: colonist.Exhaustion,
	satisfiesNeeds: []SatisfyNeed{
		{colonist.Exhaustion, 60, ease.InBounce},
	},
	agitatesNeeds: []AgitateNeed{
		{colonist.Stress, 40, ease.InBounce},
	},
}

var CryoSleepWakeup = &SimpleAction{
	status: []string{
		"waking up from cryosleep",
	},
	effort:   Painless,
	duration: Fast,
	agitatesNeeds: []AgitateNeed{
		{colonist.Stress, 30, ease.InOutQuint},
	},
}
