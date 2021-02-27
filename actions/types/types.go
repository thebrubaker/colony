package types

import "github.com/thebrubaker/colony/colonist"

type TickDuration float64

const (
	Tedious  TickDuration = 60 * 3
	Slow     TickDuration = 60
	Moderate TickDuration = 45
	Fast     TickDuration = 30
	Fastest  TickDuration = 15
)

type Effort float64

const (
	Exhausting Effort = 0.2
	Hard       Effort = 0.15
	Demanding  Effort = 0.1
	Easy       Effort = 0.05
	Painless   Effort = 0.01
)

type Actionable interface {
	Status() []string
	TakesEffort() Effort
	HasDuration() TickDuration
	WhenBagFull() bool
	HasUtilityNeed() colonist.NeedType
	HasUtilityDesire() colonist.DesireType
	SatisfiesNeeds() []SatisfyNeed
	AgitatesNeeds() []AgitateNeed
	SatisfiesDesires() []SatisfyDesire
	ProducesResources() []ProduceResource
	ConsumesResources() []ConsumeResource
	ImprovesSkills() []ImproveSkill
}
