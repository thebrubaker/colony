package colonist

import (
	"encoding/json"
	"errors"

	"github.com/thebrubaker/colony/stackable"
)

type Bag struct {
	Size  uint
	Items []interface{}
}

type Stackable interface {
	IsStackable() bool
}

type Counter interface {
	GetCount() uint
}

// MarshalJSON will marshal needs into it's attributes
func (b *Bag) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Items)
}

// UnmarshalJSON fills in the attributes of needs
func (b *Bag) UnmarshalJSON(bytes []byte) error {
	var items []interface{}

	if err := json.Unmarshal(bytes, &items); err != nil {
		return err
	}

	b.Items = items

	return nil
}

func (b *Bag) Add(elem interface{}, count uint) error {
	var items []interface{}
	var err error

	i, ok := elem.(Stackable)

	if !ok {
		return errors.New("non-stackable elements cannot be added with a count")
	}

	if i.IsStackable() {
		items, err = stackable.Put(b.Items, &stackable.Stack{
			Item:  elem,
			Count: count,
		})
	} else {
		items, err = stackable.Put(b.Items, elem)
	}

	if err != nil {
		return err
	}

	b.Items = items

	return nil
}

func (b *Bag) Remove(elem interface{}, count uint) error {
	var items []interface{}
	var err error

	i, ok := elem.(Stackable)

	if (!ok || !i.IsStackable()) && count > 0 {
		return errors.New("non-stackable elements cannot be removed with a count")
	}

	if i, ok := elem.(Stackable); ok && i.IsStackable() {
		items, err = stackable.Take(b.Items, &stackable.Stack{
			Item:  elem,
			Count: count,
		})
	} else {
		items, err = stackable.Take(b.Items, elem)
	}

	if err != nil {
		return err
	}

	b.Items = items

	return nil
}

func (b *Bag) Has(elem interface{}, count uint) bool {
	if i, ok := elem.(Stackable); (!ok || !i.IsStackable()) && count > 0 {
		return false
	}

	if i, ok := elem.(Stackable); ok && i.IsStackable() {
		return stackable.Has(b.Items, &stackable.Stack{
			Item:  elem,
			Count: count,
		})
	}

	return stackable.Has(b.Items, elem)
}

func (b *Bag) GetAvailableSpace() uint {
	var itemCount uint = 0

	for _, i := range b.Items {
		itemCount += getCount(i)
	}

	return uint(b.Size - itemCount)
}

func (b *Bag) GetItemCount() uint {
	var itemCount uint = 0

	for _, i := range b.Items {
		itemCount += getCount(i)
	}

	return uint(itemCount)
}

func (b *Bag) IsFull() bool {
	return b.GetAvailableSpace() == 0
}

func (b *Bag) IsEmpty() bool {
	return b.GetItemCount() == 0
}

func getCount(item interface{}) uint {
	if i, ok := item.(Counter); ok {
		return i.GetCount()
	}

	return 1
}
