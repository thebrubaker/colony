package need

import (
	"testing"
)

func Test(t *testing.T) {
	needs := NewNeeds()

	if needs.Get(Hunger) != 0 {
		t.Error(Hunger, " does not equal ", 0)
	}

	if needs.Get(Exhaustion) != 0 {
		t.Error(Exhaustion, " does not equal ", 0)
	}

	if needs.Get(Stress) != 0 {
		t.Error(Stress, " does not equal ", 0)
	}

	if needs.Get(Thirst) != 0 {
		t.Error(Thirst, " does not equal ", 0)
	}

	needs.Increase(Hunger, 100)

	if needs.Get(Hunger) != 100 {
		t.Error(Hunger, " does not equal ", 100)
	}

	needs.Increase(Hunger, 20)

	if needs.Get(Hunger) != 100 {
		t.Error(Hunger, " does not equal max ", 100)
	}

	needs.Decrease(Hunger, 100)

	if needs.Get(Hunger) != 0 {
		t.Error(Hunger, " does not equal ", 0)
	}

	needs.Decrease(Hunger, 20)

	if needs.Get(Hunger) != 0 {
		t.Error(Hunger, " does not equal min ", 0)
	}

	needs.Set(Hunger, -20)

	if needs.Get(Hunger) != 0 {
		t.Error(Hunger, " does not equal min ", 0)
	}

	needs.Set(Hunger, 120)

	if needs.Get(Hunger) != 100 {
		t.Error(Hunger, " does not equal min ", 100)
	}
}
