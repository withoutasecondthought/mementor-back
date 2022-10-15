package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mementor-back/pkg/service"
)

type Handler struct {
	Services *service.Service
	UserId   primitive.ObjectID
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		Services: service,
	}
}

// @title Mementor back
// @version         1.0
// @description     Best backend ever.

// @contact.name   @withoutasecondthought
// @contact.email    mrmarkeld@gmail.com

// @host      api.ilyaprojects.com/
// @BasePath  /mementor

func (h *Handler) InitRoutes() *echo.Echo {
	e := echo.New()
	log := logrus.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Static(""))

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

	e.GET("mentor/:id", h.getMentor)
	e.POST("mentor/:page", h.listOfMentors)

	mentor := e.Group("mentor", h.parseJWT)
	{
		mentor.GET("", h.getYourPage)
		mentor.PUT("", h.putMentor)
		mentor.DELETE("", h.deleteMentor)

	}

	return e
}
