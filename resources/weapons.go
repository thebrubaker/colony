package resources

// DamageType represents the type of damage for a weapon
type DamageType string

// Constants for the different damage types.
const (
	Sharp = "sharp"
	Blunt = "blunt"
)

// QualityLevel represents the quality of the weapon
type QualityLevel uint8

// Constants for the different quality levels that impact
// damage, accuracy and durability
const (
	_ QualityLevel = iota
	Poor
	Decent
	Good
	Great
	Excellent
	Legendary
)

// Weapon is a struct for a resource that can be equipped and is use in combat.
type Weapon struct {
	Name        string
	Description string
	DamageType  DamageType
	Quality     QualityLevel
	Durability  float64
	CreatedBy   string
}

// GetName returns the weapon's name
func (w *Weapon) GetName() string {
	return w.Name
}

// GetDescription returns the weapon's description
func (w *Weapon) GetDescription() string {
	return w.Description
}

// GetCreatedBy returns the name of the colonist who created the weapon
func (w *Weapon) GetCreatedBy() string {
	return w.CreatedBy
}

// GetQuality returns the weapon's quality level
func (w *Weapon) GetQuality() QualityLevel {
	return w.Quality
}

// GetDamage returns the weapon's damage
func (w *Weapon) GetDamage() float64 {
	return 5 * float64(w.Quality)
}

// GetDamageType returns the weapon's damage type
func (w *Weapon) GetDamageType() DamageType {
	return w.DamageType
}
