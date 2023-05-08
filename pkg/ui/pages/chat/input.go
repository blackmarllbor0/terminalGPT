package chat

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (c *Chat) initInput() *tview.TextArea {
	c.input = tview.NewTextArea().
		SetPlaceholder("Enter text: ")

	c.controlInput()

	return c.input
}

func (c *Chat) controlInput() {
	c.input.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			prompt := c.input.GetText()
			c.input.SetText("", false)
			c.setUserOutput(prompt)

			response, err := c.gptModelAPI.GenerateText(prompt, c.currentChatID)
			if err != nil {
				c.setOutput(err.Error())
				return event
			}

			c.setChatOutput(response)
		}

		return event
	})
}
