package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (ui *UserInterface) initOutput() *tview.TextView {
	ui.output = tview.NewTextView().
		SetTextColor(tcell.ColorWhiteSmoke).
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetTextAlign(2).
		SetText("Code will be displayed here")

	return ui.output
}

func (ui *UserInterface) setOutput(text string) {
	ui.output.SetText(text)
}
