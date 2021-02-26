package colonist

import (
	"testing"

	"github.com/thebrubaker/colony/resources"
)

func TestColonistBag(t *testing.T) {
	bag := Bag{
		Size: 10,
	}
	var err error

	err = bag.Add(resources.Wood, 10)

	if err != nil {
		t.Error(err)
	}

	if !bag.Has(resources.Wood, 10) {
		t.Error("wood was not placed in the colonist's bag")
	}

	err = bag.Remove(resources.Wood, 5)

	if err != nil {
		t.Error(err)
	}

	if !bag.Has(resources.Wood, 5) {
		t.Error("wood was not removed from the colonist's bag")
	}
}
