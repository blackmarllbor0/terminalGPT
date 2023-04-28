package GPT3

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"terminalGPT/config/interfaces"
)

type GPT3 struct {
	client *openai.Client
	ctx    context.Context

	configReader interfaces.IConfigReader
}

func NewGPT3(configReader interfaces.IConfigReader) *GPT3 {
	apiKey := configReader.GetString("api-key")
	return &GPT3{
		client:       openai.NewClient(apiKey),
		ctx:          context.Background(),
		configReader: configReader,
	}
}

func (g *GPT3) GenerateText(prompt string) (string, error) {
	responce, err := g.client.CreateChatCompletion(g.ctx, openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{{
			Role:    openai.ChatMessageRoleUser,
			Content: prompt,
		}},
	})

	if err != nil {
		return "", err
	}

	return responce.Choices[0].Message.Content, nil
}
