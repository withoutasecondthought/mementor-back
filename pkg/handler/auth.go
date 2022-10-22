package handler

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	mementor_back "mementor-back"
	"net/http"
)

// @Summary     sign up
// @Description sign up
// @Tags        auth
// @Accept      json
// @Produce     json
// @Params      user   body      mementor_back.Auth  true  "Account data"
// @Success     200 {object} interface{}
// @Failure     400 {object} mementor_back.Message
// @Router      /sign-up [post]

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

// @Summary     sign in
// @Description sign in
// @Tags        auth
// @Accept      json
// @Produce     json
// @Params      user   body      mementor_back.Auth  true  "Account data"
// @Success     200 {object} interface{}
// @Failure     400 {object} mementor_back.Message
// @Router      /sign-in [post]
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
