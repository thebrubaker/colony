package game

import (
	"math/rand"
)

const (
	Survival   uint = 1
	Duty       uint = 2
	Recreation uint = 3
)

type ActionType struct {
	Status     string
	Priority   uint
	EnergyCost float64
	Duration   float64
	GetUtility func(colonist *Colonist) uint
	IsAllowed  func(colonist *Colonist) bool
	OnStart    func(colonist *Colonist)
	OnTick     func(colonist *Colonist)
	OnEnd      func(colonist *Colonist)
}

var WatchClouds *ActionType = &ActionType{
	Status:     "Watching the clouds.",
	Priority:   Recreation,
	EnergyCost: 0,
	Duration:   15,
	GetUtility: func(colonist *Colonist) uint {
		if colonist.Stress > 0 {
			return 40
		}

		return 5
	},
	IsAllowed: func(colonist *Colonist) bool {
		return true
	},
	OnStart: func(colonist *Colonist) {},
	OnTick: func(colonist *Colonist) {
		colonist.SubStress(4)
	},
	OnEnd: func(colonist *Colonist) {},
}

var SearchForWater *ActionType = &ActionType{
	Status:     "Searching for water.",
	Priority:   Survival,
	EnergyCost: 2,
	Duration:   4,
	GetUtility: func(colonist *Colonist) uint {
		if colonist.Game.Inventory["water"] <= 0 {
			return 5
		}

		return uint(colonist.Thirst)
	},
	IsAllowed: func(colonist *Colonist) bool {
		return true
	},
	OnStart: func(colonist *Colonist) {
	},
	OnTick: func(colonist *Colonist) {},
	OnEnd: func(colonist *Colonist) {
		if rand.Float64() <= 0.6 {
			colonist.Game.Inventory["water"] += 3
		}
	},
}

var DrinkWater *ActionType = &ActionType{
	Status:     "Drinking water.",
	Priority:   Survival,
	EnergyCost: 1,
	Duration:   10,
	GetUtility: func(colonist *Colonist) uint {
		return uint(colonist.Thirst)
	},
	IsAllowed: func(colonist *Colonist) bool {
		return colonist.Game.Inventory["water"] > 0
	},
	OnStart: func(colonist *Colonist) {
		colonist.Game.Inventory["water"]--
	},
	OnTick: func(colonist *Colonist) {
		colonist.SubThirst(5)
	},
	OnEnd: func(colonist *Colonist) {},
}

var StandIdle *ActionType = &ActionType{
	Status:     "Standing Still.",
	Priority:   Recreation,
	EnergyCost: 1,
	Duration:   5,
	GetUtility: func(colonist *Colonist) uint {
		return 5
	},
	IsAllowed: func(colonist *Colonist) bool {
		return true
	},
	OnStart: func(colonist *Colonist) {},
	OnTick:  func(colonist *Colonist) {},
	OnEnd:   func(colonist *Colonist) {},
}

var WakingUp *ActionType = &ActionType{
	Status:     "Waking up from cryosleep.",
	Priority:   Recreation,
	EnergyCost: 1,
	Duration:   4,
	GetUtility: func(colonist *Colonist) uint {
		return 100
	},
	IsAllowed: func(colonist *Colonist) bool {
		return false
	},
	OnStart: func(colonist *Colonist) {},
	OnTick: func(colonist *Colonist) {
		colonist.AddStress(15)
	},
	OnEnd: func(colonist *Colonist) {},
}
