package service

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	mementor_back "mementor-back"
	"mementor-back/pkg/repository"
	"net/smtp"
)

type Booking struct {
	repo repository.Book
}

func (b *Booking) NewBooking(ctx context.Context, booking mementor_back.Booking) error {
	err := b.repo.NewBooking(ctx, booking)
	if err != nil {
		return err
	}
	return sendBookEmail(booking)
}

func sendBookEmail(booking mementor_back.Booking) error {
	from := viper.GetString("gmail.from")
	pass := viper.GetString("gmail.password")
	to := viper.GetString("gmail.to")

	sub := fmt.Sprintf("Subject: New Booking Request from %s\r\n\r\n", booking.CustomerName)
	msg := fmt.Sprintf("Name: %s\n\n Telegram: %s\n\n MentorId: %s \n\n MentorLink: https://ilyaprojects.com/test-drive/mementor/mentor/%s\n\n TarriffIndex: %d", booking.CustomerName, booking.CustomerTelegram, booking.MentorID, booking.MentorID, booking.TariffIndex)

	auth := smtp.PlainAuth("", from, pass, "smtp.gmail.com")

	err := smtp.SendMail("smtp.gmail.com:587", auth,
		from, []string{to}, []byte(sub+msg))

	if err != nil {
		return fmt.Errorf("smtp error: %w", err)
	}
	return nil
}

func NewBooking(repo repository.Book) *Booking {
	return &Booking{repo: repo}
}
