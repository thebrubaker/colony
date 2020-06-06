package types

import (
	"github.com/fogleman/ease"
	"github.com/thebrubaker/colony/need"
)

type WatchClouds struct {
}

func (a *WatchClouds) Status() string {
	return "watching the clouds go by"
}

func (a *WatchClouds) EnergyCost() EnergyCost {
	return Easy
}

func (a *WatchClouds) Duration() TickDuration {
	return Slow
}

func (a *WatchClouds) Priority() Priority {
	return Recreation
}

func (a *WatchClouds) Need() need.NeedType {
	return need.Stress
}

func (a *WatchClouds) Satisfies() (need.NeedType, float64, func(float64) float64) {
	return need.Stress, 70, func(w float64) float64 { return ease.InQuad(w) }
}
