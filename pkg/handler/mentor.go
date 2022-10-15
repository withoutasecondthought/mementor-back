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

func (h *Handler) putMentor(c echo.Context) error {
	validate := validator.New()
	var mentor mementor_back.Mentor

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
		sentError := c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
		if sentError != nil {
			logrus.Error(sentError)
		}
		return err
	}
	return c.String(http.StatusOK, "ok")
}

func (h *Handler) deleteMentor(c echo.Context) error {
	ctx := context.Background()

	err := h.Services.DeleteMentor(ctx, h.UserId.Hex())
	if err != nil {
		if err != nil {
			sentError := c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
			if sentError != nil {
				logrus.Error(sentError)
			}
			return err
		}
	}
	return c.String(http.StatusOK, "ok")
}

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
		sentError := c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
		if sentError != nil {
			logrus.Error(sentError)
		}
		return err
	}

	return c.JSON(http.StatusOK, mentors)
}

func (h *Handler) getYourPage(c echo.Context) error {
	ctx := context.Background()
	mentor, err := h.Services.GetMyMentor(ctx, h.UserId.Hex())
	if err != nil {
		sentError := c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
		if sentError != nil {
			logrus.Error(sentError)
		}
		return err
	}

	return c.JSON(http.StatusOK, mentor)
}
