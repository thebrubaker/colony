package recipes

import "github.com/thebrubaker/colony/resources"

type RecipeInput struct {
	Elem  interface{}
	Count uint
}

type RecipeOutput struct {
	Elem  interface{}
	Count uint
}

type Recipe struct {
	In  []RecipeInput
	Out []RecipeOutput
}

var SimpleMeal = &Recipe{
	In: []RecipeInput{
		{
			Elem:  resources.Wood,
			Count: 10,
		},
	},
	Out: []RecipeOutput{
		{
			Elem:  resources.Wood,
			Count: 10,
		},
	},
}
