package interfaces

import "github.com/sashabaranov/go-openai"

type GPTGenerateText interface {
	GenerateText(prompt string) (string, error)
	GenerateStreamText(prompt string) (*openai.ChatCompletionStream, error)
}
