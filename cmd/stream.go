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
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/thebrubaker/colony/pb"
	"google.golang.org/grpc"
)

// streamCmd represents the stream command
var streamCmd = &cobra.Command{
	Use:   "stream",
	Short: "Stream the game state over a gRPC connection.",
	Long: `Stream the game state over a gRPC connection. The game state
	will be returned and printed to the console as JSON.`,
	Run: func(cmd *cobra.Command, args []string) {
		address := "localhost:50051"
		if len(args) > 0 {
			address = args[0]
		}
		// Set up a connection to the server.
		conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewGameServerClient(conn)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		stream, err := c.StreamGameState(ctx, &pb.EmptyRequest{})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}

		for {
			gameState, err := stream.Recv()

			if err != nil {
				log.Fatalf("can not receive %v", err)
			}

			log.Printf("\033c%s\n", gameState.Json)
		}
	},
}

func init() {
	rootCmd.AddCommand(streamCmd)

	// Here you will define your flags and configuration settings.
	streamCmd.PersistentFlags().String("address", "localhost:50051", "The connection address.")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// streamCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// streamCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
