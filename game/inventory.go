package game

type Inventory struct {
	RawFood    uint `json:raw_food`
	CookedFood uint `json:cooked_food`
	SimpleMeal uint `json:simple_meal`
	Water      uint `json:water`
}
