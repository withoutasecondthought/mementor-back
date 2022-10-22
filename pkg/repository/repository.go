package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	mementor_back "mementor-back"
)

type Mentor interface {
	GetMentor(ctx context.Context, id string) (mementor_back.MentorFullInfo, error)
	GetMyMentor(ctx context.Context, id string) (mementor_back.MentorFullInfo, error)
	PutMentor(ctx context.Context, mentor mementor_back.MentorFullInfo) error
	DeleteMentor(ctx context.Context, id string) error
	ListOfMentors(ctx context.Context, page uint, params interface{}) (mementor_back.ListOfMentorsResponse, error)
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
