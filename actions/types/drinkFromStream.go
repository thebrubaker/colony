package types

import (
	"github.com/fogleman/ease"
	"github.com/thebrubaker/colony/need"
)

type DrinkFromStream struct {
}

func (a *DrinkFromStream) Status() string {
	return "drinking water from a local stream"
}

func (a *DrinkFromStream) EnergyCost() EnergyCost {
	return Easy
}

func (a *DrinkFromStream) Duration() TickDuration {
	return Slow
}

func (a *DrinkFromStream) Priority() Priority {
	return Survival
}

func (a *DrinkFromStream) Need() need.NeedType {
	return need.Thirst
}

func (a *DrinkFromStream) Satisfies() (need.NeedType, float64, func(float64) float64) {
	return need.Thirst, 70, func(w float64) float64 { return ease.InQuad(w) }
}
