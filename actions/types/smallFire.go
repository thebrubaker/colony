package types

import (
	"encoding/json"
)

type BuildSmallFire struct {
}

// MarshalJSON will marshal needs into it's attributes
func (a *BuildSmallFire) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"status":      a.Status(),
		"energy_cost": a.EnergyCost(),
		"duration":    a.Duration(),
		"priority":    a.Priority(),
	})
}

// UnmarshalJSON fills in the attributes of needs
func (a *BuildSmallFire) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, a); err != nil {
		return err
	}

	return nil
}

func (a *BuildSmallFire) Status() string {
	return "building a small fire"
}

func (a *BuildSmallFire) EnergyCost() EnergyCost {
	return Easy
}

func (a *BuildSmallFire) Duration() TickDuration {
	return Moderate
}

func (a *BuildSmallFire) Priority() Priority {
	return Job
}
