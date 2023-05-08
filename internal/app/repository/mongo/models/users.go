package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Users struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	APIKey    string             `bson:"api_key"`
	Chats     []Chat             `bson:"chats"`
	Timestamp time.Time          `bson:"timestamp"`
}
