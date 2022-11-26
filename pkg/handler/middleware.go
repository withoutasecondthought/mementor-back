package handler

import (
	jsonold "encoding/json"
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mementor_back "mementor-back"
	"mementor-back/pkg/parser"
	"net/http"
	"strings"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

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

		h.UserID = id

		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}

func (ser customSerialize) Serialize(c echo.Context, i interface{}, indent string) error {
	enc := json.NewEncoder(c.Response())
	if indent != "" {
		enc.SetIndent("", indent)
	}
	return enc.Encode(i)
}

func (ser customSerialize) Deserialize(c echo.Context, i interface{}) error {
	err := json.NewDecoder(c.Request().Body).Decode(i)
	var ute *jsonold.UnmarshalTypeError
	var se *jsonold.SyntaxError
	if ok := errors.Is(err, ute); ok {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unmarshal type error: expected=%v, got=%v, field=%v, offset=%v", ute.Type, ute.Value, ute.Field, ute.Offset)).SetInternal(err)
	} else if ok := errors.Is(err, se); ok {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Syntax error: offset=%v, error=%v", se.Offset, se.Error())).SetInternal(err)
	}
	return err
}
