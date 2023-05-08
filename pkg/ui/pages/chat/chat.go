package chat

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"terminalGPT/internal/pkg/api/interfaces"
)

type Chat struct {
	input         *tview.TextArea
	output        *tview.TextView
	currentChatID primitive.ObjectID

	gptModelAPI interfaces.GPTGenerateText
}

func NewChat(gptModelAPI interfaces.GPTGenerateText) *Chat {
	return &Chat{gptModelAPI: gptModelAPI}
}

func (c *Chat) Page() *tview.Grid {
	grid := tview.NewGrid().SetRows(1, 0).SetColumns(0).SetBorders(true)

	grid.AddItem(c.info(), 0, 0, 1, 1, 0, 0, false)
	grid.AddItem(c.initOutput(), 1, 0, 1, 1, 0, 0, false)
	grid.AddItem(c.initInput(), 2, 0, 1, 1, 0, 0, true)

	return grid
}

func (*Chat) info() *tview.TextView {
	return tview.NewTextView().
		SetTextColor(tcell.ColorWhiteSmoke).
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetTextAlign(2).
		SetText("Instructions: Press Enter to submit")
}
