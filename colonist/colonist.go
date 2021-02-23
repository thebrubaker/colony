package colonist

import (
	"math"
	"math/rand"

	"github.com/rs/xid"
)

const (
	JsonDecimalPlaces float64 = 10000
)

type DesireType string

const (
	Fulfillment DesireType = "Fulfillment"
	Belonging   DesireType = "Belonging"
	Esteem      DesireType = "Esteem"
)

type Desires map[DesireType]float64

func (d Desires) Increase(t DesireType, amount float64) {
	d[t] = math.Min(100, d[t]+amount)
}

func (d Desires) Decrease(t DesireType, amount float64) {
	d[t] = math.Max(0, d[t]-amount)
}

type NeedType string

const (
	Hunger     NeedType = "Hunger"
	Thirst     NeedType = "Thirst"
	Stress     NeedType = "Stress"
	Exhaustion NeedType = "Exhaustion"
	Fear       NeedType = "Fear"
	Pain       NeedType = "Pain"
)

type Needs map[NeedType]float64

func (n Needs) Increase(t NeedType, amount float64) {
	n[t] = math.Min(100, n[t]+amount)
}

func (n Needs) Decrease(t NeedType, amount float64) {
	n[t] = math.Max(0, n[t]-amount)
}

type SkillType string

const (
	Hunting     SkillType = "Hunting"
	Crafting    SkillType = "Crafting"
	Cooking     SkillType = "Cooking"
	Building    SkillType = "Building"
	Gathering   SkillType = "Gathering"
	Mining      SkillType = "Mining"
	Woodcutting SkillType = "Woodcutting"
	Science     SkillType = "Science"
	Combat      SkillType = "Combat"
	Charisma    SkillType = "Charisma"
	Medicine    SkillType = "Medicine"
)

type Skills map[SkillType]float64

func (n Skills) Increase(t SkillType, amount float64) {
	n[t] += amount
}

func (n Skills) Decrease(t SkillType, amount float64) {
	n[t] = math.Max(0, n[t]-amount)
}

// Colonist struct
type Colonist struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Age    uint   `json:"age"`

	Bag       *Bag       `json:"bag"`
	Equipment *Equipment `json:"equipment"`

	Desires Desires `json:"desires"`
	Needs   Needs   `json:"needs"`

	Skills Skills `json:"skills"`
}

// GenerateColonist returns a new colonist with the given name
// and a random set of skills and stats.
func NewColonist(name string) *Colonist {
	return &Colonist{
		Key:  xid.New().String(),
		Name: name,
		Age:  generateAge(),
		Bag: &Bag{
			size: 30,
		},
		Equipment: &Equipment{},
		Desires: map[DesireType]float64{
			Fulfillment: 50,
			Belonging:   50,
			Esteem:      50,
		},
		Needs: map[NeedType]float64{
			Hunger:     60,
			Thirst:     60,
			Stress:     20,
			Exhaustion: 5,
			Fear:       30,
			Pain:       10,
		},
		Skills: generateSkills(),
	}
}

func generateAge() uint {
	return 20 + uint(rand.Float64()*400)
}

func (c *Colonist) SetStatus(s string) {
	c.Status = s
}
