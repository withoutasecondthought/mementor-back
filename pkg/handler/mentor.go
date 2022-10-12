package handler

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	mementor_back "mementor-back"
	"net/http"
	"strconv"
)

func (h *Handler) getMentor(c echo.Context) error {
	id := c.Param("id")

	ctx := context.Background()
	mentor, err := h.services.GetMentor(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, mementor_back.Message{
			Message: fmt.Sprintf("Mentor not found: %s", err),
		})
		return err
	}

	return c.JSON(http.StatusOK, mentor)
}

func (h *Handler) putMentor(c echo.Context) error {
	validate := validator.New()
	var mentor mementor_back.Mentor

	err := c.Bind(&mentor)
	if err != nil {
		c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
		return err
	}
	err = validate.Struct(mentor)
	if err != nil {
		c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
		return err
	}

	ctx := context.Background()

	mentor.Id = &h.userId

	err = h.services.PutMentor(ctx, mentor)
	if err != nil {
		c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
		return err
	}
	return c.String(http.StatusOK, "ok")
}

func (h *Handler) deleteMentor(c echo.Context) error {
	ctx := context.Background()

	err := h.services.DeleteMentor(ctx, h.userId.Hex())
	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
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
		c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
		return err
	}

	err = c.Bind(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
		return err
	}

	mentors, err := h.services.ListOfMentors(ctx, uint(page), params)
	if err != nil {
		c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
		return err
	}

	c.JSON(http.StatusOK, mentors)
	return nil
}

func (h *Handler) getYourPage(c echo.Context) error {
	ctx := context.Background()
	mentor, err := h.services.GetMentor(ctx, h.userId.Hex())
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return err
	}

	return c.JSON(http.StatusOK, mentor)
}
