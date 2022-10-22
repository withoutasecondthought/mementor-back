package handler

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	mementor_back "mementor-back"
	"net/http"
	"strconv"
)

// @Summary     Show a mentor
// @Description Give you mentor wuthout personal fields
// @Tags        mentor
// @Accept      json
// @Produce     string
// @Param       id  path     string true "Account ID"
// @Success     200 {object} mementor_back.MentorFullInfo
// @Failure     404 {object} mementor_back.Message
// @Router      /mentor/{id} [get]

func (h *Handler) getMentor(c echo.Context) error {
	id := c.Param("id")

	ctx := context.Background()
	mentor, err := h.Services.GetMentor(ctx, id)
	if err != nil {
		sentError := c.JSON(http.StatusNotFound, mementor_back.Message{
			Message: fmt.Sprintf("Mentor not found: %s", err),
		})
		if sentError != nil {
			logrus.Error(sentError)
		}
		return err
	}

	return c.JSON(http.StatusOK, mentor)
}

// @Summary     change mentor
// @Description get string by ID
// @Secure      ApiAuthKey
// @Tags        mentor
// @Accept      json
// @Produce     json
// @Param       user body     mementor_back.MentorFullInfo true "Account ID"
// @Success     200  {string} "ok"
// @Failure     400  {object} mementor_back.Message
// @Failure     401  {object} mementor_back.Message
// @Router      /mentor [put]

func (h *Handler) putMentor(c echo.Context) error {
	validate := validator.New()
	var mentor mementor_back.MentorFullInfo

	err := c.Bind(&mentor)
	if err != nil {
		sentError := c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
		if sentError != nil {
			logrus.Error(sentError)
		}
		return err
	}
	err = validate.Struct(mentor)
	if err != nil {
		sentError := c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
		if sentError != nil {
			logrus.Error(sentError)
		}
		return err
	}

	ctx := context.Background()

	mentor.Id = &h.UserId

	err = h.Services.PutMentor(ctx, mentor)
	if err != nil {
		sentError := c.JSON(http.StatusInternalServerError, mementor_back.Message{Message: err.Error()})
		if sentError != nil {
			logrus.Error(sentError)
		}
		return err
	}
	return c.String(http.StatusOK, "ok")
}

// @Summary     Delete mentor
// @Description remove mentor from bd
// @Secure      ApiAuthKey
// @Tags        mentor
// @Produce     string
// @Success     200 {string} "ok"
// @Failure     400 {object} mementor_back.Message
// @Failure     401 {object} mementor_back.Message
// @Router      /mementor [delete]

func (h *Handler) deleteMentor(c echo.Context) error {
	ctx := context.Background()

	err := h.Services.DeleteMentor(ctx, h.UserId.Hex())
	if err != nil {
		if err != nil {
			sentError := c.JSON(http.StatusInternalServerError, mementor_back.Message{Message: err.Error()})
			if sentError != nil {
				logrus.Error(sentError)
			}
			return err
		}
	}
	return c.String(http.StatusOK, "ok")
}

// @Summary     return List of Mentors
// @Description get mentors
// @Tags        mentor
// @Accept      json
// @Accept      path
// @Produce     json
// @Param       page   path     int       true  "number of page"
// @Param       params body     interface false "params"
// @Success     200    {object} mementor_back.Message
// @Failure     400    {object} httputil.HTTPError

// @Router /mentor/{page} [post]

func (h *Handler) listOfMentors(c echo.Context) error {
	var params interface{}
	ctx := context.Background()
	p := c.Param("page")
	if p == "" {
		p = "0"
	}
	page, err := strconv.Atoi(p)
	if err != nil {
		sentError := c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
		if sentError != nil {
			logrus.Error(sentError)
		}
		return err
	}

	err = c.Bind(&params)
	if err != nil {
		sentError := c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
		if sentError != nil {
			logrus.Error(sentError)
		}
		return err
	}

	mentors, err := h.Services.ListOfMentors(ctx, uint(page), params)
	if err != nil {
		sentError := c.JSON(http.StatusInternalServerError, mementor_back.Message{Message: err.Error()})
		if sentError != nil {
			logrus.Error(sentError)
		}
		return err
	}

	return c.JSON(http.StatusOK, mentors)
}

// @Summary     Show your Page
// @Description get your page
// @Secure      ApiAuthKey
// @Tags        mentor
// @Produce     json
// @Success     200 {object} mementor_back.MentorFullInfo
// @Failure     400 {object} mementor_back.Message
// @Failure     401 {object} mementor_back.Message
// @Router      /mentor [get]

func (h *Handler) getYourPage(c echo.Context) error {
	ctx := context.Background()
	mentor, err := h.Services.GetMyMentor(ctx, h.UserId.Hex())
	if err != nil {
		sentError := c.JSON(http.StatusInternalServerError, mementor_back.Message{Message: err.Error()})
		if sentError != nil {
			logrus.Error(sentError)
		}
		return err
	}

	return c.JSON(http.StatusOK, mentor)
}
