package types

func InitTypes() []Buildable {
	return []Buildable{}
}

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

type Buildable interface {
	Name() string
	Description() string
	EnergyCost() EnergyCost
	Duration() TickDuration
}

type OnTick interface {
	Buildable
	OnTick()
}

type OnStart interface {
	Buildable
	OnStart()
}

type OnEnd interface {
	Buildable
	OnEnd()
}
