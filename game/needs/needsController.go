package needs

import (
	"math"

	"github.com/thebrubaker/colony/colonist"
)

type needType string

// Constants for each need type
const (
	Rest        needType = "Rest"
	Food        needType = "Food"
	Water       needType = "Water"
	Security    needType = "Security"
	Belonging   needType = "Belonging"
	Fulfillment needType = "Fulfillment"
	Power       needType = "Power"
	Joy         needType = "Joy"
	Wealth      needType = "Wealth"
	Status      needType = "Status"
)

type NeedsController struct {
	state   map[colonist.ColonistKey]map[needType]*need
	actionc chan func()
	quitc   chan struct{}
}

func NewNeedsController(k colonist.ColonistKey) *NeedsController {
	nc := &NeedsController{}
	go nc.start()
	return nc
}

func (nc *NeedsController) start() {
	for {
		select {
		case f := <-nc.actionc:
			f()
		case <-nc.quitc:
			return
		}
	}
}

func (nc *NeedsController) update(tickElapsed float64) {
	c := make(chan struct{})
	nc.actionc <- func() {
		var size int
		for _, needs := range nc.state {
			for range needs {
				size++
			}
		}
		d := make(chan bool, size)
		for _, needs := range nc.state {
			for _, n := range needs {
				go func(n *need) {
					n.update(tickElapsed)
					d <- true
				}(n)
			}
		}
		for i := 0; i < size; i++ {
			<-d
		}
		close(c)
	}
	<-c
}

func (nc *NeedsController) NewColonist(k colonist.ColonistKey) {
	c := make(chan struct{})
	nc.actionc <- func() {
		nc.state[k] = map[needType]*need{
			Rest:        newNeed(50, -0.1),
			Food:        newNeed(50, -0.1),
			Water:       newNeed(50, -0.1),
			Security:    newNeed(50, -0.1),
			Belonging:   newNeed(50, -0.1),
			Fulfillment: newNeed(50, -0.1),
			Power:       newNeed(50, -0.1),
			Joy:         newNeed(50, -0.1),
			Wealth:      newNeed(50, -0.1),
			Status:      newNeed(50, -0.1),
		}
		close(c)
	}
	<-c
}

func (nc *NeedsController) GetValue(k colonist.ColonistKey, t needType) float64 {
	n := nc.state[k][t]

	return n.getValue()
}

func (nc *NeedsController) Adjust(k colonist.ColonistKey, t needType, adjustment float64) float64 {
	n := nc.state[k][t]
	n.setValue(n.getValue() + adjustment)
	return n.getValue()
}

func (nc *NeedsController) Target(k colonist.ColonistKey, t needType, target float64, duration float64) (cancel func()) {
	n := nc.state[k][t]

	return n.addTarget(target, duration)
}

func (nc *NeedsController) Ceiling(k colonist.ColonistKey, t needType, ceiling float64, duration float64) (cancel func()) {
	n := nc.state[k][t]

	return n.addCeiling(ceiling, duration)
}

func (nc *NeedsController) Floor(k colonist.ColonistKey, t needType, floor float64, duration float64) (cancel func()) {
	n := nc.state[k][t]

	return n.addFloor(floor, duration)
}

type need struct {
	state   *needState
	actionc chan func()
	quitc   chan struct{}
}

type needState struct {
	value    float64
	baseRate float64
	rates    []*expiringValue
	ceilings []*expiringValue
	floors   []*expiringValue
}

type expiringValue struct {
	value    float64
	duration float64
	elapsed  float64
}

func newNeed(startingValue float64, baseRate float64) *need {
	n := &need{
		state: &needState{
			value:    startingValue,
			baseRate: baseRate,
		},
		actionc: make(chan func()),
		quitc:   make(chan struct{}),
	}
	go n.start()
	return n
}

func (n *need) start() {
	for {
		select {
		case f := <-n.actionc:
			f()
		case <-n.quitc:
			return
		}
	}
}

