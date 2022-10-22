package service

import (
	"context"
	mementor_back "mementor-back"
	"mementor-back/pkg/repository"
)

type Mentor interface {
	GetMentor(ctx context.Context, id string) (mementor_back.MentorFullInfo, error)
	GetMyMentor(ctx context.Context, id string) (mementor_back.MentorFullInfo, error)
	PutMentor(ctx context.Context, mentor mementor_back.MentorFullInfo) error
	DeleteMentor(ctx context.Context, id string) error
	ListOfMentors(ctx context.Context, page uint, params interface{}) (mementor_back.ListOfMentorsResponse, error)
}

type Authorization interface {
	SignIn(ctx context.Context, user mementor_back.Auth) (string, error)
	SignUp(ctx context.Context, user mementor_back.Auth) (string, error)
}

type Service struct {
	Mentor
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Mentor:        NewMentorService(repos.Mentor),
		Authorization: NewAuthService(repos.Authorization),
	}
}
