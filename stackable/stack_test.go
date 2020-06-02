package stackable

import (
	"testing"
)

func TestPut(t *testing.T) {
	elem := struct {
		Type string
	}{
		"some_elem",
	}

	var slice []interface{}

	slice, err := Put(slice, elem)

	if err != nil {
		t.Error("error returned from Put method")
	}

	if len(slice) != 1 {
		t.Error("slice length is not equal to one")
	}

	if slice[0] != elem {
		t.Error("elem failed to be placed into slice")
	}
}

func TestPutStack(t *testing.T) {
	elem := struct {
		Type string
	}{
		"some_elem",
	}

	stack := &Stack{
		elem,
		10,
	}

	var slice []interface{}

	slice, err := Put(slice, stack)

	if err != nil {
		t.Error("error returned from Put method")
	}

	if len(slice) != 1 {
		t.Error("slice length is not equal to one")
	}

	stackable, ok := slice[0].(Stackable)

	if !ok {
		t.Error("stack failed to be placed into slice as stackable")
	}

	if stackable.GetElement() != stack.GetElement() {
		t.Error("stack failed to be placed into slice")
	}

	if stackable.GetCount() != stack.GetCount() {
		t.Error("stack count does not match stack within slice")
	}
}

func TestPutCombineStack(t *testing.T) {
	elem := struct {
		Type string
	}{
		"some_elem",
	}

	stackA := &Stack{
		elem,
		10,
	}

	stackB := &Stack{
		elem,
		10,
	}

	var slice []interface{}

	slice, err := Put(slice, stackA)

	if err != nil {
		t.Error("error returned from Put method")
	}

	slice, err = Put(slice, stackB)

	if len(slice) != 1 {
		t.Error("slice length is not equal to one")
	}

	stackable, ok := slice[0].(Stackable)

	if !ok {
		t.Error("stack failed to be placed into slice as stackable")
	}

	if stackable.GetElement() != stackA.GetElement() {
		t.Error("stack failed to be placed into slice")
	}

	if stackable.GetCount() != stackA.GetCount()+stackB.GetCount() {
		t.Error("stack count does not match stack within slice")
	}

	slice, err = Put(slice, elem)

	if err != nil {
		t.Error("error returned from Put method")
	}

	if len(slice) != 2 {
		t.Error("slice length is not equal to two")
		t.Errorf("%v", slice)
	}
}
