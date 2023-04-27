package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (ui *UserInterface) initInput() *tview.TextArea {
	ui.input = tview.NewTextArea().
		SetLabel("Enter text: ").
		SetPlaceholder("Hello, what is your name?")

	ui.input.SetInputCapture(ui.controlInput)

	return ui.input
}

func (ui *UserInterface) controlInput(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyEnter:
		text := ui.input.GetText()
		ui.setOutput(text)

		ui.input.SetText("", false)
	case tcell.KeyEscape:
		ui.stop()

		fmt.Println("До свидания! 👋")
	}

	return event
}
