package types

import (
	"github.com/fogleman/ease"
	"github.com/thebrubaker/colony/need"
)

type StartingAction struct {
}

func (a *StartingAction) Status() string {
	return "waking up from cryosleep"
}

func (a *StartingAction) EnergyCost() EnergyCost {
	return Easiest
}

func (a *StartingAction) Duration() TickDuration {
	return Fast
}

func (a *StartingAction) Priority() Priority {
	return FightOrFlight
}

func (a *StartingAction) GetUtility() float64 {
	return 1
}

func (a *StartingAction) Agitates() (need.NeedType, float64, func(float64) float64) {
	return need.Stress, 100, func(w float64) float64 { return ease.InOutQuad(w) }
}
