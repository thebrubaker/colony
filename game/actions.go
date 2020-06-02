// package game

// import (
// 	"math"
// 	"math/rand"

// 	"github.com/fogleman/ease"
// )

// var (
// 	FightOrFlight float64 = 100
// 	Physiological float64 = 90
// 	Social        float64 = 80
// 	Obligation    float64 = 70
// 	Recreation    float64 = 60
// 	NoPriority    float64 = 50
// )

// type ActionType struct {
// 	Status     string                                                   `json:status`
// 	Priority   float64                                                  `json:priority`
// 	EnergyCost float64                                                  `json:energy_cost`
// 	Duration   float64                                                  `json:duration`
// 	GetUtility func(actionType *ActionType, colonist *Colonist) float64 `json:-`
// 	IsAllowed  func(actionType *ActionType, colonist *Colonist) bool    `json:-`
// 	OnStart    func(action *Action, ticker *Ticker)                     `json:-`
// 	OnTick     func(action *Action, ticker *Ticker)                     `json:-`
// 	OnEnd      func(action *Action, ticker *Ticker)                     `json:-`
// }

// type Action struct {
// 	Game           *GameState  `json:-`
// 	Colonist       *Colonist   `json:-`
// 	Type           *ActionType `json:type`
// 	TickExpiration float64     `json:-`
// }

// func (action *Action) SetTickExpiration() {
// 	action.TickExpiration = action.Game.Ticker.Count + action.Type.Duration
// }

// func (action *Action) OnStart() {
// 	if action.Type.OnStart == nil {
// 		return
// 	}

// 	action.Type.OnStart(action, action.Game.Ticker)
// }

// func (action *Action) OnTick() {
// 	if action.Type.OnTick == nil {
// 		return
// 	}

// 	action.Type.OnTick(action, action.Game.Ticker)
// }

// func (action *Action) OnEnd() {
// 	if action.Type.OnEnd == nil {
// 		return
// 	}

// 	action.Type.OnEnd(action, action.Game.Ticker)
// }

// func (action *Action) IsExpired(ticker *Ticker) bool {
// 	return ticker.Count >= action.TickExpiration
// }

// var WatchClouds *ActionType = &ActionType{
// 	Status:     "Watching the clouds",
// 	Priority:   Recreation,
// 	EnergyCost: 0,
// 	Duration:   8,
// 	GetUtility: func(actionType *ActionType, colonist *Colonist) float64 {
// 		return ease.InOutQuad(colonist.Stress.Value / 100)
// 	},
// 	OnTick: func(action *Action, ticker *Ticker) {
// 		action.Colonist.Stress.Sub(2 * ticker.Elapsed)
// 	},
// }

// var ExploringStream *ActionType = &ActionType{
// 	Status:     "Exploring a nearby stream",
// 	Priority:   Recreation,
// 	EnergyCost: 0,
// 	Duration:   11,
// 	GetUtility: func(actionType *ActionType, colonist *Colonist) float64 {
// 		return ease.InOutQuad(colonist.Stress.Value / 100)
// 	},
// 	OnTick: func(action *Action, ticker *Ticker) {
// 		action.Colonist.Stress.Sub(2 * ticker.Elapsed)
// 	},
// }

// var ShortWalk *ActionType = &ActionType{
// 	Status:     "Taking a short walk",
// 	Priority:   Recreation,
// 	EnergyCost: 0,
// 	Duration:   12,
// 	GetUtility: func(actionType *ActionType, colonist *Colonist) float64 {
// 		return ease.InOutQuad(colonist.Stress.Value / 100)
// 	},
// 	OnTick: func(action *Action, ticker *Ticker) {
// 		action.Colonist.Stress.Sub(2 * ticker.Elapsed)
// 	},
// }

// var SearchForWater *ActionType = &ActionType{
// 	Status:     "Collecting water from a stream",
// 	Priority:   Obligation,
// 	EnergyCost: 1,
// 	Duration:   8,
// 	GetUtility: func(actionType *ActionType, colonist *Colonist) float64 {
// 		return 0.6
// 	},
// 	IsAllowed: func(actionType *ActionType, colonist *Colonist) bool {
// 		return true
// 	},
// 	OnEnd: func(action *Action, ticker *Ticker) {
// 		action.Game.Stockpile.Water += 5
// 	},
// }

// var DrinkWater *ActionType = &ActionType{
// 	Status:     "Drinking water",
// 	Priority:   Physiological,
// 	EnergyCost: 1,
// 	Duration:   6,
// 	GetUtility: func(actionType *ActionType, colonist *Colonist) float64 {
// 		return ease.InQuart(colonist.Thirst.Value / 100)
// 	},
// 	IsAllowed: func(actionType *ActionType, colonist *Colonist) bool {
// 		return colonist.GameState.Stockpile.Water > 0
// 	},
// 	OnStart: func(action *Action, ticker *Ticker) {
// 		action.Game.Stockpile.Water--
// 	},
// 	OnTick: func(action *Action, ticker *Ticker) {
// 		action.Colonist.Thirst.Sub(10 * ticker.Elapsed)
// 	},
// }

