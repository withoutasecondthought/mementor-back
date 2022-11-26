package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	mementor_back "mementor-back"
	"net/http"
)

// @Summary     Upload Image
// @Description Upload your best photo
// @Secure      ApiAuthKey
// @Tags        mentor
// @Accept      json
// @Produce     json
// @Param       user body     mementor_back.PostImage true "base64"
// @Success     200  {object} mementor_back.Image "ok"
// @Failure     400  {object} mementor_back.Message "error occurred"
// @Failure     401  {object} mementor_back.Message "Unauthorized"
// @Failure     500  {object} mementor_back.Message "error occurred"
// @Router      /mentor/image [post]
func (h *Handler) PostImage(c echo.Context) error {
	var req mementor_back.PostImage
	req.ID = &h.UserID

	err := c.Bind(&req)
	if err != nil {
		sentError := c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
		if sentError != nil {
			logrus.Error(sentError)
		}
		return err
	}

	image, err := h.Services.Image.NewImage(c.Request().Context(), req)
	if err != nil {
		sentError := c.JSON(http.StatusInternalServerError, mementor_back.Message{Message: err.Error()})
		if sentError != nil {
			logrus.Error(sentError)
		}
		return err
	}

	return c.JSON(http.StatusOK, image)
}
