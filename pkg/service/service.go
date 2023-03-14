package service

import (
	"context"
	mementor_back "mementor-back"
	"mementor-back/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Mentor interface {
	GetMentor(ctx context.Context, id string) (mementor_back.MentorFullInfo, error)
	GetMyMentor(ctx context.Context, id string) (mementor_back.MentorFullInfo, error)
	PutMentor(ctx context.Context, mentor mementor_back.MentorFullInfo) error
	DeleteMentor(ctx context.Context, id string) error
	ListOfMentors(ctx context.Context, page uint, params mementor_back.SearchParameters) (mementor_back.ListOfMentorsResponse, error)
}

type Authorization interface {
	SignIn(ctx context.Context, user mementor_back.Auth) (string, error)
	SignUp(ctx context.Context, user mementor_back.Auth) (string, error)
}

type Book interface {
	NewBooking(ctx context.Context, booking mementor_back.Booking) error
}

type Image interface {
	NewImage(ctx context.Context, image mementor_back.PostImage) (mementor_back.Image, error)
}

type Service struct {
	Mentor
	Authorization
	Book
	Image
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Mentor:        NewMentorService(repo.Mentor),
		Authorization: NewAuthService(repo.Authorization),
		Book:          NewBooking(repo.Book),
		Image:         NewImageService(repo.Image),
	}
}
