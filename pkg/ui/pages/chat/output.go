package chat

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/sashabaranov/go-openai"
)

func (c *Chat) initOutput() *tview.TextView {
	c.output = tview.NewTextView().
		SetTextColor(tcell.ColorWhiteSmoke).
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetTextAlign(tview.AlignLeft).
		SetScrollable(true)

	content, err := c.gptModelAPI.GetAllChats()
	if err != nil {
		c.output.SetText(fmt.Sprintf("err: %v", err))
	}

	if len(content) == 0 {
		ID, err := c.gptModelAPI.CreateNewChat()
		if err != nil {
			c.output.SetText(fmt.Sprintf("err: %v", err))
		}

		c.currentChatID = ID
	} else {
		c.currentChatID = content[0].ID
		if err := c.gptModelAPI.UpdateChatMessages(c.currentChatID); err != nil {
			c.output.SetText(err.Error())
		}

		c.writeContent(content[0].Content)
	}

	return c.output
}

func (c *Chat) setOutput(text string) {
	c.output.SetText(fmt.Sprintf("%s\n\n%s", c.output.GetText(false), text))
}

func (c *Chat) setUserOutput(text string) {
	c.output.SetText(fmt.Sprintf("%s\n\n%s", c.output.GetText(false), fmt.Sprintf("🧑‍💻: %v", text)))
}

func (c *Chat) setChatOutput(text string) {
	c.output.SetText(fmt.Sprintf("%s\n\n%s", c.output.GetText(false), fmt.Sprintf("🤖: %v", text)))
}

func (c *Chat) writeContent(content []openai.ChatCompletionMessage) {
	for _, message := range content {
		if message.Role == openai.ChatMessageRoleUser {
			c.setUserOutput(message.Content)
		} else if message.Role == openai.ChatMessageRoleAssistant {
			c.setChatOutput(message.Content)
		}
	}
}
