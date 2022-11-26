package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	mementor_back "mementor-back"
	"net/http"
)

// @Summary     sign up
// @Description sign up
// @Tags        auth
// @ID sign-up
// @Accept      json
// @Produce     json
// @Param      userinfo body mementor_back.Auth true "Account data"
// @Success     200 {object} loginResponse
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

	token, err := h.Services.SignUp(c.Request().Context(), user)
	if err != nil {
		sendError := c.JSON(http.StatusInternalServerError, mementor_back.Message{Message: err.Error()})
		if sendError != nil {
			logrus.Error(sendError)
		}
		return err
	}
	return c.JSON(http.StatusOK, loginResponse{Token: token})
}

// @Summary     sign in
// @Description sign in
// @Tags        auth
// @ID sign-in
// @Accept      json
// @Produce     json
// @Param      userinfo body mementor_back.Auth true "Account data"
// @Success     200 {object} loginResponse
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

	token, err := h.Services.SignIn(c.Request().Context(), user)
	if err != nil {
		sendError := c.JSON(http.StatusInternalServerError, mementor_back.Message{Message: err.Error()})
		if sendError != nil {
			logrus.Error(sendError)
		}
		return err
	}
	sendError := c.JSON(http.StatusOK, loginResponse{Token: token})
	if sendError != nil {
		logrus.Error(sendError)
	}
	return nil
}

type loginResponse struct {
	Token string `json:"token" example:"Bearer token"`
} // @name PostAuthResponse
