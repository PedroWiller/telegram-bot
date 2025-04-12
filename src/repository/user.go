package repository

import (
	"context"

	"telegram-bot/src/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type UserRepository struct {
	db *mongo.Client
}

func NewUserRepository(db *mongo.Client) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetBydId(id int64) (model.User, error) {
	filter := bson.M{"externalId": id}
	var user model.User
	err := r.db.Database("telegram").Collection("users").FindOne(context.Background(), filter).Decode(&user)
	if err != nil && err != mongo.ErrNoDocuments {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) Save(user model.User) error {
	filter := bson.M{"externalId": user.ExternalId}
	update := bson.M{"$set": user}
	_, err := r.db.Database("telegram").Collection("users").UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetAll() ([]model.User, error) {
	filter := bson.M{"sendMessage": true}
	cursor, err := r.db.Database("telegram").Collection("users").Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var users []model.User
	for cursor.Next(context.Background()) {
		var user model.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
