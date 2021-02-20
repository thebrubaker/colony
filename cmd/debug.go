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
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/thebrubaker/colony/debug"
	"github.com/thebrubaker/colony/game"
)

// type Debug struct {
// 	Ticker    *ticker.Ticker
// 	Colonists []*colonist.Colonist
// 	Actions   []*actions.ColonistAction
// }

// startCmd represents the start command
var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "Debug the game server.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		sc := debug.NewController()
		gc := game.NewController(sc)

		gameKey := gc.CreateGame()
		sc.CreateStream(gameKey)

		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		log.Println("\nStopping Game Resources...")
		gc.Stop()
		sc.Stop()
		os.Exit(0)
	},
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())

	rootCmd.AddCommand(debugCmd)
}
