package colonist

import (
	"errors"
)

// Equippable items can be assigned to the head, weapon or body slot
// of a colonist's equipment.
type Equippable interface {
	StrengthRequirement() float64
	DexterityRequirement() float64
}

// Equipment for the Colonist
type Equipment struct {
	Head   Equippable `json:"head"`
	Weapon Equippable `json:"weapon"`
	Body   Equippable `json:"body"`
}

// EquipHelmet equips the helmet if the requirements are met. Returns
// the currently equipped helmet or nil if no helmet is equipped.
func (c *Colonist) EquipHelmet(e Equippable) (Equippable, error) {
	s := e.StrengthRequirement()

	if s != 0 && s > c.Stats.Strength {
		return nil, errors.New("not strong enough to equip helmet")
	}

	d := e.DexterityRequirement()

	if d != 0 && d > c.Stats.Dexterity {
		return nil, errors.New("not enough dexterity to equip helmet")
	}

	var existingHelmet Equippable

	if c.Equipment.Head != nil {
		existingHelmet = c.Equipment.Head
	}

	c.Equipment.Head = e

	return existingHelmet, nil
}

// EquipWeapon equips the weapon if the requirements are met. Returns
// the currently equipped weapon or nil if no weapon is equipped.
func (c *Colonist) EquipWeapon(e Equippable) (Equippable, error) {
	s := e.StrengthRequirement()

	if s != 0 && s > c.Stats.Strength {
		return nil, errors.New("not strong enough to equip weapon")
	}

	d := e.DexterityRequirement()

	if d != 0 && d > c.Stats.Dexterity {
		return nil, errors.New("not enough dexterity to equip weapon")
	}

	var existingWeapon Equippable

	if c.Equipment.Weapon != nil {
		existingWeapon = c.Equipment.Weapon
	}

	c.Equipment.Weapon = e

	return existingWeapon, nil
}

// EquipArmor equips the armor if the requirements are met. Returns
// the currently equipped armor or nil if no armor is equipped.
func (c *Colonist) EquipArmor(e Equippable) (Equippable, error) {
	s := e.StrengthRequirement()

	if s != 0 && s > c.Stats.Strength {
		return nil, errors.New("not strong enough to equip armor")
	}

	d := e.DexterityRequirement()

	if d != 0 && d > c.Stats.Dexterity {
		return nil, errors.New("not enough dexterity to equip armor")
	}

	var existingArmor Equippable

	if c.Equipment.Body != nil {
		existingArmor = c.Equipment.Body
	}

	c.Equipment.Body = e

	return existingArmor, nil
}
