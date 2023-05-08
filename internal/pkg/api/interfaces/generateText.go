package interfaces

import (
	"github.com/sashabaranov/go-openai"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"terminalGPT/internal/app/repository/interfaces/database"
)

type GPTGenerateText interface {
	GenerateText(prompt string, chatID primitive.ObjectID) (string, error)
	GenerateStreamText(prompt string) (*openai.ChatCompletionStream, error)
	UpdateChatMessages(chatID primitive.ObjectID) error
	database.Chats
}