// update processes how the need should be adjusted for the amount of tick time
// that has elapsed. Various rates, ceilings and floors may expire.
func (n *need) update(tickElapsed float64) {
	n.actionc <- func() {
		n.state.rates = updateExpiringValues(n.state.rates, tickElapsed)
		n.state.ceilings = updateExpiringValues(n.state.ceilings, tickElapsed)
		n.state.floors = updateExpiringValues(n.state.floors, tickElapsed)

		adjustment := sumExpiringValues(n.state.rates) * tickElapsed
		ceiling := minExpiringValue(n.state.ceilings)
		floor := maxExpiringValue(n.state.floors)

		n.state.value = math.Min(ceiling, math.Max(floor, n.state.value+adjustment))
	}
}

// addTarget adds a new target for a given duration and returns a cancel callback to remove the target.
func (n *need) addTarget(target float64, duration float64) (cancel func()) {
	c := make(chan func())
	n.actionc <- func() {
		// convert target to rate
		r := (n.state.value - target) / duration
		v := &expiringValue{
			value:    r,
			duration: duration,
		}
		cancel := func() {
			n.state.rates = removeExpiringValue(n.state.rates, v)
		}
		// add rate
		n.state.rates = append(n.state.rates, v)
		c <- cancel
	}
	return <-c
}

// addCeiling adds a new ceiling for an optional duration and returns a cancel callback to remove the ceiling.
func (n *need) addCeiling(ceiling float64, duration float64) (cancel func()) {
	c := make(chan func())
	n.actionc <- func() {
		// convert target to rate
		v := &expiringValue{
			value:    ceiling,
			duration: duration,
		}
		cancel := func() {
			n.state.ceilings = removeExpiringValue(n.state.ceilings, v)
		}
		n.state.ceilings = append(n.state.ceilings, v)
		c <- cancel
	}
	return <-c
}

// addFloor adds a new floor for an optional duration and returns a cancel callback to remove the floor.
func (n *need) addFloor(floor float64, duration float64) (cancel func()) {
	c := make(chan func())
	n.actionc <- func() {
		// convert target to rate
		v := &expiringValue{
			value:    floor,
			duration: duration,
		}
		cancel := func() {
			n.state.floors = removeExpiringValue(n.state.floors, v)
		}
		n.state.floors = append(n.state.floors, v)
		c <- cancel
	}
	return <-c
}

// getValue returns the current value of the need.
func (n *need) getValue() float64 {
	c := make(chan float64)
	n.actionc <- func() {
		c <- n.state.value
	}
	return <-c
}

// setValue sets the value of the need within a set ceiling and floor
func (n *need) setValue(v float64) {
	c := make(chan struct{})
	n.actionc <- func() {
		ceiling := minExpiringValue(n.state.ceilings)
		floor := maxExpiringValue(n.state.floors)

		n.state.value = math.Min(ceiling, math.Max(floor, v))
		close(c)
	}
	<-c
}

func sumExpiringValues(values []*expiringValue) float64 {
	var sum float64
	for _, v := range values {
		sum += v.value
	}
	return sum
}

// Returns the floor of this need (the need cannot go any lower than this value)
// If there is more than one floor, this method returns the highest floor.
func maxExpiringValue(values []*expiringValue) float64 {
	max := float64(0)
	for _, v := range values {
		if v.value > max {
			max = v.value
		}
	}
	return max
}

// Returns the ceiling of this need (the need cannot go any higher than this value)
// If there is more than one ceiling, this method returns the lowest ceiling.
// A default value of 100 is returned which is the default ceiling for all expiring values
func minExpiringValue(values []*expiringValue) float64 {
	min := float64(100)
	for _, v := range values {
		if v.value < min {
			min = v.value
		}
	}
	return min
}

func updateExpiringValues(values []*expiringValue, tickElapsed float64) []*expiringValue {
	var nonExpiredValues []*expiringValue
	for i, v := range values {
		if v.duration == 0 || v.elapsed < v.duration {
			v.elapsed += tickElapsed
			nonExpiredValues[i] = v
		}
	}
	return nonExpiredValues
}

func removeExpiringValue(values []*expiringValue, targetValue *expiringValue) []*expiringValue {
	for i, v := range values {
		if v == targetValue {
			values[i] = values[len(values)-1]
			values[len(values)-1] = v
			values = values[:len(values)-1]
		}
	}
	return values
}
