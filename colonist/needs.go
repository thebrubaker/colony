package colonist

import (
	"math"
	"reflect"
)

type Needs struct {
	Thirst     float64 `json:"thirst"`
	Stress     float64 `json:"stress"`
	Exhaustion float64 `json:"exhaustion"`
	Hunger     float64 `json:"hunger"`
}

func (n *Needs) Decrease(key string, value float64) {
	need := reflect.ValueOf(n).Elem().FieldByName(key)

	if !need.IsValid() {
		return
	}

	if !need.CanSet() {
		return
	}

	if need.Kind() != reflect.Float64 {
		return
	}

	newValue := math.Max(0, need.Float()-value)

	need.SetFloat(newValue)
}

func (n *Needs) Increase(key string, value float64) {
	need := reflect.ValueOf(n).Elem().FieldByName(key)

	if !need.IsValid() {
		return
	}

	if !need.CanSet() {
		return
	}

	if need.Kind() != reflect.Float64 {
		return
	}

	newValue := math.Min(100, need.Float()+value)

	need.SetFloat(newValue)
}
