package colonist

import (
	"math/rand"

	"github.com/thebrubaker/colony/stackable"
)

// Colonist struct
type Colonist struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Age    uint   `json:"age"`

	Bag       []interface{} `json:"bag"`
	Equipment *Equipment    `json:"equipment"`

	Needs  *Needs  `json:"needs"`
	Skills *Skills `json:"skills"`
	Stats  *Stats  `json:"stats"`
}

// AddToBag adds an element or a stackable element to the colonist's bag.
func (c *Colonist) AddToBag(elem interface{}) error {
	bag, err := stackable.Put(c.Bag, elem)

	if err != nil {
		return err
	}

	c.Bag = bag

	return nil
}

// RemoveFromBag removes an element or a stackable element from the colonist's bag.
func (c *Colonist) RemoveFromBag(elem interface{}) error {
	bag, err := stackable.Take(c.Bag, elem)

	if err != nil {
		return err
	}

	c.Bag = bag

	return nil
}

// Has determines if the colonist has the given element or stackable in their hands or bag.
func (c *Colonist) Has(elem interface{}) bool {
	return stackable.Has(c.Bag, elem)
}

// GenerateColonist returns a new colonist with the given name
// and a random set of skills and stats.
func GenerateColonist(name string) *Colonist {
	return &Colonist{
		Name: name,
		Age:  generateAge(),
		Needs: &Needs{
			Thirst:     30,
			Stress:     30,
			Exhaustion: 30,
			Hunger:     30,
		},
		Skills:    generateSkills(),
		Stats:     generateStats(),
		Equipment: &Equipment{},
	}
}

func generateAge() uint {
	return 20 + uint(rand.Float64()*400)
}

func generateSkills() *Skills {
	var values [11]float64

	for i := 0; i < 2; i++ {
		values[i] = 4 + (rand.Float64() * 10)
	}

	for i := 2; i < 8; i++ {
		values[i] = 3 + (rand.Float64() * 5)
	}

	for i := 8; i < 11; i++ {
		values[i] = rand.Float64() * 2
	}

	rand.Shuffle(len(values), func(i, j int) {
		values[i], values[j] = values[j], values[i]
	})

	return &Skills{
		Hunting:     values[0],
		Crafting:    values[1],
		Cooking:     values[2],
		Building:    values[3],
		Gathering:   values[4],
		Mining:      values[5],
		Woodcutting: values[6],
		Science:     values[7],
		Combat:      values[8],
		Charisma:    values[9],
		Medicine:    values[10],
	}
}

func generateStats() *Stats {
	return &Stats{
		Strength:  8 + rand.Float64()*10,
		Dexterity: 8 + rand.Float64()*10,
	}
}

func (c *Colonist) SetStatus(s string) {
	c.Status = s
}

type Skills struct {
	Hunting     float64 `json:"hunting"`
	Crafting    float64 `json:"crafting"`
	Cooking     float64 `json:"cooking"`
	Building    float64 `json:"building"`
	Gathering   float64 `json:"gathering"`
	Mining      float64 `json:"mining"`
	Woodcutting float64 `json:"woodcutting"`
	Science     float64 `json:"science"`
	Combat      float64 `json:"combat"`
	Charisma    float64 `json:"charisma"`
	Medicine    float64 `json:"medicine"`
}

type Stats struct {
	Strength  float64 `json:"strength"`
	Dexterity float64 `json:"dexterity"`
}
