package handler

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	mementor_back "mementor-back"
	"net/http"
)

func (h *Handler) signUp(c echo.Context) error {
	validate := validator.New()
	var user mementor_back.Auth

	err := c.Bind(&user)
	if err != nil {
		sendError := c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
		if sendError != nil {
			logrus.Error(sendError)
		}
		return err
	}
	err = validate.Struct(user)
	if err != nil {
		sendError := c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
		if sendError != nil {
			logrus.Error(sendError)
		}
		return err
	}

	ctx := context.Background()

	token, err := h.Services.SignUp(ctx, user)
	if err != nil {
		sendError := c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
		if sendError != nil {
			logrus.Error(sendError)
		}
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) signIn(c echo.Context) error {
	validate := validator.New()
	var user mementor_back.Auth

	err := c.Bind(&user)
	if err != nil {
		sendError := c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
		if sendError != nil {
			logrus.Error(sendError)
		}
		return err
	}
	err = validate.Struct(user)
	if err != nil {
		sendError := c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
		if sendError != nil {
			logrus.Error(sendError)
		}
		return err
	}

	ctx := context.Background()

	token, err := h.Services.SignIn(ctx, user)
	if err != nil {
		sendError := c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
		if sendError != nil {
			logrus.Error(sendError)
		}
		return err
	}
	sendError := c.JSON(http.StatusOK, map[string]string{"token": token})
	if sendError != nil {
		logrus.Error(sendError)
	}
	return nil
}
