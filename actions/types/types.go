package types

type TickDuration float64

const (
	Tedious  TickDuration = 20
	Slow     TickDuration = 15
	Moderate TickDuration = 12
	Fast     TickDuration = 8
	Fastest  TickDuration = 5
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
	SatisfiesNeeds() []SatisfyNeed
	AgitatesNeeds() []AgitateNeed
	ProducesResources() []ProduceResource
	ConsumesResources() []ConsumeResource
	ImprovesSkills() []ImproveSkill
}
