package game

import "fmt"

type JsonGame struct {
	Tick      string        `json:tick`
	Colonist  JsonColonist  `json:colonist`
	Inventory JsonInventory `json:inventory`
}

func ToJsonGame(gameState *GameState) JsonGame {
	return JsonGame{
		RenderFloat(gameState.Ticker.Count),
		ToJsonColonist(gameState.ActiveColonist),
		ToJsonInventory(gameState.Inventory),
	}
}

type JsonColonist struct {
	Name      string        `json:name`
	Status    string        `json:status`
	Thirst    string        `json:thirst`
	Stress    string        `json:stress`
	Energy    string        `json:energy`
	Hunger    string        `json:hunger`
	Inventory JsonInventory `json:inventory`
}

func ToJsonColonist(colonist *Colonist) JsonColonist {
	return JsonColonist{
		colonist.Name,
		colonist.CurrentAction.Type.Status,
		RenderAttribute(colonist.Thirst),
		RenderAttribute(colonist.Stress),
		RenderAttribute(colonist.Energy),
		RenderAttribute(colonist.Hunger),
		ToJsonInventory(colonist.Inventory),
	}
}

type JsonInventory struct {
	RawFood    uint `json:raw_food`
	CookedFood uint `json:cooked_food`
	SimpleMeal uint `json:simple_meal`
	Water      uint `json:water`
}

func ToJsonInventory(inventory *Inventory) JsonInventory {
	return JsonInventory{
		inventory.RawFood,
		inventory.CookedFood,
		inventory.SimpleMeal,
		inventory.Water,
	}
}

func RenderFloat(number float64) string {
	return fmt.Sprintf("%.2f", number)
}

func RenderAttribute(attribute *Attribute) string {
	return RenderFloat(attribute.Value)
}
