package game

import (
	"math"
	"math/rand"

	"github.com/fogleman/ease"
)

var (
	FightOrFlight float64 = 100
	Physiological float64 = 90
	Social        float64 = 80
	Obligation    float64 = 70
	Recreation    float64 = 60
)

type ActionType struct {
	ID         string                                                   `json:id`
	Status     string                                                   `json:status`
	Priority   float64                                                  `json:priority`
	EnergyCost float64                                                  `json:energy_cost`
	Duration   float64                                                  `json:duration`
	GetUtility func(actionType *ActionType, colonist *Colonist) float64 `json:-`
	IsAllowed  func(actionType *ActionType, colonist *Colonist) bool    `json:-`
	OnStart    func(action *Action, ticker *Ticker)                     `json:-`
	OnTick     func(action *Action, ticker *Ticker)                     `json:-`
	OnEnd      func(action *Action, ticker *Ticker)                     `json:-`
}

type Action struct {
	Game           *GameState  `json:-`
	Colonist       *Colonist   `json:-`
	Type           *ActionType `json:type`
	TickExpiration float64     `json:-`
}

func (action *Action) SetTickExpiration() {
	action.TickExpiration = action.Game.Ticker.Count + action.Type.Duration
}

func (action *Action) OnStart() {
	if action.Type.OnStart == nil {
		return
	}

	action.Type.OnStart(action, action.Game.Ticker)
}

func (action *Action) OnTick() {
	if action.Type.OnTick == nil {
		return
	}

	action.Type.OnTick(action, action.Game.Ticker)
}

func (action *Action) OnEnd() {
	if action.Type.OnEnd == nil {
		return
	}

	action.Type.OnEnd(action, action.Game.Ticker)
}

func (action *Action) IsExpired(ticker *Ticker) bool {
	return ticker.Count >= action.TickExpiration
}

func (actionType *ActionType) GetUtilityForColonist(colonist *Colonist) float64 {
	if actionType.GetUtility == nil {
		return 0.1
	}

	return actionType.GetUtility(actionType, colonist)
}

func (actionType *ActionType) IsAllowedForColonist(colonist *Colonist) bool {
	if actionType.IsAllowed == nil {
		return true
	}

	return actionType.IsAllowed(actionType, colonist)
}

var WatchClouds *ActionType = &ActionType{
	ID:         "1",
	Status:     "Watching the clouds.",
	Priority:   Recreation,
	EnergyCost: 0,
	Duration:   9,
	GetUtility: func(actionType *ActionType, colonist *Colonist) float64 {
		return ease.InQuad(colonist.Stress.Value / 100)
	},
	OnTick: func(action *Action, ticker *Ticker) {
		action.Colonist.Stress.Sub(4 * ticker.Elapsed)
	},
}

var SearchForWater *ActionType = &ActionType{
	ID:         "2",
	Status:     "Collecting water from a stream.",
	Priority:   Obligation,
	EnergyCost: 2,
	Duration:   4,
	GetUtility: func(actionType *ActionType, colonist *Colonist) float64 {
		return ease.InQuad(colonist.Thirst.Value / 100)
	},
	IsAllowed: func(actionType *ActionType, colonist *Colonist) bool {
		return colonist.GameState.Inventory.Water <= 0
	},
	OnEnd: func(action *Action, ticker *Ticker) {
		action.Game.Inventory.Water += 3
	},
}

var DrinkNaturalWater *ActionType = &ActionType{
	ID:         "3",
	Status:     "Drinking water from a stream.",
	Priority:   Physiological,
	EnergyCost: 1,
	Duration:   5,
	GetUtility: func(actionType *ActionType, colonist *Colonist) float64 {
		return ease.InQuart(colonist.Thirst.Value / 100)
	},
	OnTick: func(action *Action, ticker *Ticker) {
		action.Colonist.Thirst.Sub(4 * ticker.Elapsed)
		action.Colonist.Stress.Add(1 * ticker.Elapsed)
	},
}

var DrinkWater *ActionType = &ActionType{
	ID:         "3",
	Status:     "Drinking water.",
	Priority:   Physiological,
	EnergyCost: 1,
	Duration:   4,
	GetUtility: func(actionType *ActionType, colonist *Colonist) float64 {
		return ease.InQuart(colonist.Thirst.Value / 100)
	},
	IsAllowed: func(actionType *ActionType, colonist *Colonist) bool {
		return colonist.GameState.Inventory.Water > 0
	},
	OnStart: func(action *Action, ticker *Ticker) {
		action.Game.Inventory.Water--
		action.Colonist.Inventory.Water++
	},
	OnEnd: func(action *Action, ticker *Ticker) {
		action.Colonist.Inventory.Water--
	},
	OnTick: func(action *Action, ticker *Ticker) {
		action.Colonist.Thirst.Sub(5 * ticker.Elapsed)
	},
}

