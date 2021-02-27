package storage

import (
	"encoding/json"
	"errors"

	"github.com/thebrubaker/colony/stackable"
)

type Storage struct {
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
func (s *Storage) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Items)
}

// UnmarshalJSON fills in the attributes of needs
func (s *Storage) UnmarshalJSON(bytes []byte) error {
	var items []interface{}

	if err := json.Unmarshal(bytes, &items); err != nil {
		return err
	}

	s.Items = items

	return nil
}

func (s *Storage) Add(elem interface{}, count uint) error {
	var items []interface{}
	var err error

	i, ok := elem.(Stackable)

	if !ok {
		return errors.New("non-stackable elements cannot be added with a count")
	}

	if i.IsStackable() {
		items, err = stackable.Put(s.Items, &stackable.Stack{
			Item:  elem,
			Count: count,
		})
	} else {
		items, err = stackable.Put(s.Items, elem)
	}

	if err != nil {
		return err
	}

	s.Items = items

	return nil
}

func (s *Storage) Remove(elem interface{}, count uint) error {
	var items []interface{}
	var err error

	i, ok := elem.(Stackable)

	if (!ok || !i.IsStackable()) && count > 0 {
		return errors.New("non-stackable elements cannot be removed with a count")
	}

	if i, ok := elem.(Stackable); ok && i.IsStackable() {
		items, err = stackable.Take(s.Items, &stackable.Stack{
			Item:  elem,
			Count: count,
		})
	} else {
		items, err = stackable.Take(s.Items, elem)
	}

	if err != nil {
		return err
	}

	s.Items = items

	return nil
}

func (s *Storage) Has(elem interface{}, count uint) bool {
	if i, ok := elem.(Stackable); (!ok || !i.IsStackable()) && count > 0 {
		return false
	}

	if i, ok := elem.(Stackable); ok && i.IsStackable() {
		return stackable.Has(s.Items, &stackable.Stack{
			Item:  elem,
			Count: count,
		})
	}

	return stackable.Has(s.Items, elem)
}

func (s *Storage) GetAvailableSpace() uint {
	var itemCount uint = 0

	for _, i := range s.Items {
		itemCount += getCount(i)
	}

	return uint(s.Size - itemCount)
}

func (s *Storage) GetItemCount() uint {
	var itemCount uint = 0

	for _, i := range s.Items {
		itemCount += getCount(i)
	}

	return uint(itemCount)
}

func (s *Storage) IsFull() bool {
	return s.GetAvailableSpace() == 0
}

func (s *Storage) IsEmpty() bool {
	return s.GetItemCount() == 0
}

func getCount(item interface{}) uint {
	if i, ok := item.(Counter); ok {
		return i.GetCount()
	}

	return 1
}
