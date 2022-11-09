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
// @Description Give you mentor without personal fields
// @Tags        mentor
// @Produce     json
// @Param       id  path     string true "Account ID"
// @Success     200 {object} mementor_back.MentorFullInfo
// @Failure     404 {object} mementor_back.Message "error occured"
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
// @Description Send full mentor info to update your profile
// @Secure      ApiAuthKey
// @Tags        mentor
// @Accept      json
// @Produce     json
// @Param       user body     mementor_back.PutMentorRequest true "Account info"
// @Success     200  {object} mementor_back.Message "ok"
// @Failure     400  {object} mementor_back.Message "error occurred"
// @Failure     401  {object} mementor_back.Message "Unauthorized"
// @Failure     500  {object} mementor_back.Message "error occurred"
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
	return c.JSON(http.StatusOK, mementor_back.Message{Message: "ok"})
}

// @Summary     Delete mentor
// @Description remove mentor from bd
// @Secure      ApiAuthKey
// @Tags        mentor
// @Produce     json
// @Success     200 {object} mementor_back.Message "ok"
// @Failure     500 {object} mementor_back.Message "error occurred"
// @Router      /mentor [delete]
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
	return c.JSON(http.StatusOK, mementor_back.Message{Message: "ok"})
}

// @Summary     return List of Mentors
// @Description get mentors
// @Tags        mentor
// @Accept      json
// @Produce     json
// @Param       page   path     int       true  "number of page"
// @Param       params body     mementor_back.SearchParameters false "params"
// @Success     200    {object} mementor_back.ListOfMentorsResponse
// @Failure     400  {object} mementor_back.Message "error occurred"
// @Failure     500  {object} mementor_back.Message "error occurred"
// @Router /mentor/{page} [post]
func (h *Handler) listOfMentors(c echo.Context) error {
	var params mementor_back.SearchParameters
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

	if mentors.Mentors == nil {
		mentors.Mentors = []mementor_back.Mentor{}
	}

	return c.JSON(http.StatusOK, mentors)

}

// @Summary     Show your Page
// @Description get your page
// @Secure      ApiAuthKey
// @Tags        mentor
// @Produce     json
// @Success     200 {object} mementor_back.MentorFullInfo
// @Failure     500 {object} mementor_back.Message "error occurred"
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
