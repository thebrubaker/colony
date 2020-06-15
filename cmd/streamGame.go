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
	"log"

	tm "github.com/buger/goterm"
	"github.com/spf13/cobra"
	"github.com/thebrubaker/colony/client"
	"github.com/thebrubaker/colony/pb"
)

// streamGameCmd represents the stream command
var streamGameCmd = &cobra.Command{
	Use:   "streamGame",
	Short: "Stream the game state over a gRPC connection.",
	Long: `Stream the game state over a gRPC connection. The game state
	will be returned and printed to the console as JSON.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			panic("Missing required argument for game Key")
		}

		key := args[0]
		address := cmd.Flag("address").Value.String()

		client, connection, context, _ := client.CreateClient(address)
		defer connection.Close()

		stream, err := client.StreamGame(context, &pb.StreamGameRequest{
			GameKey: key,
		})
		if err != nil {
			log.Fatal("Failed to open client stream.")
		}

		for {
			gameState, err := stream.Recv()
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			tm.Clear()
			var state map[string]interface{}
			json.Unmarshal([]byte(gameState.Json), &state)
			formatted, err := json.MarshalIndent(state, "", "    ")
			tm.Clear()
			tm.MoveCursor(1, 1)
			tm.Println(string(formatted))
			tm.Flush()
		}
	},
}

func init() {
	rootCmd.AddCommand(streamGameCmd)

	// Here you will define your flags and configuration settings.
	streamGameCmd.PersistentFlags().String("address", "localhost:50051", "The connection address.")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// streamCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// streamCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
