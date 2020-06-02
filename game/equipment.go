// package game

// import (
// 	"encoding/json"
// 	"errors"
// 	"math"
// )

// type Slottable interface {
// 	GetSlot() Item
// 	GetCount() uint
// }

// type ItemStack struct {
// 	Item  Item `json:"item"`
// 	Count uint `json:"count"`
// }

// func (s *ItemStack) MarshalJSON() ([]byte, error) {
// 	m := struct {
// 		Type  string `json:"stack`
// 		Item  Item   `json:"item`
// 		Count uint   `json:"count"`
// 	}{
// 		"basic",
// 		s.Item,
// 		s.Count,
// 	}

// 	return json.Marshal(m)
// }

// func (s *ItemStack) GetSlot() Item {
// 	return s.Item
// }

// func (s *ItemStack) GetCount() uint {
// 	return s.Count
// }

// type Carryable interface {
// 	CanCarry() bool
// }

// type Equipment struct {
// }

// func (e *Equipment) HoldItem(i Carryable) error {
// 	if e.Hands == i {
// 		return errors.New("Item is already in hands.")
// 	}

// 	e.Hands = i

// 	return nil
// }

// func (e *Equipment) PutInBag(i Item) error {
// 	for _, stack := range e.Bag {
// 		if stack.Item == i {
// 			stack.Count++
// 			return nil
// 		}
// 	}

// 	e.Bag = append(e.Bag, &Slottable{i, 1})

// 	return nil
// }

// type Item interface {
// 	GetName() string
// 	GetCount() uint
// }

// type Consumable interface {
// 	GetNourishment() uint
// 	GetPleasure() uint
// }

// type Food struct {
// 	Name    string
// 	Quality float64
// }

// func (f *Food) GetName() string {
// 	return f.Name
// }

// func (f *Food) GetQuality() float64 {
// 	return f.Quality
// }

// func (f *Food) GetNourishment() float64 {
// 	return math.Max(100, f.Quality*30)
// }

// func (f *Food) GetPleasure() float64 {
// 	return math.Max(100, f.Quality*10)
// }

// var SimpleMeal *Food = &Food{
// 	Name:    "Simple Meal",
// 	Quality: 2,
// }
