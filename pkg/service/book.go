package service

import (
	"context"
	mementor_back "mementor-back"
	"mementor-back/pkg/repository"
)

type Booking struct {
	repo repository.Book
}

func (b *Booking) NewBooking(ctx context.Context, booking mementor_back.Booking) error {
	return b.repo.NewBooking(ctx, booking)
}

func NewBooking(repo repository.Book) *Booking {
	return &Booking{repo: repo}
}
