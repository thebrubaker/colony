package resources

func CreateWoodenSpear(name string, skill float64) *Weapon {
	return &Weapon{
		Name:        "Wooden Spear",
		Description: "A long stick with a sharp point. Good for poking but not very durable or deadly.",
		Quality:     Poor,
		Durability:  100,
	}
}
