package repository

import (
	"context"
	"errors"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	mementor_back "mementor-back"
)

type AuthMongo struct {
	db *mongo.Collection
}

type getId struct {
	Id *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
}

func (a *AuthMongo) CreateUser(ctx context.Context, user mementor_back.Auth) (string, error) {
	test, err := a.db.CountDocuments(ctx, bson.M{"email": user.Email})
	if test != 0 {
		return "", errors.New("you have already creates account")
	}
	if err != nil {
		return "", err
	}
	res, err := a.db.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (a *AuthMongo) GetUser(ctx context.Context, user mementor_back.Auth) (string, error) {
	var response getId
	err := a.db.FindOne(ctx, bson.M{"email": user.Email, "password": user.Password}).Decode(&response)
	if err != nil {
		return "", err
	}

	if response.Id == nil {
		return "", errors.New("no such user")
	}

	return response.Id.Hex(), nil
}

func NewAuthMongo(db *mongo.Database) *AuthMongo {
	return &AuthMongo{db: db.Collection(viper.GetString("db.collection"))}
}
