package handler

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	mementor_back "mementor-back"
	"net/http"
)

func (h *Handler) signUp(c echo.Context) error {
	validate := validator.New()
	var user mementor_back.Auth

	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
		return err
	}
	err = validate.Struct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
		return err
	}

	ctx := context.Background()

	token, err := h.services.SignUp(ctx, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
		return err
	}
	return c.JSON(http.StatusOK, token)
}

func (h *Handler) signIn(c echo.Context) error {
	validate := validator.New()
	var user mementor_back.Auth

	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
		return err
	}
	err = validate.Struct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
		return err
	}

	ctx := context.Background()

	token, err := h.services.SignIn(ctx, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
		return err
	}
	return c.JSON(http.StatusOK, token)
}
