package buildings

import (
	"github.com/thebrubaker/colony/resources"
	"github.com/thebrubaker/colony/stackable"
)

type CampFire struct {
}

func (b *CampFire) Name() string {
	return "Campfire"
}

func (b *CampFire) Description() string {
	return "A small campfire for cooking small meals and staying warm."
}

func (b *CampFire) BuildCost() []stackable.Stack {
	return []stackable.Stack{
		{
			Item:  resources.Wood,
			Count: 10,
		},
	}
}

func (b *CampFire) BuildDuration() float64 {
	return 10
}
