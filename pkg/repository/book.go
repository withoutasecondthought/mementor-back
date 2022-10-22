package repository

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	mementor_back "mementor-back"
)

type BookMongo struct {
	db *mongo.Collection
}

func (b *BookMongo) NewBooking(ctx context.Context, booking mementor_back.Booking) error {
	_, err := b.db.InsertOne(ctx, booking)

	return err
}

func NewBookMongo(db *mongo.Database) *BookMongo {
	return &BookMongo{db: db.Collection(viper.GetString("db.book-collection"))}
}
