package service

import "mementor-back/pkg/repository"

type Mentor interface {
}

type Service struct {
	Mentor
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Mentor: NewMentorService(repos.Mentor),
	}
}
