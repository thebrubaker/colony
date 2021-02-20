package types

import (
	"github.com/fogleman/ease"
	"github.com/thebrubaker/colony/colonist"
	"github.com/thebrubaker/colony/need"
)

type PickBerries struct {
}

func (a *PickBerries) Status() string {
	return "scavenging wild berries from a nearby hillside"
}

func (a *PickBerries) EnergyCost() EnergyCost {
	return Easy
}

func (a *PickBerries) Duration() TickDuration {
	return Slow
}

func (a *PickBerries) Priority() Priority {
	return Survival
}

func (a *PickBerries) Need() need.NeedType {
	return need.Hunger
}

func (a *PickBerries) Satisfies() (need.NeedType, float64, func(float64) float64) {
	return need.Hunger, 70, func(w float64) float64 { return ease.InQuad(w) }
}

func (a *PickBerries) SkillUp() (colonist.SkillType, float64, func(float64) float64) {
	return colonist.Gathering, 0.5, func(w float64) float64 { return w }
}
