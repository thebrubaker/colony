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
				if e.Type == ui.KeyboardEvent {
					close(quitc)
					return
				}
				case <-time.Tick(33 * time.Millisecond):
					p := widgets.NewParagraph()
					p.Text = fmt.Sprintf("%f", gc.Render(key).Ticker.Count)
					p.SetRect(0, 0, 25, 5)
					ui.Render(p)
					// data, err := json.MarshalIndent(gc.Render(key), "", "    ")

					// if err != nil {
					// 	close(quitc)
					// 	return
					// }

					// tm.Clear()
					// tm.MoveCursor(1, 1)
					// tm.Println(string(data))
					// tm.Flush()
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
