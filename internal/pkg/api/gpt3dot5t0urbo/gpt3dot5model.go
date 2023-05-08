package gpt3dot5t0urbo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"terminalGPT/config/interfaces"
	"terminalGPT/internal/app/repository/interfaces/database"

	"github.com/sashabaranov/go-openai"
)

type GPT3dot5Turbo struct {
	client   *openai.Client
	ctx      context.Context
	messages []openai.ChatCompletionMessage

	configReader interfaces.ConfigReader

	database.Chats
}

func NewGPT3(configReader interfaces.ConfigReader, chatsRepository database.Chats) *GPT3dot5Turbo {
	apiKey := configReader.GetString("api-key")
	return &GPT3dot5Turbo{
		client:       openai.NewClient(apiKey),
		ctx:          context.Background(),
		configReader: configReader,
		Chats:        chatsRepository,
	}
}

func (g *GPT3dot5Turbo) GenerateText(prompt string, chatID primitive.ObjectID) (string, error) {
	g.messages = append(g.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: prompt,
	})

	response, err := g.client.CreateChatCompletion(g.ctx, openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo,
		Messages: g.messages,
	})

	if err != nil {
		return "", err
	}

	content := response.Choices[0].Message.Content
	g.messages = append(g.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: content,
	})

	go func() {
		if err := g.UpdateChatHistory(chatID, g.messages); err != nil {
			log.Printf("a write error occurred in the database: %v", err)
		}
	}()

	return content, nil
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
func (g *GPT3dot5Turbo) UpdateChatMessages(chatID primitive.ObjectID) error {
	currentChat, err := g.GetChat(chatID)
	if err != nil {
		return err
	}

	if len(currentChat.Content) > 0 {
		g.messages = currentChat.Content
	}

	return nil
}
