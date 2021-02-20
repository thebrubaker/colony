package types

import (
	"github.com/thebrubaker/colony/colonist"
	"github.com/thebrubaker/colony/need"
)

func InitTypes() []Actionable {
	return []Actionable{
		&PickBerries{},
		&DrinkFromStream{},
		&WatchClouds{},
		&GatherWood{},
	}
}

type StorageType uint8

const (
	ColonistBag StorageType = iota
	Stockpile
)

type Priority float64

const (
	FightOrFlight Priority = 80
	Survival      Priority = 60
	Social        Priority = 40
	Recreation    Priority = 20
	Job           Priority = 0
)

type TickDuration float64

const (
	Tedious  TickDuration = 20
	Slow     TickDuration = 15
	Moderate TickDuration = 12
	Fast     TickDuration = 8
	Fastest  TickDuration = 5
)

type EnergyCost float64

const (
	Exhausting EnergyCost = 0.2
	Hard       EnergyCost = 0.1
	Easy       EnergyCost = 0.05
	Easiest    EnergyCost = 0.01
)

type Actionable interface {
	Status() string
	EnergyCost() EnergyCost
	Duration() TickDuration
	Priority() Priority
}

type OnTick interface {
	Actionable
	OnTick()
}

type OnStart interface {
	Actionable
	OnStart()
}

type OnEnd interface {
	Actionable
	OnEnd()
}

type Weighted interface {
	Actionable
	Weight() float64
}

type GetUtility interface {
	Actionable
	GetUtility() float64
}

type SimpleNeed interface {
	Actionable
	Need() need.NeedType
}

type SimpleFulfillment interface {
	Actionable
	Satisfies() (need.NeedType, float64, func(float64) float64)
}

type SimpleAgitate interface {
	Actionable
	Agitates() (need.NeedType, float64, func(float64) float64)
}

type Gathers interface {
	Actionable
	Gather() (StorageType, interface{}, float64)
}

type EasedWeight interface {
	Actionable
	EaseWeight(v float64) float64
}
type SimpleSkillUp interface {
	Actionable
	SkillUp() (colonist.SkillType, float64, func(float64) float64)
}
