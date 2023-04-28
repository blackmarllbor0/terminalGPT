package GPT3dot5Turbo

import (
	"context"
	openai "github.com/sashabaranov/go-openai"
	"terminalGPT/config/interfaces"
)

type GPT3dot5Turbo struct {
	client *openai.Client
	ctx    context.Context

	configReader interfaces.IConfigReader
}

func NewGPT3(configReader interfaces.IConfigReader) *GPT3dot5Turbo {
	apiKey := configReader.GetString("api-key")
	return &GPT3dot5Turbo{
		client:       openai.NewClient(apiKey),
		ctx:          context.Background(),
		configReader: configReader,
	}
}

func (g *GPT3dot5Turbo) GenerateText(prompt string) (string, error) {
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

func (g *GPT3dot5Turbo) GenerateStreamText(prompt string) (*openai.ChatCompletionStream, error) {
	stream, err := g.client.CreateChatCompletionStream(g.ctx, openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{{
			Role:    openai.ChatMessageRoleUser,
			Content: prompt,
		}},
		Stream:    true,
		MaxTokens: 20,
	})

	if err != nil {
		return nil, err
	}

	return stream, nil
}
