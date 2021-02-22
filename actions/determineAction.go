package actions

import (
	"log"

	"github.com/jmcvetta/randutil"
	"github.com/thebrubaker/colony/actions/types"
)

func (ctx *Context) DetermineAction(a *Action) *Action {
	if a != nil && a.TickProgress < float64(a.Type.HasDuration()) {
		return a
	}

	return ctx.NextAction(a.Type)
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
	allowedActions := ctx.FilterActions([]types.Actionable{
		types.DrinkGroundWater,
		types.GatherWildBerries,
		types.GatherWood,
		types.BasicRelax,
		types.BasicRest,
	})

	action := ctx.ChooseWeightedAction(currentType, allowedActions)

	return &Action{
		Type: action,
	}
}

func (ctx *Context) FilterActions(actions []types.Actionable) []types.Actionable {
	var allowedActions []types.Actionable

	for _, a := range actions {
		// colonist has required resources that will be consumed
		if !ctx.MeetsResourceQuantities(a.ConsumesResources()) {
			continue
		}
		allowedActions = append(allowedActions, a)
	}

	return allowedActions
}

func GetEasedValue(total float64, ease func(float64) float64, duration float64, tickProgress float64, tickElapsed float64) float64 {
	if ease == nil {
		return (total / duration) * tickElapsed
	}

	previousTick := ease(tickProgress / duration)
	currentTick := ease((tickProgress + tickElapsed) / duration)

	return total * (currentTick - previousTick)
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
	return 0
}
