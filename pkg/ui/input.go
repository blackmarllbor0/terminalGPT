package ui

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (ui *UserInterface) initInput() *tview.TextArea {
	ui.input = tview.NewTextArea().
		SetLabel("Enter text: ")

	ui.controlInput()

	return ui.input
}

func (ui *UserInterface) controlInput() {
	ui.input.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			promt := ui.input.GetText()
			ui.input.SetText("", false)
			ui.setOutput(fmt.Sprintf("You: %v", promt))

			responce, err := ui.gptModelApi.GenerateText(promt)
			if err != nil {
				ui.setOutput(err.Error())
				return event
			}

			ui.setOutput(fmt.Sprintf("GPTModel: %v", responce))
		case tcell.KeyEscape:
			ui.stop()

			fmt.Println("Goodbye! 👋")
		}

		return event
	})
}
