// package game

// import (
// 	"math"

// 	wr "github.com/mroth/weightedrand"
// )

// type Attribute struct {
// 	Value float64 `json:value`
// }

// func (attribute *Attribute) Add(amount float64) {
// 	attribute.Value = math.Min(100, attribute.Value+amount)
// }

// func (attribute *Attribute) Sub(amount float64) {
// 	attribute.Value = math.Max(0, attribute.Value-amount)
// }

// type Colonist struct {
// 	Name   string `json:name`
// 	Status string `json:status`

// 	Hands     Carryable   `json:hands`
// 	Bag       []Slottable `json:bag`
// 	Equipment *Equipment  `json:equipment`

// 	Needs Needs `json:needs`
// }

// type Needs struct {
// 	Thirst     *Attribute `json:thirst`
// 	Stress     *Attribute `json:stress`
// 	Exhaustion *Attribute `json:exhaustion`
// 	Hunger     *Attribute `json:hunger`
// }

// func (_ *Needs) Add(a *Attribute, amount float64) {
// 	a.Value = math.Min(100, a.Value+amount)
// }

// func (_ *Needs) Sub(a *Attribute, amount float64) {
// 	a.Value = math.Max(0, a.Value-amount)
// }

// func (colonist *Colonist) OnTick(currentAction *Action, ticker *Ticker) {
// 	colonist.Needs.Add(colonist.Needs.Hunger, 0.3*ticker.Elapsed)
// 	colonist.Needs.Add(colonist.Needs.Thirst, 0.3*ticker.Elapsed)

// 	cost := currentAction.Type.EnergyCost

// 	colonist.Needs.Add(colonist.Needs.Exhaustion, cost*ticker.Elapsed)

// 	currentAction.OnTick()
// }

// func SetActionType(actionType *ActionType, currentAction *Action) {
// 	if currentAction != nil {
// 		currentAction.OnEnd()
// 	}

// 	currentAction = &Action{
// 		Type: actionType,
// 	}

// 	currentAction.OnStart()

// 	currentAction.SetTickExpiration()
// }

// func ProcessActions(colonist *Colonist, currentAction *Action, ticker *Ticker) *Action {
// 	if continueAction(currentAction, ticker) {
// 		return currentAction
// 	}

// 	choices := CreateChoices(ActionTypes, colonist)

// 	// for _, choice := range colonist.Choices {
// 	// 	fmt.Printf("%s: %d\n", choice.Item.(*ActionType).Status, choice.Weight)
// 	// }

// 	// os.Exit(1)

// 	actionType := wr.NewChooser(colonist.Choices...).Pick().(*ActionType)

// 	SetActionType(actionType)

// 	return colonist
// }

// func continueAction(currentAction *Action, ticker *Ticker) bool {
// 	if currentAction == nil {
// 		return false
// 	}

// 	if currentAction.IsExpired(ticker) {
// 		return false
// 	}

// 	return true
// }

// func CreateChoices(actionTypes []*ActionType, colonist *Colonist) {
// 	for _, actionType := range actionTypes {
// 		if CannotTakeAction(actionType, colonist) {
// 			continue
// 		}

// 		weight := colonist.GetWeight(actionType)

// 		colonist.Choices = append(colonist.Choices, wr.Choice{
// 			Item:   actionType,
// 			Weight: uint(math.Pow(weight, 3)),
// 		})
// 	}
// }

// func (colonist *Colonist) GetWeight(actionType *ActionType) float64 {
// 	if actionType.GetUtility == nil {
// 		return actionType.Priority - 9
// 	}

// 	return actionType.GetUtility(actionType, colonist) * actionType.Priority
// }

// func CanTakeAction(actionType *ActionType, colonist *Colonist) bool {
// 	if actionType.IsAllowed == nil {
// 		return true
// 	}

// 	return actionType.IsAllowed(actionType, colonist)
// }

// func CannotTakeAction(actionType *ActionType, colonist *Colonist) bool {
// 	return !CanTakeAction(actionType, colonist)
// }
