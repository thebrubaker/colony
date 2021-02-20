package colonist

import (
	"math/rand"

	"github.com/rs/xid"
	"github.com/thebrubaker/colony/need"
)

const (
	JsonDecimalPlaces float64 = 10000
)

// Colonist struct
type Colonist struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Age    uint   `json:"age"`

	Bag       *Bag       `json:"bag"`
	Equipment *Equipment `json:"equipment"`

	Needs  *need.Needs `json:"needs"`
	Skills *Skills     `json:"skills"`
	Stats  *Stats      `json:"stats"`
}

type ColonistKey string

// GenerateColonist returns a new colonist with the given name
// and a random set of skills and stats.
func NewColonist(name string) *Colonist {
	return &Colonist{
		Key:       newColonistKey(),
		Name:      name,
		Age:       generateAge(),
		Needs:     need.NewNeeds(),
		Skills:    generateSkills(),
		Stats:     generateStats(),
		Equipment: &Equipment{},
		Bag:       &Bag{},
	}
}

func generateAge() uint {
	return 20 + uint(rand.Float64()*400)
}

func (c *Colonist) SetStatus(s string) {
	c.Status = s
}

func newColonistKey() string {
	return xid.New().String()
}
