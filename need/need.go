package need

import (
	"encoding/json"
	"math"
)

type NeedType string

// The constants for each need type.
const (
	Thirst     NeedType = "thirst"
	Stress     NeedType = "stress"
	Exhaustion NeedType = "exhaustion"
	Hunger     NeedType = "hunger"
)

type attributes struct {
	Thirst     float64
	Stress     float64
	Exhaustion float64
	Hunger     float64
}

// Needs are the basic needs of a colonist
type Needs struct {
	attributes *attributes
}

// NewNeeds returns a set of attributes starting at 0
func NewNeeds() *Needs {
	return &Needs{
		attributes: &attributes{
			Thirst:     60,
			Stress:     0,
			Exhaustion: 0,
			Hunger:     60,
		},
	}
}

// MarshalJSON will marshal needs into it's attributes
func (n *Needs) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.attributes)
}

// UnmarshalJSON fills in the attributes of needs
func (n *Needs) UnmarshalJSON(b []byte) error {
	a := &attributes{}

	if err := json.Unmarshal(b, a); err != nil {
		return err
	}

	n.attributes = a

	return nil
}

// Increase will increase the given need type by the value
// with a ceiling of 0
func (n *Needs) Increase(t NeedType, value float64) {
	switch t {
	case Hunger:
		n.attributes.Hunger = math.Min(100, n.attributes.Hunger+value)
	case Thirst:
		n.attributes.Thirst = math.Min(100, n.attributes.Thirst+value)
	case Exhaustion:
		n.attributes.Exhaustion = math.Min(100, n.attributes.Exhaustion+value)
	case Stress:
		n.attributes.Stress = math.Min(100, n.attributes.Stress+value)
	}
}

// Decrease will decrease the given need type by the value
// with a floor of 0
func (n *Needs) Decrease(t NeedType, value float64) {
	switch t {
	case Hunger:
		n.attributes.Hunger = math.Max(0, n.attributes.Hunger-value)
	case Thirst:
		n.attributes.Thirst = math.Max(0, n.attributes.Thirst-value)
	case Exhaustion:
		n.attributes.Exhaustion = math.Max(0, n.attributes.Exhaustion-value)
	case Stress:
		n.attributes.Stress = math.Max(0, n.attributes.Stress-value)
	}
}

// Set sets the need type `t` with the given value. Value will be
// capped at 100 with a floor of 0.
func (n *Needs) Set(t NeedType, value float64) {
	if value > 100 {
		value = 100
	}

	if value < 0 {
		value = 0
	}

	switch t {
	case Hunger:
		n.attributes.Hunger = value
	case Thirst:
		n.attributes.Thirst = value
	case Exhaustion:
		n.attributes.Exhaustion = value
	case Stress:
		n.attributes.Stress = value
	}
}

// Get returns the attribute value for the need type t
func (n *Needs) Get(t NeedType) float64 {
	switch t {
	case Hunger:
		return n.attributes.Hunger
	case Thirst:
		return n.attributes.Thirst
	case Exhaustion:
		return n.attributes.Exhaustion
	case Stress:
		return n.attributes.Stress
	}

	return 0
}
