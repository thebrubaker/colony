package types

import (
	"github.com/thebrubaker/colony/resources"
)

type GatherWood struct {
}

func (a *GatherWood) Status() string {
	return "gathering wood from a nearby forest"
}

func (a *GatherWood) EnergyCost() EnergyCost {
	return Hard
}

func (a *GatherWood) Duration() TickDuration {
	return Slow
}

func (a *GatherWood) Priority() Priority {
	return Job
}

func (a *GatherWood) AddToBag() (StorageType, interface{}, uint, float64) {
	return ColonistBag, resources.Wood, 1, 0.5
}