// var StandIdle *ActionType = &ActionType{
// 	Status:     "Standing Still",
// 	Priority:   NoPriority,
// 	EnergyCost: 0.1,
// 	Duration:   2,
// }

// var WakingUp *ActionType = &ActionType{
// 	Status:     "Waking up from cryosleep",
// 	Priority:   Recreation,
// 	EnergyCost: 0,
// 	Duration:   3,
// 	OnTick: func(action *Action, ticker *Ticker) {
// 		action.Colonist.Stress.Add(25 * ticker.Elapsed)
// 	},
// }

// var ConsumeRawFood *ActionType = &ActionType{
// 	Status:     "Consuming raw berries",
// 	Priority:   Physiological,
// 	EnergyCost: 1,
// 	Duration:   8,
// 	GetUtility: func(actionType *ActionType, colonist *Colonist) float64 {
// 		return ease.InOutQuad(colonist.Hunger.Value / 100)
// 	},
// 	IsAllowed: func(actionType *ActionType, colonist *Colonist) bool {
// 		return colonist.GameState.Stockpile.SimpleMeal == 0
// 	},
// 	OnTick: func(action *Action, ticker *Ticker) {
// 		action.Colonist.Hunger.Sub(10 * ticker.Elapsed)
// 	},
// }

// var ConsumeCookedFood *ActionType = &ActionType{
// 	Status:     "Eating a cooked meal",
// 	Priority:   Physiological,
// 	EnergyCost: 1,
// 	Duration:   8,
// 	GetUtility: func(actionType *ActionType, colonist *Colonist) float64 {
// 		return ease.InOutQuad(colonist.Hunger.Value / 100)
// 	},
// 	IsAllowed: func(actionType *ActionType, colonist *Colonist) bool {
// 		if colonist.GameState.Stockpile.SimpleMeal > 0 {
// 			return true
// 		}

// 		return false
// 	},
// 	OnStart: func(action *Action, ticker *Ticker) {
// 		action.Game.Stockpile.SimpleMeal--
// 	},
// 	OnTick: func(action *Action, ticker *Ticker) {
// 		action.Colonist.Hunger.Sub(5 * ticker.Elapsed)
// 	},
// }

// var CookRawFood *ActionType = &ActionType{
// 	Status:     "Cooking food",
// 	Priority:   Physiological,
// 	EnergyCost: 1,
// 	Duration:   12,
// 	IsAllowed: func(actionType *ActionType, colonist *Colonist) bool {
// 		if colonist.GameState.Stockpile.CookingFire == 0 {
// 			return false
// 		}

// 		if colonist.GameState.Stockpile.RawFood < 3 {
// 			return false
// 		}

// 		return true
// 	},
// 	GetUtility: func(actionType *ActionType, colonist *Colonist) float64 {
// 		return ease.InOutQuad(colonist.Hunger.Value / 100)
// 	},
// 	OnStart: func(action *Action, ticker *Ticker) {
// 		action.Game.Stockpile.RawFood -= 3
// 	},
// 	OnEnd: func(action *Action, ticker *Ticker) {
// 		action.Game.Stockpile.SimpleMeal += 3
// 	},
// }

// var RestingOnGround *ActionType = &ActionType{
// 	Status:     "Sleeping on the cold floor",
// 	Priority:   Physiological,
// 	EnergyCost: 0,
// 	Duration:   2,
// 	GetUtility: func(actionType *ActionType, colonist *Colonist) float64 {
// 		// Colonist wants to continue with the current action until satisfied
// 		if actionType == colonist.CurrentAction.Type && colonist.Exhaustion.Value >= 5 {
// 			return 1
// 		}

// 		return ease.InExpo(colonist.Exhaustion.Value / 100)
// 	},
// 	OnTick: func(action *Action, ticker *Ticker) {
// 		action.Colonist.Exhaustion.Sub(5 * ticker.Elapsed)
// 	},
// }

// var GatherWood *ActionType = &ActionType{
// 	Status:     "Gathering wood from the forest",
// 	Priority:   Obligation,
// 	EnergyCost: 2,
// 	Duration:   8,
// 	IsAllowed: func(actionType *ActionType, colonist *Colonist) bool {
// 		return true
// 	},
// 	GetUtility: func(actionType *ActionType, colonist *Colonist) float64 {
// 		return 0.6
// 	},
// 	OnEnd: func(action *Action, ticker *Ticker) {
// 		action.Game.Stockpile.Wood += uint(1 + rand.Float64()*3)
// 	},
// }

