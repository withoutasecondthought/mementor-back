package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mementor-back/pkg/service"
)

type Handler struct {
	services *service.Service
	userId   primitive.ObjectID
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		services: service,
	}
}

func (h *Handler) InitRoutes() *echo.Echo {
	e := echo.New()
	log := logrus.New()

	e.Use(middleware.Recover())

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(logrus.Fields{
				"URI":    values.URI,
				"status": values.Status,
				"error":  values.Error,
			}).Info("request")

			return nil
		},
	}))

	e.POST("/sign-in", h.signIn)
	e.POST("/sign-up", h.signUp)

	e.GET("mentor/:id", h.getMentor)
	e.POST("mentor/:page", h.listOfMentors)

	mentor := e.Group("mentor", h.parseJWT)
	{
		mentor.GET("/", h.getYourPage)
		mentor.PUT("/", h.putMentor)
		mentor.DELETE("/", h.deleteMentor)

	}

	return e
}
