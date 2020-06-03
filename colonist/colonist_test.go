package colonist

import (
	"testing"

	"github.com/thebrubaker/colony/stackable"
)

func TestColonistBag(t *testing.T) {
	colonist := GenerateColonist("Joel")

	wood := struct {
		Name string
	}{
		"Wood",
	}

	colonist.AddToBag(&stackable.Stack{
		Element: wood,
		Count:   10,
	})

	if !colonist.Has(&stackable.Stack{
		Element: wood,
		Count:   10,
	}) {
		t.Error("wood was not placed in the colonist's bag")
	}

	colonist.RemoveFromBag(&stackable.Stack{
		Element: wood,
		Count:   5,
	})

	if !colonist.Has(&stackable.Stack{
		Element: wood,
		Count:   5,
	}) {
		t.Error("wood was not removed from the colonist's bag")
	}
}