// var CreateHuntingSpear *ActionType = &ActionType{
// 	Status:     "Creating a hunting spear",
// 	Priority:   Obligation,
// 	EnergyCost: 2,
// 	Duration:   10,
// 	IsAllowed: func(actionType *ActionType, colonist *Colonist) bool {
// 		return colonist.GameState.Stockpile.Wood > 0
// 	},
// 	GetUtility: func(actionType *ActionType, colonist *Colonist) float64 {
// 		return 0.6
// 	},
// 	OnStart: func(action *Action, ticker *Ticker) {
// 		action.Game.Stockpile.Wood -= 1
// 	},
// 	OnEnd: func(action *Action, ticker *Ticker) {
// 		action.Game.Stockpile.HuntingSpear += 1
// 	},
// }

// var BuildCookingFire *ActionType = &ActionType{
// 	Status:     "Building a cooking fire",
// 	Priority:   Obligation,
// 	EnergyCost: 2,
// 	Duration:   10,
// 	IsAllowed: func(actionType *ActionType, colonist *Colonist) bool {
// 		if colonist.GameState.Stockpile.CookingFire > 0 {
// 			return false
// 		}

// 		return colonist.GameState.Stockpile.Wood >= 6
// 	},
// 	GetUtility: func(actionType *ActionType, colonist *Colonist) float64 {
// 		return 0.6
// 	},
// 	OnStart: func(action *Action, ticker *Ticker) {
// 		action.Game.Stockpile.Wood -= 6
// 	},
// 	OnEnd: func(action *Action, ticker *Ticker) {
// 		action.Game.Stockpile.CookingFire += 1
// 	},
// }

// var EquipSpear *ActionType = &ActionType{
// 	Status:     "Taking a hunting spear",
// 	Priority:   Obligation,
// 	EnergyCost: 1,
// 	Duration:   6,
// 	IsAllowed: func(actionType *ActionType, colonist *Colonist) bool {
// 		if colonist.Equipment.HuntingSpear > 0 {
// 			return false
// 		}

// 		if colonist.GameState.Stockpile.HuntingSpear == 0 {
// 			return false
// 		}

// 		return true
// 	},
// 	GetUtility: func(actionType *ActionType, colonist *Colonist) float64 {
// 		return 0.6
// 	},
// 	OnStart: func(action *Action, ticker *Ticker) {
// 		action.Game.Stockpile.HuntingSpear -= 1
// 	},
// 	OnEnd: func(action *Action, ticker *Ticker) {
// 		action.Colonist.Equipment.HuntingSpear += 1
// 	},
// }

// var HuntWildAnimal *ActionType = &ActionType{
// 	Status:     "Hunting a wild animal",
// 	Priority:   Obligation,
// 	EnergyCost: 2,
// 	Duration:   10,
// 	GetUtility: func(actionType *ActionType, colonist *Colonist) float64 {
// 		return 0.6
// 	},
// 	IsAllowed: func(actionType *ActionType, colonist *Colonist) bool {
// 		return colonist.Equipment.HuntingSpear > 0
// 	},
// 	OnEnd: func(action *Action, ticker *Ticker) {
// 		if rand.Float64() <= 0.8 {
// 			action.Game.Stockpile.RawFood += uint(rand.Float64() * 8)
// 		}
// 	},
// }

// var EatingFood *ActionType = &ActionType{
// 	Status:     "Hunting a wild animal",
// 	Priority:   Obligation,
// 	EnergyCost: 2,
// 	Duration:   10,
// 	GetUtility: func(actionType *ActionType, colonist *Colonist) float64 {
// 		return ease.InOutQuint(colonist.Hunger.Value / 100)
// 	},
// 	IsAllowed: func(actionType *ActionType, colonist *Colonist) bool {
// 		return true
// 	},
// 	OnStart: func(action *Action, ticker *Ticker) {
// 		action.Game.Stockpile.SimpleMeal > 0 {
// 			action.Colonist.Equipped.Hands += 1
// 		}

// 		action.Game.Stockpile.UncookedFood > 0 {
// 			action.Colonist.Equipped.Hands += 1
// 		}
// 	},
// 	OnEnd: func(action *Action, ticker *Ticker) {
// 		action.Game.Stockpile.SimpleMeal > 0 {
// 			action.Game.Stockpile.SimpleMeal -= 1
// 		}
// 	},
// }

// func GetMinEaseValue(need float64, min float64) float64 {
// 	if need < min {
// 		return 0
// 	}

// 	return math.Max(0, need-min) / (100 - min)
// }

// var ActionTypes []*ActionType = []*ActionType{
// 	EatingFood,
// 	DrinkingWater,
// 	Sleeping,
// 	// BuildCookingFire,
// 	// CreateHuntingSpear,
// 	// ConsumeCookedFood,
// 	// ConsumeRawFood,
// 	// CookRawFood,
// 	// DrinkWater,
// 	// EquipSpear,
// 	// ExploringStream,
// 	// GatherWood,
// 	// HuntWildAnimal,
// 	// RestingOnGround,
// 	// SearchForWater,
// 	// ShortWalk,
// 	// StandIdle,
// 	// WatchClouds,
// }
