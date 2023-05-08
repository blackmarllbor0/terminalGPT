package models

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	ID        primitive.ObjectID             `bson:"_id,omitempty"`
	Content   []openai.ChatCompletionMessage `bson:"content"`
	Timestamp time.Time                      `bson:"timestamp"`
}

type Chats struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewChats(collection *mongo.Collection, ctx context.Context) *Chats {
	return &Chats{
		collection: collection,
		ctx:        ctx,
	}
}

func (c Chats) CreateNewChat() (primitive.ObjectID, error) {
	chat := &Chat{Timestamp: time.Now()}
	result, err := c.collection.InsertOne(c.ctx, chat)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func (c Chats) GetAllChats() ([]Chat, error) {
	var chatHistory []Chat

	cursor, err := c.collection.Find(c.ctx, bson.D{})
	if err != nil {
		return chatHistory, err
	}
	defer cursor.Close(c.ctx)

	for cursor.Next(c.ctx) {
		var chat Chat
		if err := cursor.Decode(&chat); err != nil {
			return chatHistory, err
		}

		chatHistory = append(chatHistory, chat)
	}

	if cursor.Err() != nil {
		return chatHistory, err
	}

	return chatHistory, nil
}

func (c Chats) UpdateChatHistory(chatID primitive.ObjectID, content []openai.ChatCompletionMessage) error {
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "content", Value: content},
		}},
	}

	_, err := c.collection.UpdateOne(c.ctx, bson.D{{Key: "_id", Value: chatID}}, update)
	if err != nil {
		return err
	}

	return nil
}

func (c Chats) GetChat(chatID primitive.ObjectID) (Chat, error) {
	var chat Chat

	if err := c.collection.FindOne(c.ctx, bson.M{"_id": chatID}).Decode(&chat); err != nil {
		if err == mongo.ErrNoDocuments {
			return Chat{}, fmt.Errorf("chat with ID: %v not found", chatID.Hex())
		}

		return Chat{}, fmt.Errorf("failed to get chat: %v", err)
	}

	return chat, nil
}
