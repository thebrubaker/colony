package colonist

import (
	"encoding/json"
	"testing"
)

func TestJsonMarshalSkills(t *testing.T) {
	skills := &Skills{
		Hunting:     0.1231230123123,
		Crafting:    23.34530123123,
		Cooking:     54.5675671230123123,
		Building:    68.6791230123123,
		Gathering:   35.12431230123123,
		Mining:      14.2351230123123,
		Woodcutting: 85.547230123123,
		Science:     36.253451230123123,
		Combat:      63.23451230123123,
		Charisma:    86.23451230123123,
		Medicine:    31.4357630123123,
	}

	match := `{"building":68.68,"charisma":86.23,"combat":63.23,"cooking":54.57,"crafting":23.35,"gathering":35.12,"hunting":0.12,"medicine":31.44,"mining":14.24,"science":36.25,"woodcutting":85.55}`

	json, err := json.Marshal(skills)

	if err != nil {
		t.Error(err)
	}

	if string(json) != match {
		t.Error("Marshal output of skills does not match", string(json), match)
	}
}
