package actions

import (
	"log"
	"math/rand"

	"github.com/jmcvetta/randutil"
	"github.com/thebrubaker/colony/actions/types"
)

func (ctx *Context) DetermineAction(a *Action) *Action {
	// action duration has ended
	if a.ActionComplete || a.TickProgress >= float64(a.TickDuration) {
		return ctx.NextAction(a.Type)
	}

	return a
}

func (ctx *Context) shouldEndEarly(a *Action) bool {
	// once a need is met, it ends the action
	if needType := a.Type.HasUtilityNeed(); needType != "" {
		return ctx.ActiveColonist.Needs[needType] <= 0
	}

	// desires never end early
	if desireType := a.Type.HasUtilityDesire(); desireType != "" {
		return false
	}

	// if it produces and the bag is full, it ends early
	if produces := a.Type.ProducesResources(); produces != nil {
		return ctx.ActiveColonist.Bag.IsFull()
	}

	// stops dropping off items after bag is empty
	if a.Type == types.HaulItems && ctx.ActiveColonist.Bag.IsEmpty() {
		return true
	}

	return false
}

func (ctx *Context) NextAction(previousAction types.Actionable) *Action {
	ctx.EndAction(previousAction)

	nextAction := ctx.SelectNextAction(previousAction)

	ctx.StartAction(nextAction.Type)

	return nextAction
}

// From all available actions, filter to those that meet pre-requirements
// then generate a weight score for how motivated the colonist is to choose
// that action (with some randomness).
func (ctx *Context) SelectNextAction(currentType types.Actionable) *Action {
	allowedActions := ctx.FilterActions(ColonistActions)

	action := ctx.ChooseWeightedAction(currentType, allowedActions)
	variance := float64(action.HasDuration()) * 0.2
	duration := types.TickDuration(float64(action.HasDuration()) - variance/2 + (rand.Float64() * variance))

	return &Action{
		Type:         action,
		TickDuration: duration,
	}
}

func (ctx *Context) FilterActions(actions []types.Actionable) []types.Actionable {
	var allowedActions []types.Actionable

	for _, a := range actions {
		// colonist has required resources that will be consumed
		if a.ConsumesResources() != nil && !ctx.MeetsResourceQuantities(a.ConsumesResources()) {
			continue
		}
		// will not produce resources while bag is full
		if a.ProducesResources() != nil && ctx.ActiveColonist.Bag.IsFull() {
			continue
		}

		allowedActions = append(allowedActions, a)
	}

	return allowedActions
}

// Goes through the game context to confirm the resources exist
// and are accessible to the colonist. Returns true if the resources
// are available to be consumed.
func (ctx *Context) MeetsResourceQuantities(resources []types.ConsumeResource) bool {
	for _, resourceQuantity := range resources {
		resource := resourceQuantity.Resource
		amount := resourceQuantity.Amount
		if !ctx.ActiveColonist.Bag.Has(resource, amount) {
			return false
		}
	}

	return true
}

type ActionWeight struct {
	Type   types.Actionable
	Weight float64
}

func (ctx *Context) ChooseWeightedAction(previousAction types.Actionable, actions []types.Actionable) types.Actionable {
	choices := ctx.GetWeightedChoices(previousAction, actions)

	choice, err := randutil.WeightedChoice(choices)
	if err != nil {
		log.Fatal("Failed to retrieve a weighted choice", err)
	}

	i, ok := choice.Item.(types.Actionable)
	if !ok {
		log.Fatal("Could not choose next action, choice was not actionable.")
	}

	return i
}

func (ctx *Context) GetWeightedChoices(currentAction types.Actionable, availableActions []types.Actionable) []randutil.Choice {
	var choices []randutil.Choice

	for _, action := range availableActions {
		choices = append(choices, randutil.Choice{
			Item:   action,
			Weight: ctx.GetUtility(currentAction, action),
		})
	}

	return choices
}

func (ctx *Context) GetUtility(currentAction types.Actionable, nextAction types.Actionable) int {
	if nextAction.WhenBagFull() {
		return 80
	}

	if needType := nextAction.HasUtilityNeed(); needType != "" {
		return int(ctx.ActiveColonist.Needs[needType])
	}

	if desireType := nextAction.HasUtilityDesire(); desireType != "" {
		return 100 - int(ctx.ActiveColonist.Desires[desireType])
	}

	return 10
}
