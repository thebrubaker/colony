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
	"math/rand"
	"os"
	"time"

	tm "github.com/buger/goterm"
	"github.com/spf13/cobra"
	"github.com/thebrubaker/colony/actions"
	"github.com/thebrubaker/colony/colonist"
	"github.com/thebrubaker/colony/region"
	"github.com/thebrubaker/colony/server"
	"github.com/thebrubaker/colony/ticker"
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
		renderToConsole := cmd.Flag("render").Value.String() == "true"

		t := ticker.CreateTick()

		r := &region.Region{}

		c := []*colonist.Colonist{
			colonist.GenerateColonist("Joel"),
		}

		a := actions.InitActions(r, c)

		// debug := &Debug{
		// 	Ticker:    t,
		// 	Colonists: c,
		// 	Actions:   a,
		// }

		go t.OnTick(func(t *ticker.Ticker) {
			actions.Update(t.TickElapsed, r, c, a)
		})

		if renderToConsole {
			tm.Clear()

			ticker.OnTickMilliseconds(16, func() {
				json, err := json.MarshalIndent(a[0], "", "    ")

				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				output := string(json)
				tm.Clear()
				tm.MoveCursor(1, 1)
				tm.Println(string(output))
				tm.Flush()
			})
		} else {
			server.StartServer()
		}
	},
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())

	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	startCmd.PersistentFlags().String("render", "false", "Streams the game state to the console.")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
