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
		t.Error(err)
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
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("slice length is not equal to one")
	}

	stackable, ok := slice[0].(Stackable)

	if !ok {
		t.Error("stack failed to be placed into slice as stackable")
	}

	if stackable.GetItem() != stack.GetItem() {
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
		t.Error(err)
	}

	slice, err = Put(slice, stackB)

	if len(slice) != 1 {
		t.Error("slice length is not equal to one")
	}

	stackable, ok := slice[0].(Stackable)

	if !ok {
		t.Error("stack failed to be placed into slice as stackable")
	}

	if stackable.GetItem() != stackA.GetItem() {
		t.Error("stack failed to be placed into slice")
	}

	if stackable.GetCount() != stackA.GetCount()+stackB.GetCount() {
		t.Error("stack count does not match stack within slice")
	}

	slice, err = Put(slice, elem)

	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("slice length is not equal to two")
		t.Errorf("%v", slice)
	}
}

func TestHas(t *testing.T) {
	elem := struct {
		Type string
	}{
		"some_elem",
	}

	slice := []interface{}{elem}

	if ok := Has(slice, elem); !ok {
		t.Error("failed to detect element in slice")
	}
}

func TestHasStack(t *testing.T) {
	elem := struct {
		Type string
	}{
		"some_elem",
	}

	stack := &Stack{elem, 10}

	slice := []interface{}{stack}

	if ok := Has(slice, &Stack{elem, 5}); !ok {
		t.Error("failed to detect correct count of element in slice")
	}

	if ok := Has(slice, &Stack{elem, 10}); !ok {
		t.Error("failed to detect correct count of element in slice")
	}

	if ok := Has(slice, &Stack{elem, 20}); ok {
		t.Error("failed to detect correct count of element in slice")
	}
}

func TestTake(t *testing.T) {
	elem := struct {
		Type string
	}{
		"some_elem",
	}

	slice := []interface{}{elem}

	slice, err := Take(slice, elem)

	if err != nil {
		t.Error(err)
	}

	if len(slice) != 0 {
		t.Error("failed to take element from slice")
	}
}

func TestTakeStack(t *testing.T) {
	elem := struct {
		Type string
	}{
		"some_elem",
	}

	stack := &Stack{elem, 10}

	slice := []interface{}{stack}

	slice, err := Take(slice, &Stack{elem, 5})

	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("failed to leave stack in slice")
	}

	s, ok := slice[0].(Stackable)

	if !ok {
		t.Error("slice fails to contain stackable")
		t.Errorf("%v", slice)
	}

	if s.GetCount() != 5 {
		t.Error("failed to take correct count of element from slice")
		t.Errorf("%d", s.GetCount())
	}

	slice, err = Take(slice, &Stack{elem, 5})

	if err != nil {
		t.Error(err)
	}

	if len(slice) != 0 {
		t.Error("slice fails to be empty after taking remaining count of element")
		t.Errorf("%v", slice)
	}
}
