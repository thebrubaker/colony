package colonist

import (
	"encoding/json"
	"math"
	"math/rand"
)

func generateStats() *Stats {
	return &Stats{
		Strength:  8 + rand.Float64()*10,
		Dexterity: 8 + rand.Float64()*10,
	}
}

type Stats struct {
	Strength  float64 `json:"strength"`
	Dexterity float64 `json:"dexterity"`
}

func (s Stats) MarshalJSON() ([]byte, error) {
	data := map[string]float64{
		"dexterity": math.Round(s.Dexterity*JsonDecimalPlaces) / JsonDecimalPlaces,
		"strength":  math.Round(s.Strength*JsonDecimalPlaces) / JsonDecimalPlaces,
	}

	return json.Marshal(data)
}
