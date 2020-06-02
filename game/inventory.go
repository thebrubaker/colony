// package game

// type Inventory struct {
// 	RawFood      uint `json:raw_food`
// 	CookedFood   uint `json:cooked_food`
// 	SimpleMeal   uint `json:simple_meal`
// 	Water        uint `json:water`
// 	Wood         uint `json:wood`
// 	HuntingSpear uint `json:hunting_spear`
// 	CookingFire  uint `json:cooking_fire`
// }

// type JsonInventory struct {
// 	RawFood      uint `json:raw_food`
// 	CookedFood   uint `json:cooked_food`
// 	SimpleMeal   uint `json:simple_meal`
// 	Water        uint `json:water`
// 	Wood         uint `json:wood`
// 	HuntingSpear uint `json:hunting_spear`
// 	CookingFire  uint `json:cooking_fire`
// }

// func ToJsonInventory(inventory *Inventory) JsonInventory {
// 	return JsonInventory{
// 		inventory.RawFood,
// 		inventory.CookedFood,
// 		inventory.SimpleMeal,
// 		inventory.Water,
// 		inventory.Wood,
// 		inventory.HuntingSpear,
// 		inventory.CookingFire,
// 	}
// }
