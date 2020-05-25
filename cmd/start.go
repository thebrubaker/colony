/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"time"

	wr "github.com/mroth/weightedrand"
	"github.com/spf13/cobra"
)

type Colonist struct {
	Name                   string
	Status                 string
	Thirst                 float64
	Stress                 float64
	LastActionTaken        time.Time
	CurrentAction          *Action
	CurrentActionExpiresAt float64
}

func SubAttribute(attribute float64, amount float64) float64 {
	return math.Max(0, attribute-amount)
}

func AddAttribute(attribute float64, amount float64) float64 {
	return math.Max(0, attribute+amount)
}

type GameState struct {
	ActiveColonist *Colonist
	LastTick       time.Time
	TimeElapsed    time.Duration
	TickElapsed    float64
	Ticks          float64
	Inventory      Stock
}

type Stock map[string]uint

const (
	Survival   uint = 1
	Duty       uint = 2
	Recreation uint = 3
)

type Action struct {
	Status     string
	Priority   uint
	EnergyCost float64
	Duration   float64
	GetUtility func(gameState *GameState, colonist *Colonist) uint
	IsAllowed  func(gameState *GameState, colonist *Colonist) bool
	OnStart    func(gameState *GameState, colonist *Colonist)
	OnTick     func(gameState *GameState, colonist *Colonist)
	OnEnd      func(gameState *GameState, colonist *Colonist)
}

var WatchClouds *Action = &Action{
	Status:     "Watching the clouds.",
	Priority:   Recreation,
	EnergyCost: 0,
	Duration:   15,
	GetUtility: func(gameState *GameState, colonist *Colonist) uint {
		return 40
	},
	IsAllowed: func(gameState *GameState, colonist *Colonist) bool {
		return colonist.Stress > 10
	},
	OnStart: func(gameState *GameState, colonist *Colonist) {},
	OnTick: func(gameState *GameState, colonist *Colonist) {
		colonist.Stress = SubAttribute(colonist.Stress, 4*gameState.TickElapsed)
	},
	OnEnd: func(gameState *GameState, colonist *Colonist) {},
}

var SearchForWater *Action = &Action{
	Status:     "Searching for water.",
	Priority:   Survival,
	EnergyCost: 2,
	Duration:   4,
	GetUtility: func(gameState *GameState, colonist *Colonist) uint {
		return uint(colonist.Thirst / 2)
	},
	IsAllowed: func(gameState *GameState, colonist *Colonist) bool {
		return gameState.Inventory["water"] == 0
	},
	OnStart: func(gameState *GameState, colonist *Colonist) {
	},
	OnTick: func(gameState *GameState, colonist *Colonist) {},
	OnEnd: func(gameState *GameState, colonist *Colonist) {
		if rand.Float64() <= 0.6 {
			gameState.Inventory["water"] += 3
		}
	},
}

var DrinkWater *Action = &Action{
	Status:     "Drinking water.",
	Priority:   Survival,
	EnergyCost: 1,
	Duration:   10,
	GetUtility: func(gameState *GameState, colonist *Colonist) uint {
		if colonist.Thirst >= 60 {
			return 70
		}

		return 0
	},
	IsAllowed: func(gameState *GameState, colonist *Colonist) bool {
		return gameState.Inventory["water"] > 0
	},
	OnStart: func(gameState *GameState, colonist *Colonist) {
		gameState.Inventory["water"]--
	},
	OnTick: func(gameState *GameState, colonist *Colonist) {
		colonist.Thirst = SubAttribute(colonist.Thirst, 5*gameState.TickElapsed)
	},
	OnEnd: func(gameState *GameState, colonist *Colonist) {},
}

var StandIdle *Action = &Action{
	Status:     "Standing Still.",
	Priority:   Recreation,
	EnergyCost: 1,
	Duration:   5,
	GetUtility: func(gameState *GameState, colonist *Colonist) uint {
		return 1
	},
	IsAllowed: func(gameState *GameState, colonist *Colonist) bool {
		return true
	},
	OnStart: func(gameState *GameState, colonist *Colonist) {},
	OnTick:  func(gameState *GameState, colonist *Colonist) {},
	OnEnd:   func(gameState *GameState, colonist *Colonist) {},
}

