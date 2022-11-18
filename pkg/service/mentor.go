package service

import (
	"context"
	"errors"
	mementor_back "mementor-back"
	"mementor-back/pkg/repository"
	"time"
)

type MentorService struct {
	repos repository.Mentor
}

func (m *MentorService) GetMentor(ctx context.Context, id string) (mementor_back.MentorFullInfo, error) {
	return m.repos.GetMentor(ctx, id)
}

func (m *MentorService) GetMyMentor(ctx context.Context, id string) (mementor_back.MentorFullInfo, error) {
	return m.repos.GetMyMentor(ctx, id)
}

func (m *MentorService) PutMentor(ctx context.Context, mentor mementor_back.MentorFullInfo) error {
	mentor.ValidProfile = true
	if mentor.Tariff[0].Price > mentor.Tariff[1].Price || mentor.Tariff[1].Price > mentor.Tariff[2].Price {
		return errors.New("wrong tariffs order")
	}
	return m.repos.PutMentor(ctx, mentor)
}

func (m *MentorService) DeleteMentor(ctx context.Context, id string) error {
	return m.repos.DeleteMentor(ctx, id)
}

func (m *MentorService) ListOfMentors(ctx context.Context, page uint, params mementor_back.SearchParameters) (mementor_back.ListOfMentorsResponse, error) {
	params.ValidProfile = true

	if params.MaxPrice == 0 {
		params.MaxPrice = 100000000
	}

	if params.MaxPrice < params.MinPrice {
		return mementor_back.ListOfMentorsResponse{Mentors: []mementor_back.Mentor{}}, errors.New("min price can't be greater than max")
	}
	if len(params.Grade) == 0 {
		params.Grade = []string{"junior", "middle", "senior"}
	}
	if params.ExperienceSince == 0 {
		params.ExperienceSince = time.Now().Year()
	}

	resp, err := m.repos.ListOfMentors(ctx, page, params)

	if err == nil && resp.Mentors != nil {
		resp.Pages += 1
	}

	if resp.Mentors == nil {
		resp.Mentors = []mementor_back.Mentor{}
	}

	return resp, err

}

func NewMentorService(repo repository.Mentor) *MentorService {
	return &MentorService{
		repos: repo,
	}
}
