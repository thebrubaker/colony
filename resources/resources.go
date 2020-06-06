package resources

type SimpleResource struct {
	Name        string
	Description string
	stackable   bool
}

func (w *SimpleResource) GetName() string {
	return w.Name
}

func (w *SimpleResource) GetDescription() string {
	return w.Description
}

func (w *SimpleResource) IsStackable() bool {
	return w.stackable
}

var Wood *SimpleResource = &SimpleResource{
	Name:        "wood",
	Description: "a stack of logs",
	stackable:   true,
}

var RawLeather *SimpleResource = &SimpleResource{
	Name:        "raw leather",
	Description: "raw leather stripped from a hunted animal",
	stackable:   true,
}

var CuredLeather *SimpleResource = &SimpleResource{
	Name:        "raw leather",
	Description: "prepared leather used in crafting apparel and other goods",
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
