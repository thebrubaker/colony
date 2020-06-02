// package game

// import (
// 	"fmt"
// )

// type JsonGame struct {
// 	Tick      string         `json:tick`
// 	Colonists []JsonColonist `json:colonist`
// 	Stockpile JsonInventory  `json:inventory`
// }

// func ToJsonGame(gameState *GameState) JsonGame {
// 	var colonists []JsonColonist

// 	for _, v := range gameState.Colonists {
// 		colonists = append(colonists, ToJsonColonist(v))
// 	}

// 	return JsonGame{
// 		RenderFloat(gameState.Ticker.Count),
// 		colonists,
// 		ToJsonInventory(gameState.Stockpile),
// 	}
// }

// func GetStatuses(colonists []*Colonist) map[string]string {
// 	statuses := make(map[string]string)
// 	for _, v := range colonists {
// 		statuses[v.Name] = v.CurrentAction.Type.Status
// 	}
// 	return statuses
// }

// func RenderFloat(number float64) string {
// 	return fmt.Sprintf("%.2f", number)
// }

// func RenderAttribute(attribute *Attribute) string {
// 	bar := "["

// 	filled := int(attribute.Value / 5)
// 	empty := 20 - filled

// 	for i := 0; i < filled; i++ {
// 		bar = bar + "o"
// 	}

// 	for i := 0; i < empty; i++ {
// 		bar = bar + "-"
// 	}

// 	bar = bar + "]"

// 	return bar
// }
