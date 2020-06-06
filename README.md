# Colony (Working Title)

Colony is a colony / survival / simulation game written in Go. The game is currently experimental and a work in progress.

The premise is your colonist has woken up from cryosleep alone on a brand new planet. Unfortunately things haven't gone as planned, and you wake up alone, with most of your resources already looted or destroyed by others who woke up before you.

Slowly build your colony from the ground up despite this setback. Create a settlement, hunt for food, recruit new colonists, and discover advanced technology hidden away. But be careful. There seems to be some strange, dangerous creatures roaming about. Maybe something much worse.

## Planned technical features include

- gRPC API for streaming game and sending in commands
- Fully documented API for automating the game while you are away
- Multiplayer support to play with or against others
- OAuth support enabling custom API-driven clients such as voice commands and push notifications

## Planned game features include

- Utility-based decision trees. Assign colonists to leadership roles to improve coordinated effort.
- Construct farms, mines, workshops, fortificatons, trade depots and other buildings to sustain your colonists needs and desires
- Train your colonists and prepare them for battle from raiders, bugs and worse things crawling in the dark recesses of the planet
- Raid hidden bunkers for advanced technology

# How to start the game server

To do some basic streaming of the game server, you can run the following command.

Warning: you should run this in a very tall console window.

```script
go run main.go start --render true
```

# How to stream the game server

```script
Coming Soon
go run main.go stream
```

# Contributing

Guide coming soon.

## How To Add Actions

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
