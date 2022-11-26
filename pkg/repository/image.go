package repository

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	mementor_back "mementor-back"
)

type ImageMongo struct {
	db *mongo.Collection
}

func (i *ImageMongo) NewImage(ctx context.Context, image mementor_back.PostImage) error {
	_, err := i.db.UpdateByID(ctx, image.ID, bson.M{"$set": bson.M{"image": image.Image}})
	return err
}

func NewImageMongo(db *mongo.Database) *ImageMongo {
	return &ImageMongo{db: db.Collection(viper.GetString("db.collection"))}
}
