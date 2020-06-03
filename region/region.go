package region

import (
	"github.com/fogleman/ease"
	"github.com/thebrubaker/colony/colonist"
)

type ColonistAction struct {
	Action       interface{}
	TickProgress float64
	TickEnds     float64
}

func (a *ColonistAction) processAction(tickElapsed float64, region *Region, colonist *colonist.Colonist) *ColonistAction {
	a.TickProgress += tickElapsed
	return a
}

type Region struct {
	Colonists map[string]*colonist.Colonist `json:"colonists"`
	Stockpile []interface{}                 `json:"stockpile"`
	actions   map[string]*ColonistAction
}

// Update processes all colonist actions in the region
func (r *Region) Update(tickElapsed float64) {
	for name, colonist := range r.Colonists {
		r.actions[name] = r.updateActions(tickElapsed, colonist, r.actions[name])

	}
}

func (r *Region) updateActions(tickElapsed float64, colonist *colonist.Colonist, currentAction *ColonistAction) *ColonistAction {
	if currentAction.TickProgress < currentAction.TickEnds {
		return currentAction.processAction(tickElapsed, r, colonist)
	}

	return NextAction(r, colonist, currentAction)
}

func NextAction(region *Region, colonist *colonist.Colonist, currentAction *ColonistAction) *ColonistAction {

}

type Context struct {
	Region        *Region
	Colonist      *colonist.Colonist
	CurrentAction *ColonistAction
	TickElapsed   float64
}

type Action interface {
	EnergyCost() float64
	Duration() float64
	Priority() float64
}

type SimpleNeed interface {
	Need() (float64, float64)
}

type SimpleFulfillment interface {
	Satisfies() (string, float64)
}

type Weighted interface {
	Weight() float64
}

type EasedWeight interface {
	EaseWeight(v float64) float64
}

type Prioritized interface {
	Priority() float64
}

type PickBerries struct {
	Context *Context
}

func (a *PickBerries) EnergyCost() float64 {
	return 2
}

func (a *PickBerries) Duration() float64 {
	return 10
}

func (a *PickBerries) Priority() float64 {
	return 60
}

func (a *PickBerries) Need() (float64, float64) {
	return a.Context.Colonist.Needs.Hunger, 80
}

func (a *PickBerries) EaseWeight(w float64) float64 {
	return ease.InExpo(w)
}

func (a *PickBerries) Satisfies() (string, float64) {
	return "Hunger", 1
}

func GetWeight(i SimpleNeed) float64 {
	need, threshold := i.Need()

	if need < threshold {
		return 0
	}

	w := (need - threshold) / (100 - threshold)

	if e, ok := i.(EasedWeight); ok {
		return e.EaseWeight(w)
	}

	return w
}

func GetUtility(a Prioritized) float64 {
	if i, ok := a.(SimpleNeed); ok {
		return a.Priority() + GetWeight(i)*20
	}

	if i, ok := a.(Weighted); ok {
		return a.Priority() + i.Weight()*20
	}

	return a.Priority() + 10
}

func OnTick(a Action, c *Context) {
	if i, ok := a.(SimpleFulfillment); ok {
		key, value := i.Satisfies()
		c.Colonist.Needs.Decrease(key, value*c.TickElapsed)
	}
}
