package handler

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	mementor_back "mementor-back"
	"net/http"
)

type PostImageResponse struct {
	FullImage  string `json:"256x256"`
	SmallImage string `json:"72x72"`
}

func (h *Handler) PostImage(c echo.Context) error {
	var req mementor_back.PostImage
	var resp PostImageResponse
	req.Id = &h.UserId

	err := c.Bind(&req)
	if err != nil {
		sentError := c.JSON(http.StatusBadRequest, mementor_back.Message{Message: err.Error()})
		if sentError != nil {
			logrus.Error(sentError)
		}
		return err
	}

	cld, err := cloudinary.NewFromParams("dnpjr8yvk", "345764535751477", "C4au6sUKzqT_beAC0zYvz712p1Y")
	if err != nil {
		sentError := c.JSON(http.StatusInternalServerError, mementor_back.Message{Message: err.Error()})
		if sentError != nil {
			logrus.Error(sentError)
		}
		return err
	}
	cld.Config.URL.Secure = true

	ImagePublicId := uuid.New().String()

	_, err = cld.Upload.Upload(context.Background(), req.Base64, uploader.UploadParams{
		PublicID: ImagePublicId,
	})
	if err != nil {
		sentError := c.JSON(http.StatusInternalServerError, mementor_back.Message{Message: err.Error()})
		if sentError != nil {
			logrus.Error(sentError)
		}
		return err
	}

	Image, err := cld.Image(ImagePublicId)
	if err != nil {
		sentError := c.JSON(http.StatusInternalServerError, mementor_back.Message{Message: err.Error()})
		if sentError != nil {
			logrus.Error(sentError)
		}
		return err
	}

	Image.Transformation = "h_256,w_256"

	Image256, err := Image.String()
	if err != nil {
		sentError := c.JSON(http.StatusInternalServerError, mementor_back.Message{Message: err.Error()})
		if sentError != nil {
			logrus.Error(sentError)
		}
		return err
	}

	resp.FullImage = Image256

	Image.Transformation = "h_72,w_72"
	Image72, err := Image.String()
	if err != nil {
		sentError := c.JSON(http.StatusInternalServerError, mementor_back.Message{Message: err.Error()})
		if sentError != nil {
			logrus.Error(sentError)
		}
		return err
	}
	resp.SmallImage = Image72

	return c.JSON(200, resp)
}
