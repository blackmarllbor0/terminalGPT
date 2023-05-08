package database

import (
	"github.com/sashabaranov/go-openai"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"terminalGPT/internal/app/repository/mongo/models"
)

type Chats interface {
	GetAllChats() ([]models.Chat, error)
	GetChat(chatID primitive.ObjectID) (models.Chat, error)
	CreateNewChat() (primitive.ObjectID, error)
	UpdateChatHistory(chatID primitive.ObjectID, content []openai.ChatCompletionMessage) error
}
