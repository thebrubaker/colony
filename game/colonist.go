package game

import (
	"math"

	wr "github.com/mroth/weightedrand"
)

type Action struct {
	Colonist       *Colonist
	Type           *ActionType
	TickExpiration float64
}

func (action *Action) OnStart() {
	action.Type.OnStart(action.Colonist)
}

func (action *Action) OnTick() {
	action.Type.OnTick(action.Colonist)
}

func (action *Action) OnEnd() {
	action.Type.OnEnd(action.Colonist)
}

func (action *Action) IsExpired() bool {
	return action.Colonist.Game.Ticker.Count >= action.TickExpiration
}

type Colonist struct {
	Game          *GameState
	Name          string
	Status        string
	Thirst        float64
	Stress        float64
	CurrentAction *Action
}

func (colonist *Colonist) SubThirst(amount float64) {
	realAmount := amount * colonist.Game.Ticker.Elapsed

	colonist.Thirst = math.Max(0, colonist.Thirst-realAmount)
}

func (colonist *Colonist) AddThirst(amount float64) {
	realAmount := amount * colonist.Game.Ticker.Elapsed

	colonist.Thirst = math.Min(100, colonist.Thirst+realAmount)
}

func (colonist *Colonist) SubStress(amount float64) {
	realAmount := amount * colonist.Game.Ticker.Elapsed

	colonist.Stress = math.Max(0, colonist.Stress-realAmount)
}

func (colonist *Colonist) AddStress(amount float64) {
	realAmount := amount * colonist.Game.Ticker.Elapsed

	colonist.Stress = math.Min(100, colonist.Stress+realAmount)
}

func (colonist *Colonist) OnTick() {
	colonist.AddThirst(0.1)
	colonist.AddStress(0.1)

	colonist.CurrentAction.OnTick()
}

func (colonist *Colonist) SetActionType(actionType *ActionType) {
	if colonist.CurrentAction != nil {
		colonist.CurrentAction.OnEnd()
	}

	colonist.CurrentAction = &Action{
		Colonist:       colonist,
		Type:           actionType,
		TickExpiration: colonist.Game.Ticker.Count + actionType.Duration,
	}

	colonist.CurrentAction.OnStart()
}

func (colonist *Colonist) MakeChoice(actionType *ActionType) wr.Choice {
	return wr.Choice{
		Item:   actionType,
		Weight: actionType.GetUtility(colonist) ^ 3,
	}
}

func (colonist *Colonist) DetermineAction(actionTypes []*ActionType) {
	if colonist.CurrentAction != nil && !colonist.CurrentAction.IsExpired() {
		return // continue with current action
	}

	choices := []wr.Choice{}

	for _, actionType := range actionTypes {
		if actionType.IsAllowed(colonist) {
			choices = append(choices, colonist.MakeChoice(actionType))
		}
	}

	colonist.SetActionType(wr.NewChooser(choices...).Pick().(*ActionType))
}
