package render

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/thebrubaker/colony/actions"
	"github.com/thebrubaker/colony/colonist"
	"github.com/thebrubaker/colony/game"
	"github.com/willf/pad"
)

const (
	Foreground  = 254
	Background  = -1
	BorderLabel = 254
	BorderLine  = 96
	Highlight   = 31
)

type Controller struct {
	Grid *ui.Grid
}

type Namer interface {
	GetName() string
}

type Stackable interface {
	GetItem() interface{}
	GetCount() uint
}

func Render(g game.GameState) {
	var elements []ui.Drawable

	for i, c := range g.Colonists {
		action := g.Actions[c.Key]
		elements = append(elements, colonistRow(i, c, action)...)
	}

	// var stockpileRows string
	var row string
	colSize := 20

	for index, item := range g.Region.Stockpile.Items {
		name := "Fix Me"
		count := 1
		if i, ok := item.(Stackable); ok {
			item = i.GetItem()
			count = int(i.GetCount())
		}
		if i, ok := item.(Namer); ok {
			name = i.GetName()
		}
		text := pad.Right(fmt.Sprintf("[%v] %s", count, name), colSize, " ")
		row = row + text
		if index+1%3 == 0 {
			row = row + "\n"
		}
	}

	y1 := (len(g.Colonists) * 3)
	y2 := y1 + 12
	x1 := 0
	x2 := colSize * 3
	stockpile := widgets.NewParagraph()
	stockpile.Title = "Stockpile"
	stockpile.SetRect(x1, y1, x2, y2)
	stockpile.Text = row

	elements = append(elements, stockpile)

	ui.Render(elements...)
}

func colonistRow(index int, c *colonist.Colonist, a *actions.Action) []ui.Drawable {
	rowSize := 3
	y1 := 0 + (rowSize * index)
	y2 := rowSize + (rowSize * index)
	colSize := 16
	x1 := 0 * colSize
	x2 := 1 * (colSize * 2)

	actionDuration := a.TickDuration
	actionProgress := a.TickProgress

	name := widgets.NewParagraph()
	if index == 0 {
		name.Title = "Name"
	}
	name.Text = fmt.Sprintf(" %s", c.Name)
	name.SetRect(x1, y1, x2, y2)
	x1 = x2
	x2 = x1 + (colSize * 3)
	status := widgets.NewParagraph()
	if index == 0 {
		status.Title = "Status"
	}
	status.Text = fmt.Sprintf(" %v", c.Status)
	status.SetRect(x1, y1, x2, y2)
	x1 = x2
	x2 = x1 + colSize
	progress := widgets.NewGauge()
	if index == 0 {
		progress.Title = "Progress"
	}
	progress.Percent = int(actionProgress / float64(actionDuration) * 100)
	progress.BarColor = ui.Color(BorderLine)
	progress.SetRect(x1, y1, x2, y2)
	x1 = x2
	x2 = x1 + (colSize / 2)
	bag := widgets.NewParagraph()
	if index == 0 {
		bag.Title = "Bag"
	}
	bag.Text = fmt.Sprintf(" %v/%v", c.Bag.GetItemCount(), c.Bag.Size)
	bag.SetRect(x1, y1, x2, y2)
	x1 = x2
	x2 = x1 + colSize/8*7
	stress := widgets.NewGauge()
	if index == 0 {
		stress.Title = "Stress"
	}
	stress.Percent = int(c.Needs[colonist.Stress])
	stress.BarColor = getColor(stress.Percent)
	stress.SetRect(x1, y1, x2, y2)
	x1 = x2
	x2 = x1 + colSize/8*7
	hunger := widgets.NewGauge()
	if index == 0 {
		hunger.Title = "Hunger"
	}
	hunger.Percent = int(c.Needs[colonist.Hunger])
	hunger.BarColor = getColor(hunger.Percent)
	hunger.SetRect(x1, y1, x2, y2)
	x1 = x2
	x2 = x1 + colSize/8*7
	thirst := widgets.NewGauge()
	if index == 0 {
		thirst.Title = "Thirst"
	}
	thirst.Percent = int(c.Needs[colonist.Thirst])
	thirst.BarColor = getColor(thirst.Percent)
	thirst.SetRect(x1, y1, x2, y2)
	x1 = x2
	x2 = x1 + colSize/8*7
	exhaustion := widgets.NewGauge()
	if index == 0 {
		exhaustion.Title = "Exhaustion"
	}
	exhaustion.Percent = int(c.Needs[colonist.Exhaustion])
	exhaustion.BarColor = getColor(exhaustion.Percent)
	exhaustion.SetRect(x1, y1, x2, y2)

	return []ui.Drawable{
		name,
		status,
		progress,
		bag,
		stress,
		hunger,
		thirst,
		exhaustion,
	}
}

func getColor(count int) ui.Color {
	switch true {
	case count >= 90:
		return ui.Color(130)
	case count >= 75:
		return ui.Color(94)
	case count >= 60:
		return ui.Color(58)
	default:
		return ui.Color(22)
	}
}
