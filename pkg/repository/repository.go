package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	mementor_back "mementor-back"
)

type Mentor interface {
	GetMentor(ctx context.Context, id string) (mementor_back.Mentor, error)
	PutMentor(ctx context.Context, mentor mementor_back.Mentor) error
	DeleteMentor(ctx context.Context, id string) error
	ListOfMentors(ctx context.Context, page uint, params interface{}) ([]*mementor_back.Mentor, error)
}

type Authorization interface {
	CreateUser(ctx context.Context, user mementor_back.Auth) (string, error)
	GetUser(ctx context.Context, user mementor_back.Auth) (string, error)
}

type Repository struct {
	Mentor
	Authorization
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Mentor:        NewMentor(db),
		Authorization: NewAuthMongo(db),
	}

}
