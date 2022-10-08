package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

type MentorMongo struct {
	db *mongo.Collection
}

func NewMentor(db *mongo.Database) *MentorMongo {
	return &MentorMongo{
		db: db.Collection(os.Getenv("DB.COLLECTION")),
	}
}
