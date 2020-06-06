# How to start the game server

To do some basic streaming of the game server, you can run the following command.

Warning: you should run this in a very tall console window.

```script
go run main.go start --render true
```

# How to stream the game server

```script
Coming Soon
```

# How To Add Actions

To create a new action, create a new file in the `game/actions/types` directory and then append it to the InitTypes() method at the top of `game/actions/types/types.go`. The following is an example action for gathering wood.

```go
package types

import (
	"encoding/json"

	"github.com/thebrubaker/colony/resources"
)

type GatherWood struct {
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
```

Open `game/actions/types/types.go` to view the different constants available to you.
