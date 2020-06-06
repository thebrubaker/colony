package colonist

import (
	"encoding/json"
	"math"
	"math/rand"
)

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

func (s Skills) MarshalJSON() ([]byte, error) {
	data := map[string]float64{
		"hunting":     math.Round(s.Hunting*JsonDecimalPlaces) / JsonDecimalPlaces,
		"crafting":    math.Round(s.Crafting*JsonDecimalPlaces) / JsonDecimalPlaces,
		"cooking":     math.Round(s.Cooking*JsonDecimalPlaces) / JsonDecimalPlaces,
		"building":    math.Round(s.Building*JsonDecimalPlaces) / JsonDecimalPlaces,
		"gathering":   math.Round(s.Gathering*JsonDecimalPlaces) / JsonDecimalPlaces,
		"mining":      math.Round(s.Mining*JsonDecimalPlaces) / JsonDecimalPlaces,
		"woodcutting": math.Round(s.Woodcutting*JsonDecimalPlaces) / JsonDecimalPlaces,
		"science":     math.Round(s.Science*JsonDecimalPlaces) / JsonDecimalPlaces,
		"combat":      math.Round(s.Combat*JsonDecimalPlaces) / JsonDecimalPlaces,
		"charisma":    math.Round(s.Charisma*JsonDecimalPlaces) / JsonDecimalPlaces,
		"medicine":    math.Round(s.Medicine*JsonDecimalPlaces) / JsonDecimalPlaces,
	}

	return json.Marshal(data)
}
