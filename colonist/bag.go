package colonist

import (
	"encoding/json"
	"errors"

	"github.com/thebrubaker/colony/stackable"
)

type Bag struct {
	items []interface{}
}

type Stackable interface {
	IsStackable() bool
}

// MarshalJSON will marshal needs into it's attributes
func (b *Bag) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.items)
}

// UnmarshalJSON fills in the attributes of needs
func (b *Bag) UnmarshalJSON(bytes []byte) error {
	var items []interface{}

	if err := json.Unmarshal(bytes, items); err != nil {
		return err
	}

	b.items = items

	return nil
}

func (b *Bag) Add(elem interface{}, count uint) error {
	var items []interface{}
	var err error

	if i, ok := elem.(Stackable); !ok && i.IsStackable() && count > 0 {
		return errors.New("non-stackable elements cannot be added with a count")
	}

	if i, ok := elem.(Stackable); ok && i.IsStackable() {
		items, err = stackable.Put(b.items, &stackable.Stack{elem, count})
	} else {
		items, err = stackable.Put(b.items, elem)
	}

	if err != nil {
		return err
	}

	b.items = items

	return nil
}

func (b *Bag) Remove(elem interface{}, count uint) error {
	var items []interface{}
	var err error

	if i, ok := elem.(Stackable); !ok && i.IsStackable() && count > 0 {
		return errors.New("non-stackable elements cannot be removed with a count")
	}

	if i, ok := elem.(Stackable); ok && i.IsStackable() {
		items, err = stackable.Take(b.items, &stackable.Stack{elem, count})
	} else {
		items, err = stackable.Take(b.items, elem)
	}

	if err != nil {
		return err
	}

	b.items = items

	return nil
}

func (b *Bag) Has(elem interface{}, count uint) bool {
	if i, ok := elem.(Stackable); (!ok || !i.IsStackable()) && count > 0 {
		return false
	}

	if i, ok := elem.(Stackable); ok && i.IsStackable() {
		return stackable.Has(b.items, &stackable.Stack{elem, count})
	}

	return stackable.Has(b.items, elem)
}
