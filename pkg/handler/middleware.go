package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mementor_back "mementor-back"
	"mementor-back/pkg/parser"
	"net/http"
	"strings"
)

func (h *Handler) parseJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		errorTXT := "token invalid"
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, mementor_back.Message{Message: "1"})
			return errors.New(errorTXT)
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 {
			c.JSON(http.StatusUnauthorized, mementor_back.Message{Message: "2"})
			return errors.New(errorTXT)
		}

		if headerParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, mementor_back.Message{Message: "3"})
			return errors.New(errorTXT)
		}

		token, err := parser.ParseToken(headerParts[1], []byte(viper.GetString("signing_key")))
		if err != nil {
			c.JSON(http.StatusUnauthorized, mementor_back.Message{Message: err.Error()})
			return err
		}

		id, err := primitive.ObjectIDFromHex(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, mementor_back.Message{Message: err.Error()})
			return err
		}

		h.userId = id

		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}
