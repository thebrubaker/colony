package colonist

import (
	"errors"

	"github.com/thebrubaker/colony/stackable"
)

type Bag struct {
	items []interface{}
}

type Stackable interface {
	IsStackable() bool
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
