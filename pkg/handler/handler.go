package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "mementor-back/docs"
	"mementor-back/pkg/service"
)

type Handler struct {
	Services *service.Service
	UserID   primitive.ObjectID
}

type customSerialize struct {
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		Services: service,
	}
}

func (h *Handler) InitRoutes() *echo.Echo {
	e := echo.New()
	log := logrus.New()

	e.JSONSerializer = customSerialize{}

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

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

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/sign-in", h.signIn)
	e.POST("/sign-up", h.signUp)

	e.POST("/book", h.newBooking)

	mentor := e.Group("/mentor")
	{
		mentor.GET("/:id", h.getMentor)
		mentor.POST("/:page", h.listOfMentors)

		authmentor := mentor.Group("", h.parseJWT)
		{
			authmentor.POST("/image", h.PostImage)
			authmentor.GET("", h.getYourPage)
			authmentor.PUT("", h.putMentor)
			authmentor.DELETE("", h.deleteMentor)
		}
	}
	return e
}
