# How to start the game server

```script
go run main.go serve
```

# How to stream the game server

```script
go run main.go stream
```

# How to add an action

Create a new Struct in `game/actions.go` for your action. Then at the bottom of the file, append your
action to `ActionTypes []*ActionType`

```go
var WatchClouds *ActionType = &ActionType{
	ID:         "3", // unique string ID
	Status:     "Drinking water.", // make it interesting
	Priority:   Recreation, // caps the utility to this category
	EnergyCost: 0, // per tick
  Duration:   9, // how many ticks
  // Determine if this is one of the actions allowed to the colonist
  // when the colonist is deciding what to do next. Is this action
  // physically possible for the given game state?
  IsAllowed: func(actionType *ActionType, colonist *Colonist) bool {
		if colonist.GameState.Inventory.Water > 0 {
      return true
    }

    if colonist.Inventory.Water > 0 {
      return true
    }

    return false
	},
  // Return a score for the utility this action provides to the colonist.
  // Typically scales with the related need.
  // Ease functions provide more interesting AI curves
  // Visual Here: https://github.com/fogleman/ease
	GetUtility: func(actionType *ActionType, colonist *Colonist) float64 {
    // |                      *
    // |                     *
    // |                    *
    // |                   *
    // |                 **
    // |               **
    // |         ******
    // |*********
    // |-------------------------
		return ease.InQuart(colonist.Thirst.Value / 100)
  },
  // When this action is selected, you can modify the game state such
  // as equiping an item from the global inventory to the colonist.
  OnStart: func(action *Action, ticker *Ticker) {
		action.Game.Inventory.Water--
		action.Colonist.Inventory.Water++
	},
  // While taking this action, what do you want to change about
  // the state per tick? Make sure you reduce it by the fraction of
  // the tick that has passed this compute cycle.
	OnTick: func(action *Action, ticker *Ticker) {
		action.Colonist.Thirst.Sub(5 * ticker.Elapsed)
	},
  // When the action comes to an end, you might want to reduce
  // some inventory, or conclude the action by having a result
  OnEnd: func(action *Action, ticker *Ticker) {
		action.Colonist.Inventory.Water--
	},
}
```
