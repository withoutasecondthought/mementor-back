package handler

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	mementor_back "mementor-back"
	"net/http"
)

// @Summary     New Booking
// @Description book mentor of your dream
// @Tags        booking
// @Accept      json
// @Produce     json
// @Param       params body     mementor_back.Booking true "params"
// @Success     200    {object} mementor_back.Message
// @Failure     400    {object} mementor_back.Message
// @Router /book [post]
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
