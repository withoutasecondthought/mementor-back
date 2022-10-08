package service

import "mementor-back/pkg/repository"

type MentorService struct {
	repos repository.Mentor
}

func NewMentorService(repo repository.Mentor) *MentorService {
	return &MentorService{
		repos: repo,
	}
}
