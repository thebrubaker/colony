package game

import (
	"math"

	wr "github.com/mroth/weightedrand"
)

type Attribute struct {
	Value float64
}

func (attribute *Attribute) Add(amount float64) {
	attribute.Value = math.Min(100, attribute.Value+amount)
}

func (attribute *Attribute) Sub(amount float64) {
	attribute.Value = math.Max(0, attribute.Value-amount)
}

type Colonist struct {
	Name          string     `json:name`
	CurrentAction *Action    `json:action`
	Thirst        *Attribute `json:thirst`
	Stress        *Attribute `json:stress`
	Energy        *Attribute `json:energy`
	Hunger        *Attribute `json:hunger`
	Inventory     *Inventory `json:inventory`
	GameState     *GameState
}

func (colonist *Colonist) OnTick(ticker *Ticker) {
	colonist.Hunger.Add(0.1 * ticker.Elapsed)
	colonist.Thirst.Add(0.1 * ticker.Elapsed)
	colonist.Stress.Add(0.1 * ticker.Elapsed)
	colonist.Energy.Sub(0.1 * ticker.Elapsed)

	colonist.Energy.Sub(colonist.CurrentAction.Type.EnergyCost * ticker.Elapsed)

	colonist.CurrentAction.OnTick()
}

func (colonist *Colonist) SetActionType(actionType *ActionType, gameState *GameState) {
	if colonist.CurrentAction != nil {
		colonist.CurrentAction.OnEnd()
	}

	colonist.CurrentAction = &Action{
		Colonist: colonist,
		Game:     gameState,
		Type:     actionType,
	}

	colonist.CurrentAction.OnStart()

	colonist.CurrentAction.SetTickExpiration()
}

func (colonist *Colonist) ContinueWithCurrentAction(ticker *Ticker) bool {
	if colonist.CurrentAction == nil {
		return false
	}

	if colonist.CurrentAction.IsExpired(ticker) {
		return false
	}

	return true
}

func (colonist *Colonist) ProcessActions(actionTypes []*ActionType, gameState *GameState) *Colonist {
	if colonist.ContinueWithCurrentAction(gameState.Ticker) {
		return colonist
	}

	choices := CreateChoices(colonist, actionTypes)

	actionType := wr.NewChooser(choices...).Pick().(*ActionType)

	colonist.SetActionType(actionType, gameState)

	return colonist
}

func CreateChoices(colonist *Colonist, actionTypes []*ActionType) []wr.Choice {
	choices := []wr.Choice{}

	for _, actionType := range actionTypes {
		if !actionType.IsAllowedForColonist(colonist) {
			continue
		}

		choice := MakeChoice(colonist, actionType)

		choices = append(choices, choice)
	}

	return choices
}

func MakeChoice(colonist *Colonist, actionType *ActionType) wr.Choice {
	utility := actionType.GetUtilityForColonist(colonist) * 100

	return wr.Choice{
		Item:   actionType,
		Weight: uint(math.Pow(utility, 3)),
	}
}
