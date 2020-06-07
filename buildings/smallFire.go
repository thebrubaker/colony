package buildings

import "github.com/thebrubaker/colony/recipes"

type SmallFire struct {
}

func (b *SmallFire) Allows() []*recipes.Recipe {
	return []*recipes.Recipe{
		recipes.SimpleMeal,
	}
}
