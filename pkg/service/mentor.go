package service

import (
	"context"
	"errors"
	mementor_back "mementor-back"
	"mementor-back/pkg/repository"
)

type MentorService struct {
	repos repository.Mentor
}

func (m *MentorService) GetMentor(ctx context.Context, id string) (mementor_back.Mentor, error) {
	return m.repos.GetMentor(ctx, id)
}

func (m *MentorService) GetMyMentor(ctx context.Context, id string) (mementor_back.Mentor, error) {
	return m.repos.GetMyMentor(ctx, id)
}

func (m *MentorService) PutMentor(ctx context.Context, mentor mementor_back.Mentor) error {
	var grade = map[string]int{
		"junior": 0,
		"middle": 1,
		"senior": 2,
	}
	if _, exist := grade[mentor.Grade]; !exist {
		return errors.New("grade invalid")
	}
	mentor.ValidProfile = true
	return m.repos.PutMentor(ctx, mentor)
}

func (m *MentorService) DeleteMentor(ctx context.Context, id string) error {
	return m.repos.DeleteMentor(ctx, id)
}

func (m *MentorService) ListOfMentors(ctx context.Context, page uint, params interface{}) ([]*mementor_back.Mentor, error) {
	return m.repos.ListOfMentors(ctx, page, params)
}

func NewMentorService(repo repository.Mentor) *MentorService {
	return &MentorService{
		repos: repo,
	}
}