var WakingUp *Action = &Action{
	Status:     "Waking up from cryosleep.",
	Priority:   Recreation,
	EnergyCost: 1,
	Duration:   4,
	GetUtility: func(gameState *GameState, colonist *Colonist) uint {
		return 1
	},
	IsAllowed: func(gameState *GameState, colonist *Colonist) bool {
		return false
	},
	OnStart: func(gameState *GameState, colonist *Colonist) {},
	OnTick: func(gameState *GameState, colonist *Colonist) {
		colonist.Stress = AddAttribute(colonist.Stress, 15*gameState.TickElapsed)
	},
	OnEnd: func(gameState *GameState, colonist *Colonist) {},
}

func increateThirst(colonist *Colonist, gameState *GameState) {
	colonist.Thirst = AddAttribute(colonist.Thirst, 0.1*gameState.TickElapsed)
}

func processColonistActions(gameState *GameState, colonist *Colonist, actions []*Action) {
	colonist.Stress = AddAttribute(colonist.Stress, colonist.CurrentAction.EnergyCost*gameState.TickElapsed)
	colonist.CurrentAction.OnTick(gameState, colonist)

	if gameState.Ticks < colonist.CurrentActionExpiresAt {
		return
	}

	colonist.CurrentAction.OnEnd(gameState, colonist)

	choices := []wr.Choice{}

	for _, action := range actions {
		if action.IsAllowed(gameState, colonist) {
			choices = append(choices, getChoice(action, action.GetUtility(gameState, colonist)))
		}
	}

	colonist.CurrentAction = wr.NewChooser(choices...).Pick().(*Action)
	colonist.CurrentActionExpiresAt = gameState.Ticks + colonist.CurrentAction.Duration
	colonist.Status = colonist.CurrentAction.Status
	colonist.CurrentAction.OnStart(gameState, colonist)
}

func findRecreation(colonist *Colonist) wr.Choice {
	return wr.Choice{Item: WatchClouds, Weight: 20}
}

func getChoice(action *Action, utility uint) wr.Choice {
	return wr.Choice{Item: action, Weight: utility ^ 3}
}

func updateGameState(gameState *GameState) {
	colonist := gameState.ActiveColonist

	increateThirst(colonist, gameState)

	processColonistActions(gameState, colonist, []*Action{
		SearchForWater,
		WatchClouds,
		StandIdle,
		DrinkWater,
	})
}

func renderGameState(gameState *GameState) {
	colonist := gameState.ActiveColonist
	data := struct {
		Tick      string
		Name      string
		Status    string
		Stress    string
		Thirst    string
		Inventory Stock
	}{
		fmt.Sprintf("%d", uint(gameState.Ticks)),
		colonist.Name,
		colonist.Status,
		fmt.Sprintf("%f", colonist.Stress),
		fmt.Sprintf("%f", colonist.Thirst),
		gameState.Inventory,
	}
	output, _ := json.MarshalIndent(data, "", "    ")
	fmt.Printf("\033c%s\n", string(output))
}

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		gameState := &GameState{
			ActiveColonist: &Colonist{
				Name:                   "Artokun",
				Status:                 WakingUp.Status,
				Thirst:                 60,
				LastActionTaken:        time.Now(),
				CurrentActionExpiresAt: WakingUp.Duration,
				CurrentAction:          WakingUp,
			},
			LastTick:    time.Now(),
			TimeElapsed: 0,
			TickElapsed: 0,
			Ticks:       0,
			Inventory: Stock{
				"water": 0,
			},
		}

		// open a grpc server
		// set a route for an action to "drink some water"
		// push that action on a queue
		// updateGameState: pull actions off the queue
		// process drink water (add 20 to hydration)

		baseTickRate := 16 * time.Millisecond
		// fastTickRate := 8 * time.Millisecond
		// fastestTickRate := 4 * time.Millisecond

		currentTickRate := baseTickRate

		for range time.Tick(currentTickRate) {
			currentTime := time.Now()
			gameState.TimeElapsed = currentTime.Sub(gameState.LastTick)
			updateGameState(gameState)
			renderGameState(gameState)
			if false {
				fmt.Println("You lost the game.")
				break
			}
			gameState.LastTick = currentTime
			gameState.TickElapsed = gameState.TimeElapsed.Seconds() * float64(baseTickRate/currentTickRate)
			gameState.Ticks += gameState.TickElapsed
		}
	},
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())

	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
