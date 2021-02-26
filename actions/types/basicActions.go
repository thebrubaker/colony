package types

import (
	"github.com/fogleman/ease"
	"github.com/thebrubaker/colony/colonist"
	"github.com/thebrubaker/colony/resources"
)

var DrinkGroundWater = &SimpleAction{
	status: []string{
		"drinking water from a nearby stream",
	},
	effort:      Easy,
	duration:    Slow,
	utilityNeed: colonist.Thirst,
	satisfiesNeeds: []SatisfyNeed{
		{colonist.Thirst, 65, ease.OutQuad},
	},
	agitatesNeeds: []AgitateNeed{
		{colonist.Stress, 10, nil},
	},
}

var GatherWildBerries = &SimpleAction{
	status: []string{
		"gathering wild berries",
	},
	effort:      Easy,
	duration:    Slow,
	utilityNeed: colonist.Hunger,
	satisfiesNeeds: []SatisfyNeed{
		{colonist.Hunger, 35, ease.OutQuad},
		{colonist.Stress, 10, ease.OutQuad},
	},
	producesResources: []ProduceResource{
		{resources.Berries, 0.7},
	},
	improvesSkills: []ImproveSkill{
		{colonist.Gathering, 0.05},
	},
}

var GatherWood = &SimpleAction{
	status: []string{
		"gathering wood nearby",
	},
	effort:        Demanding,
	duration:      Slow,
	utilityDesire: colonist.Fulfillment,
	satisfiesDesires: []SatisfyDesire{
		{colonist.Fulfillment, 35, ease.OutQuad},
	},
	producesResources: []ProduceResource{
		{resources.Wood, 0.1},
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
		"going on a walk",
		"exploring a nearby area",
	},
	effort:      Painless,
	duration:    Moderate,
	utilityNeed: colonist.Stress,
	satisfiesNeeds: []SatisfyNeed{
		{colonist.Stress, 30, ease.OutQuad},
	},
}

var BasicRest = &SimpleAction{
	status: []string{
		"sleeping on the ground",
	},
	effort:      Easy,
	duration:    Slow,
	utilityNeed: colonist.Exhaustion,
	satisfiesNeeds: []SatisfyNeed{
		{colonist.Exhaustion, 60, ease.OutQuad},
	},
	agitatesNeeds: []AgitateNeed{
		{colonist.Stress, 40, ease.OutQuad},
	},
}

var Thinking = &SimpleAction{
	status: []string{
		"",
	},
	effort:   Painless,
	duration: Fastest,
	satisfiesDesires: []SatisfyDesire{
		{colonist.Fulfillment, 5, ease.OutQuad},
	},
}