var StandIdle *ActionType = &ActionType{
	ID:         "4",
	Status:     "Standing Still.",
	Priority:   Recreation,
	EnergyCost: 1,
	Duration:   3,
}

var WakingUp *ActionType = &ActionType{
	ID:         "5",
	Status:     "Waking up from cryosleep.",
	Priority:   Recreation,
	EnergyCost: 1,
	Duration:   4,
	OnTick: func(action *Action, ticker *Ticker) {
		action.Colonist.Stress.Add(15 * ticker.Elapsed)
	},
}

var HuntForFood *ActionType = &ActionType{
	ID:         "6",
	Status:     "Hunting a wild boar.",
	Priority:   Obligation,
	EnergyCost: 2,
	Duration:   4,
	GetUtility: func(actionType *ActionType, colonist *Colonist) float64 {
		return ease.InQuad(colonist.Hunger.Value / 100)
	},
	IsAllowed: func(actionType *ActionType, colonist *Colonist) bool {
		return colonist.GameState.Inventory.RawFood <= 0
	},
	OnEnd: func(action *Action, ticker *Ticker) {
		if rand.Float64() <= 0.1 {
			action.Game.Inventory.RawFood += 3
		}
	},
}

var ConsumeRawFood *ActionType = &ActionType{
	ID:         "7",
	Status:     "Consuming raw berries.",
	Priority:   Physiological,
	EnergyCost: 2,
	Duration:   6,
	GetUtility: func(actionType *ActionType, colonist *Colonist) float64 {
		return ease.InQuad(colonist.Hunger.Value / 100)
	},
	IsAllowed: func(actionType *ActionType, colonist *Colonist) bool {
		return colonist.GameState.Inventory.SimpleMeal < 0
	},
	OnTick: func(action *Action, ticker *Ticker) {
		action.Colonist.Hunger.Sub(2 * ticker.Elapsed)
	},
}

var ConsumeCookedFood *ActionType = &ActionType{
	ID:         "8",
	Status:     "Eating a cooked meal.",
	Priority:   Physiological,
	EnergyCost: 1,
	Duration:   3,
	GetUtility: func(actionType *ActionType, colonist *Colonist) float64 {
		return ease.InQuad(colonist.Hunger.Value / 100)
	},
	IsAllowed: func(actionType *ActionType, colonist *Colonist) bool {
		return colonist.GameState.Inventory.SimpleMeal > 0
	},
	OnEnd: func(action *Action, ticker *Ticker) {
		action.Game.Inventory.SimpleMeal--
		action.Colonist.Hunger.Sub(100)
	},
}

var CookRawFood *ActionType = &ActionType{
	ID:         "9",
	Status:     "Cooking food.",
	Priority:   Physiological,
	EnergyCost: 2,
	Duration:   6,
	IsAllowed: func(actionType *ActionType, colonist *Colonist) bool {
		return colonist.GameState.Inventory.RawFood >= 3
	},
	GetUtility: func(actionType *ActionType, colonist *Colonist) float64 {
		return ease.InQuad(colonist.Hunger.Value / 100)
	},
	OnStart: func(action *Action, ticker *Ticker) {
		action.Game.Inventory.RawFood -= 3
		action.Colonist.Inventory.RawFood += 3
	},
	OnEnd: func(action *Action, ticker *Ticker) {
		action.Colonist.Inventory.RawFood -= 3
		action.Game.Inventory.SimpleMeal += 3
	},
}

var RestingOnGround *ActionType = &ActionType{
	ID:         "10",
	Status:     "Resting on a pile of leaves.",
	Priority:   Physiological,
	EnergyCost: 0,
	Duration:   1,
	GetUtility: func(actionType *ActionType, colonist *Colonist) float64 {
		// Colonist wants to continue with the current action until satisfied
		if actionType == colonist.CurrentAction.Type && colonist.Energy.Value < 95 {
			return 1
		}

		t := GetMinEaseValue(100-colonist.Energy.Value, 80)

		return ease.InQuad(t)
	},
	OnTick: func(action *Action, ticker *Ticker) {
		action.Colonist.Energy.Add(5 * ticker.Elapsed)
	},
}

func GetMinEaseValue(need float64, min float64) float64 {
	if need < min {
		return 0
	}

	return math.Max(0, need-min) / (100 - min)
}

var ActionTypes []*ActionType = []*ActionType{
	WatchClouds,
	StandIdle,
	SearchForWater,
	DrinkWater,
	HuntForFood,
	ConsumeRawFood,
	RestingOnGround,
	ConsumeCookedFood,
	CookRawFood,
}
