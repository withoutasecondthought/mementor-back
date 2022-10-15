package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
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
			return c.JSON(http.StatusUnauthorized, mementor_back.Message{Message: errorTXT})

		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 {
			return c.JSON(http.StatusUnauthorized, mementor_back.Message{Message: errorTXT})
		}

		if headerParts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, mementor_back.Message{Message: errorTXT})
		}

		token, err := parser.ParseToken(headerParts[1], []byte(viper.GetString("signing_key")))
		if err != nil {
			sendError := c.JSON(http.StatusUnauthorized, mementor_back.Message{Message: err.Error()})
			if sendError != nil {
				logrus.Error(sendError)
			}
			return err
		}

		id, err := primitive.ObjectIDFromHex(token)
		if err != nil {
			sendError := c.JSON(http.StatusUnauthorized, mementor_back.Message{Message: err.Error()})
			if sendError != nil {
				logrus.Error(sendError)
			}
			return err
		}

		h.UserId = id

		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}
