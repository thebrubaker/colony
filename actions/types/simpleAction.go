package types

import (
	"github.com/thebrubaker/colony/colonist"
	"github.com/thebrubaker/colony/resources"
)

type SatisfyNeed struct {
	NeedType colonist.NeedType
	Total    float64
	Ease     func(float64) float64
}

type AgitateNeed struct {
	NeedType colonist.NeedType
	Total    float64
	Ease     func(float64) float64
}

type SatisfyDesire struct {
	DesireType colonist.DesireType
	Total      float64
	Ease       func(float64) float64
}

type ImproveSkill struct {
	Skill         colonist.SkillType
	ChancePerTick float64
}

type ProduceResource struct {
	Resource      *resources.SimpleResource
	ChancePerTick float64
}

type ConsumeResource struct {
	Resource *resources.SimpleResource
	Amount   uint
}

type SimpleAction struct {
	status            []string
	effort            Effort
	duration          TickDuration
	utilityNeed       colonist.NeedType
	utilityDesire     colonist.DesireType
	improvesSkills    []ImproveSkill
	satisfiesNeeds    []SatisfyNeed
	satisfiesDesires  []SatisfyDesire
	agitatesNeeds     []AgitateNeed
	producesResources []ProduceResource
	consumesResources []ConsumeResource
}

func (a *SimpleAction) Status() []string                      { return a.status }
func (a *SimpleAction) TakesEffort() Effort                   { return a.effort }
func (a *SimpleAction) HasDuration() TickDuration             { return a.duration }
func (a *SimpleAction) HasUtilityNeed() colonist.NeedType     { return a.utilityNeed }
func (a *SimpleAction) HasUtilityDesire() colonist.DesireType { return a.utilityDesire }
func (a *SimpleAction) SatisfiesNeeds() []SatisfyNeed         { return a.satisfiesNeeds }
func (a *SimpleAction) AgitatesNeeds() []AgitateNeed          { return a.agitatesNeeds }
func (a *SimpleAction) SatisfiesDesires() []SatisfyDesire     { return a.satisfiesDesires }
func (a *SimpleAction) ProducesResources() []ProduceResource  { return a.producesResources }
func (a *SimpleAction) ConsumesResources() []ConsumeResource  { return a.consumesResources }
func (a *SimpleAction) ImprovesSkills() []ImproveSkill        { return a.improvesSkills }
