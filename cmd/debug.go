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
	"math/rand"
	"os"
	"time"

	tm "github.com/buger/goterm"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/spf13/cobra"
	"github.com/thebrubaker/colony/colonist"
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
		if err := ui.Init(); err != nil {
			log.Fatalf("failed to initialize termui: %v", err)
		}
		defer ui.Close()

		gc := game.NewController()
		defer gc.Stop()

		key := gc.CreateGame()

		quitc := make(chan struct{})
		defer close(quitc)

		tm.Clear()
		tm.MoveCursor(1, 1)
		tm.Flush()

		go func(quitc chan struct{}) {
			for {
				select {
				case e := <-ui.PollEvents():
					switch e.ID {
					case "q", "<C-c>":
						close(quitc)
						return
					}
				case <-time.Tick(100 * time.Millisecond):
					tick := widgets.NewParagraph()
					tick.Title = "Tick Count"
					tick.Text = fmt.Sprintf(" %v", int64(gc.Render(key).Ticker.Count))
					tick.SetRect(0, 0, 30, 3)

					c := gc.Render(key).Colonists[0]
					name := widgets.NewParagraph()
					name.Title = "Name"
					name.Text = fmt.Sprintf(" %v", c.Name)
					name.SetRect(32, 0, 62, 3)

					action := widgets.NewParagraph()
					action.Title = "Action"
					action.Text = fmt.Sprintf(" %v", c.Status)
					action.SetRect(0, 3, 62, 6)

					stress := widgets.NewGauge()
					stress.Title = "Stress"
					stress.SetRect(64, 0, 94, 3)
					stress.Percent = int(c.Needs[colonist.Stress])

					hunger := widgets.NewGauge()
					hunger.Title = "Hunger"
					hunger.SetRect(64, 3, 94, 6)
					hunger.Percent = int(c.Needs[colonist.Hunger])

					thirst := widgets.NewGauge()
					thirst.Title = "Thirst"
					thirst.SetRect(64, 6, 94, 9)
					thirst.Percent = int(c.Needs[colonist.Thirst])

					exhaustion := widgets.NewGauge()
					exhaustion.Title = "Exhaustion"
					exhaustion.SetRect(64, 9, 94, 12)
					exhaustion.Percent = int(c.Needs[colonist.Exhaustion])

					ui.Render(tick, name, action, stress, hunger, thirst, exhaustion)
				}
			}
		}(quitc)

		<-quitc
		tm.Clear()
		tm.MoveCursor(1, 1)
		tm.Flush()
		tm.Println("Stopping Game Resources...")
		os.Exit(0)
	},
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())

	rootCmd.AddCommand(debugCmd)
}
