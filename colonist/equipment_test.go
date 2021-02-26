package colonist

import (
	"testing"
)

type Helmet struct {
	Name     string
	Material string
}

func (h Helmet) StrengthRequirement() float64 {
	return 0
}

func (h Helmet) DexterityRequirement() float64 {
	return 0
}

func TestColonistEquipment(t *testing.T) {
	colonist := NewColonist()

	ironHelmet := &Helmet{
		"Iron Helmet",
		"Iron",
	}

	old, err := colonist.EquipHelmet(ironHelmet)

	if err != nil {
		t.Error(err)
	}

	if old != nil {
		t.Error("failed to return nil old helmet")
	}

	if colonist.Equipment.Head != ironHelmet {
		t.Error("failed to equip helmet")
	}
}
