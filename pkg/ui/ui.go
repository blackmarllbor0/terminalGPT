package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"terminalGPT/internal/pkg/api/interfaces"
)

type UserInterface struct {
	app    *tview.Application
	input  *tview.TextArea
	output *tview.TextView

	gptModelApi interfaces.GPTGenerateText
}

func NewUserInterface(gptModelApi interfaces.GPTGenerateText) *UserInterface {
	return &UserInterface{gptModelApi: gptModelApi}
}

func (ui *UserInterface) Start() error {
	ui.app = tview.NewApplication()

	grid := tview.NewGrid().SetRows(1, 0).SetColumns(0).SetBorders(true)

	grid.AddItem(ui.info(), 0, 0, 1, 1, 0, 0, false)
	grid.AddItem(ui.initOutput(), 1, 0, 1, 1, 0, 0, false)
	grid.AddItem(ui.initInput(), 2, 0, 1, 1, 0, 0, true)

	if err := ui.app.SetRoot(grid, true).Run(); err != nil {
		return err
	}

	return nil
}

func (ui *UserInterface) stop() {
	ui.app.Stop()
}

func (ui *UserInterface) info() *tview.TextView {
	return tview.NewTextView().
		SetTextColor(tcell.ColorWhiteSmoke).
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetTextAlign(2).
		SetText("Instructions: Press Enter to submit")
}
