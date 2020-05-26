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
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
	"github.com/thebrubaker/colony/game"
)

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
		gameState := game.InitGame()

		baseTickRate := 16 * time.Millisecond
		// fastTickRate := 8 * time.Millisecond
		// fastestTickRate := 4 * time.Millisecond

		currentTickRate := baseTickRate

		for range time.Tick(currentTickRate) {
			currentTime := time.Now()
			gameState.Ticker.Elapsed = currentTime.Sub(gameState.Ticker.LastTick).Seconds() * float64(baseTickRate/currentTickRate)

			gameState.Update()
			gameState.Render()

			if false {
				fmt.Println("You lost the game.")
				break
			}

			gameState.Ticker.LastTick = currentTime
			gameState.Ticker.Count += gameState.Ticker.Elapsed
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
