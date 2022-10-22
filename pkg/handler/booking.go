package handler

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	mementor_back "mementor-back"
	"net/http"
)

func (h *Handler) newBooking(c echo.Context) error {
	var booking mementor_back.Booking
	validate := validator.New()

	err := c.Bind(&booking)
	if err != nil {
		sentError := c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
		if sentError != nil {
			logrus.Error(sentError)
		}
		return err
	}
	err = validate.Struct(booking)
	if err != nil {
		sentError := c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
		if sentError != nil {
			logrus.Error(sentError)
		}
		return err
	}
	ctx := context.Background()

	err = h.Services.Book.NewBooking(ctx, booking)
	if err != nil {
		sentError := c.JSON(http.StatusInternalServerError, mementor_back.Message{Message: err.Error()})
		if sentError != nil {
			logrus.Error(sentError)
		}
		return err
	}

	return c.JSON(http.StatusOK, mementor_back.Message{
		Message: "ok",
	})
}
