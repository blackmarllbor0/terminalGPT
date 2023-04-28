package ui

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (ui *UserInterface) initOutput() *tview.TextView {
	ui.output = tview.NewTextView().
		SetTextColor(tcell.ColorWhiteSmoke).
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetTextAlign(tview.AlignLeft).
		SetScrollable(true)
	//SetText("Code will be displayed here")

	return ui.output
}

func (ui *UserInterface) setOutput(text string) {
	ui.output.SetText(fmt.Sprintf("%s\n\n%s", ui.output.GetText(false), text))
}
