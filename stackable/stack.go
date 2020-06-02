package stackable

import "errors"

// Stackable is an interface for an element that can be placed into
// a slice but should be incremented as a stack instead of appended
// to the slice.
type Stackable interface {
	GetElement() interface{}
	GetCount() uint
	SetCount(uint)
}

// Stack is a struct that represents the count of a given element.
type Stack struct {
	Element interface{}
	Count   uint
}

// GetElement returns the Element of the Stack
func (s *Stack) GetElement() interface{} {
	return s.Element
}

// GetCount returns the count of the Element in the Stack
func (s *Stack) GetCount() uint {
	return s.Count
}

// SetCount sets the count of the Element in the Stack
func (s *Stack) SetCount(count uint) {
	s.Count = count
}

// Put is a utility method for putting an element which may or may not implement
// the Stackable interface into a slice of other elements or stackable interfaces.
// If the element is stackable, the stack will be incremented by the argument's stack
// count. If the element is not stackable, it will simply be appended to the slice.
func Put(slice []interface{}, elem interface{}) ([]interface{}, error) {
	if stack, ok := elem.(Stackable); ok {
		return putStack(slice, stack)
	}

	return putElem(slice, elem)
}

// Has is a utility method for determining if an element, or the count
// of a Stackable element exists within the given slice. It is useful to run a check
// on Has before executing Take to remove an element or Stackable from the slice.
func Has(slice []interface{}, elem interface{}) bool {
	if stack, ok := elem.(Stackable); ok {
		return hasStack(slice, stack)
	}

	return hasElem(slice, elem)
}

// Take is a utility method for removing an element for a Stackable from a
// given slice. If the element is Stackable, and the matching Stackable
// within the slice has the same count, the Stackable will be removed from the
// slice. If the existing Stackable has a higher count, then that Stackable count
// will be deducted by the given element Stackable. Otherwise an error is returned.
func Take(slice []interface{}, elem interface{}) ([]interface{}, error) {
	if stack, ok := elem.(Stackable); ok {
		return takeStack(slice, stack)
	}

	return takeElem(slice, elem)
}

func takeStack(slice []interface{}, stack Stackable) ([]interface{}, error) {
	for i, elem := range slice {
		if s, ok := elem.(Stackable); ok && s.GetElement() == stack.GetElement() {
			if s.GetCount() == stack.GetCount() {
				return removeIndex(slice, i), nil
			}

			if stack.GetCount() < s.GetCount() {
				slice[i] = &Stack{
					s.GetElement(),
					s.GetCount() - stack.GetCount(),
				}

				return slice, nil
			}

			return nil, errors.New("stack count is too large to take from slice")
		}
	}

	return nil, errors.New("stack not found within slice")
}

func takeElem(slice []interface{}, elem interface{}) ([]interface{}, error) {
	for i, e := range slice {
		if e == elem {
			return removeIndex(slice, i), nil
		}
	}

	return nil, errors.New("element does not exist within slice")
}

func removeIndex(slice []interface{}, i int) []interface{} {
	slice[i] = slice[len(slice)-1] // Copy last element to index i.
	slice[len(slice)-1] = ""       // Erase last element (write zero value).
	return slice[:len(slice)-1]    // Truncate slice.
}

func hasStack(slice []interface{}, stack Stackable) bool {
	for _, elem := range slice {
		if s, ok := elem.(Stackable); ok && s.GetElement() == stack.GetElement() {
			return s.GetCount() >= stack.GetCount()
		}
	}

	return false
}

func hasElem(slice []interface{}, elem interface{}) bool {
	for _, e := range slice {
		if e == elem {
			return true
		}
	}

	return false
}

func putStack(slice []interface{}, stack Stackable) ([]interface{}, error) {
	for i, elem := range slice {
		if s, ok := elem.(Stackable); ok && s.GetElement() == stack.GetElement() {
			combined, err := combineStacks(s, stack)

			slice[i] = combined

			return slice, err
		}
	}

	slice = append(slice, &Stack{
		stack.GetElement(),
		stack.GetCount(),
	})

	return slice, nil
}

func putElem(slice []interface{}, elem interface{}) ([]interface{}, error) {
	return append(slice, elem), nil
}

// CombineStacks returns a new Stack that combines the counts of two Stacks
// as long as the elements in the stacks match.
func CombineStacks(a Stackable, b Stackable) (*Stack, error) {
	if a.GetElement() != b.GetElement() {
		return nil, errors.New("cannot combine stacks where items do not match")
	}

	return &Stack{
		a.GetElement(),
		a.GetCount() + b.GetCount(),
	}, nil
}
