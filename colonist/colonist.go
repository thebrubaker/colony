package colonist

import (
	"math/rand"

	"github.com/thebrubaker/colony/need"
)

const (
	JsonDecimalPlaces float64 = 100
)

// Colonist struct
type Colonist struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Age    uint   `json:"age"`

	Bag       *Bag       `json:"bag"`
	Equipment *Equipment `json:"equipment"`

	Needs  *need.Needs `json:"needs"`
	Skills *Skills     `json:"skills"`
	Stats  *Stats      `json:"stats"`
}

// GenerateColonist returns a new colonist with the given name
// and a random set of skills and stats.
func GenerateColonist(name string) *Colonist {
	return &Colonist{
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
