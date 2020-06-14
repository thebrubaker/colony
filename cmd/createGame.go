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
	"log"

	"github.com/spf13/cobra"
	"github.com/thebrubaker/colony/client"
	"github.com/thebrubaker/colony/pb"
)

// createGameCmd represents the createGame command
var createGameCmd = &cobra.Command{
	Use:   "createGame",
	Short: "Creates a new game.",
	Long:  `Creates a new game with the given name. If a game already exists with that name, returns an error.`,
	Run: func(cmd *cobra.Command, args []string) {
		address := cmd.Flag("address").Value.String()

		client, connection, context, _ := client.CreateClient(address)
		defer connection.Close()

		res, err := client.CreateGame(context, &pb.CreateGameRequest{})

		if err != nil {
			log.Fatalf("could not create game: %v", err)
		}

		fmt.Printf("%v", res)
	},
}

func init() {
	rootCmd.AddCommand(createGameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	createGameCmd.Flags().String("address", "localhost:50051", "Address for listening the client")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createGameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
