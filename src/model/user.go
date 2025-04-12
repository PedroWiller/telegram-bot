package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	ExternalId    int64              `bson:"externalId"`
	FirstName     string             `bson:"firstName"`
	UserName      string             `bson:"userName"`
	SendedMessage bool               `bson:"sendMessage"`
	ChatId        int64              `bson:"chatId"`
}
