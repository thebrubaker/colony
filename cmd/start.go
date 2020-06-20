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
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/thebrubaker/colony/game"
	"github.com/thebrubaker/colony/server"
	"github.com/thebrubaker/colony/streams"
)

// type Debug struct {
// 	Ticker    *ticker.Ticker
// 	Colonists []*colonist.Colonist
// 	Actions   []*actions.ColonistAction
// }

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the game server.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// render, err := cmd.Flags().GetBool("render")

		// if err != nil {
		// 	log.Fatal(err)
		// }
		port := ":50051"
		lis, err := net.Listen("tcp", port)
		log.Printf("Listening on %s", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		sc := streams.NewController()
		gc := game.NewController(sc)
		server := server.NewServer(lis, server.NewGameService(gc, sc))

		gc.CreateGame()

		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		log.Println("\nStopping Server and Game Resources...")
		gc.Stop()
		sc.Stop()
		server.GracefulStop()
		os.Exit(0)
	},
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())

	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	startCmd.PersistentFlags().Bool("render", false, "Streams the game state to the console.")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
