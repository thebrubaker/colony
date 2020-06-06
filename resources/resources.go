package resources

type SimpleResource struct {
	name        string
	description string
	stackable   bool
}

func (w *SimpleResource) GetName() string {
	return w.name
}

func (w *SimpleResource) GetDescription() string {
	return w.description
}

func (w *SimpleResource) IsStackable() bool {
	return w.stackable
}

var Wood *SimpleResource = &SimpleResource{
	name:        "wood",
	description: "a stack of logs",
	stackable:   true,
}

var RawLeather *SimpleResource = &SimpleResource{
	name:        "raw leather",
	description: "raw leather stripped from a hunted animal",
	stackable:   true,
}

var CuredLeather *SimpleResource = &SimpleResource{
	name:        "raw leather",
	description: "prepared leather used in crafting apparel and other goods",
	stackable:   true,
}

type WoodSpear struct {
	quality    float64
	durability float64
	createdBy  string
}

func (w *WoodSpear) GetName() string {
	return "wood spear"
}

func (w *WoodSpear) GetDescription() string {
	return "a sharp wooden spear used for hunting or combat"
}

func (w *WoodSpear) CreatedBy() string {
	return w.createdBy
}

func (w *WoodSpear) GetQuality() float64 {
	return w.quality
}

func (w *WoodSpear) IsStackable() bool {
	return false
}

func (w *WoodSpear) GetDamage() float64 {
	return 5 * w.quality
}

func (w *WoodSpear) GetDamageType() string {
	return "sharp"
}
