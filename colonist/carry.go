package colonist

import (
	"errors"

	"github.com/thebrubaker/colony/stackable"
)

// Carryable interface
type Carryable interface {
	CanCarry() bool
}

// Carry will place the given element in the colonist's hands. If
// the element is stackable.Stackable, the element will be stacked.
func (c *Colonist) Carry(elem interface{}) error {
	if c.Hands != nil {
		return errors.New("colonist's hands are not empty")
	}

	if stack, ok := elem.(stackable.Stackable); ok {
		return c.carryStack(stack)
	}

	c.Hands = elem

	return nil
}

// Drop will remove the current element in the colonist's hands and return it.
func (c *Colonist) Drop() (interface{}, error) {
	if c.Hands == nil {
		return nil, errors.New("colonist has nothing to drop")
	}

	elem := c.Hands

	c.Hands = nil

	return elem, nil
}

func (c *Colonist) carryStack(stack stackable.Stackable) error {
	if i, ok := c.Hands.(stackable.Stackable); ok {
		combined, err := stackable.CombineStacks(i, stack)

		if err != nil {
			return err
		}

		c.Hands = combined

		return nil
	}

	return errors.New("element in colonist's hands does not match the given stack")
}
