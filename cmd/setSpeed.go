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
	"strconv"

	"github.com/spf13/cobra"
	"github.com/thebrubaker/colony/client"
	"github.com/thebrubaker/colony/pb"
)

// setSpeedCmd represents the setSpeed command
var setSpeedCmd = &cobra.Command{
	Use:   "setSpeed",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		address := cmd.Flag("address").Value.String()

		if len(args) == 0 {
			panic("Missing required argument for game Key")
		}

		key := args[0]

		if len(args) == 1 {
			panic("Missing required argument for speed")
		}

		speed, err := strconv.ParseInt(args[1], 10, 64)

		if err != nil {
			log.Fatal("cannot parse second argument as speed type", err)
		}

		client, connection, context, _ := client.CreateClient(address)
		defer connection.Close()

		res, err := client.SetSpeed(context, &pb.SetSpeedRequest{GameKey: key, Speed: pb.SetSpeedRequest_Speed(speed)})

		if err != nil {
			log.Fatalf("could not create game: %v", err)
		}

		fmt.Printf("%v", res)
	},
}

func init() {
	rootCmd.AddCommand(setSpeedCmd)

	// Here you will define your flags and configuration settings.
	setSpeedCmd.PersistentFlags().String("address", "localhost:50051", "The connection address.")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setSpeedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setSpeedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
