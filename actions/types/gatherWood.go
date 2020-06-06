package types

import (
	"encoding/json"

	"github.com/thebrubaker/colony/resources"
)

type GatherWood struct {
	Gathered float64
}

// MarshalJSON will marshal needs into it's attributes
func (a *GatherWood) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"status":      a.Status(),
		"energy_cost": a.EnergyCost(),
		"duration":    a.Duration(),
		"priority":    a.Priority(),
	})
}

// UnmarshalJSON fills in the attributes of needs
func (a *GatherWood) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, a); err != nil {
		return err
	}

	return nil
}

func (a *GatherWood) Status() string {
	return "gathering wood from a nearby forest"
}

func (a *GatherWood) EnergyCost() EnergyCost {
	return Hard
}

func (a *GatherWood) Duration() TickDuration {
	return Slow
}

func (a *GatherWood) Priority() Priority {
	return Job
}

func (a *GatherWood) Gather() (StorageType, interface{}, float64) {
	return ColonistBag, resources.Wood, 0.5
}
