package colonist

// Equippable items can be assigned to the head, weapon or body slot
// of a colonist's equipment.
type Equippable interface {
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
	var existingArmor Equippable

	if c.Equipment.Body != nil {
		existingArmor = c.Equipment.Body
	}

	c.Equipment.Body = e

	return existingArmor, nil
}
