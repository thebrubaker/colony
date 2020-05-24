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
	Name           string
	Status         string
	Alive          bool
	Hydration      float64
	LastActionTake time.Time
}

type GameState struct {
	ActiveColonist *Colonist
	LastTick       time.Time
	Elapsed        time.Duration
}

type Action struct {
	Status string
}

func gameIsLost(gameState *GameState) bool {
	return gameState.ActiveColonist.Alive == false
}

func searchForWater(colonist *Colonist) wr.Choice {
	action := &Action{Status: "Searching for water."}

	if colonist.Hydration <= 4 {
		return wr.Choice{Item: action, Weight: 70}
	}

	if colonist.Hydration <= 3 {
		return wr.Choice{Item: action, Weight: 80}
	}

	if colonist.Hydration <= 2 {
		return wr.Choice{Item: action, Weight: 90}
	}

	return wr.Choice{Item: action, Weight: 5}
}

func decreaseHydration(colonist *Colonist, gameState *GameState) {
	ratePerSecond := 0.1

	decrement := colonist.Hydration - float64(ratePerSecond)*gameState.Elapsed.Seconds()

	colonist.Hydration = math.Round(decrement*1000) / 1000
}

func isColonistAlive(colonist *Colonist) bool {
	if colonist.Hydration <= 0 {
		return false
	}

	return true
}

func processColonistActions(c wr.Chooser) *Action {
	return c.Pick().(*Action)
}

func findRecreation(colonist *Colonist) wr.Choice {
	action := &Action{Status: "Taking a nice walk."}

	return wr.Choice{Item: action, Weight: 40}
}

func updateGameState(gameState *GameState) {
	colonist := gameState.ActiveColonist

	decreaseHydration(colonist, gameState)

	actionToTake := processColonistActions(wr.NewChooser(
		searchForWater(colonist),
		findRecreation(colonist),
	))

	if colonist.LastActionTake.Sub(time.Now())*time.Second >= 6 {
		colonist.Status = actionToTake.Status
		colonist.LastActionTake = time.Now()
	}

	colonist.Alive = isColonistAlive(colonist)
}

func renderGameState(gameState *GameState) {
	output, _ := json.MarshalIndent(gameState, "", "    ")
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
				Name:           "Artokun",
				Status:         "Waking Up",
				Alive:          true,
				Hydration:      10,
				LastActionTake: time.Now(),
			},
			LastTick: time.Now(),
			Elapsed:  0,
		}

		// open a grpc server
		// set a route for an action to "drink some water"
		// push that action on a queue
		// updateGameState: pull actions off the queue
		// process drink water (add 20 to hydration)

		for range time.Tick(16 * time.Millisecond) {
			currentTime := time.Now()
			gameState.Elapsed = currentTime.Sub(gameState.LastTick)
			updateGameState(gameState)
			renderGameState(gameState)
			if gameIsLost(gameState) {
				fmt.Println("You lost the game.")
				break
			}
			gameState.LastTick = currentTime
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
