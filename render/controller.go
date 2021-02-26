package render

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/thebrubaker/colony/colonist"
	"github.com/thebrubaker/colony/game"
)

type Controller struct {
	Grid *ui.Grid
}

func Render(g game.GameState) {
	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	var rows []interface{}

	for _, c := range g.Colonists {
		rows = append(rows, ui.NewRow(1.0/10,
			colonistRow(c)...,
		))
	}

	grid.Set(rows...)
	ui.Render(grid)
}

func colonistRow(c *colonist.Colonist) []interface{} {
	w1 := widgets.NewParagraph()
	w1.Title = "Name"
	w1.Text = fmt.Sprintf(" %s", c.Name)
	w2 := widgets.NewParagraph()
	w2.Title = "Status"
	w2.Text = c.Status
	w3 := widgets.NewParagraph()
	w3.Title = "Progress"
	w4 := widgets.NewGauge()
	w4.Title = "Stress"
	w4.Percent = int(c.Needs[colonist.Stress])
	w5 := widgets.NewGauge()
	w5.Title = "Hunger"
	w5.Percent = int(c.Needs[colonist.Hunger])
	w6 := widgets.NewGauge()
	w6.Title = "Thirst"
	w6.Percent = int(c.Needs[colonist.Thirst])
	w7 := widgets.NewGauge()
	w7.Title = "Exhaustion"
	w7.Percent = int(c.Needs[colonist.Exhaustion])

	width := 10.0

	return []interface{}{
		ui.NewCol(2.0/width, w1),
		ui.NewCol(3.0/width, w2),
		ui.NewCol(1.0/width, w3),
		ui.NewCol(1.0/width, w4),
		ui.NewCol(1.0/width, w5),
		ui.NewCol(1.0/width, w6),
		ui.NewCol(1.0/width, w7),
	}
}
