package service

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	mementor_back "mementor-back"
	"mementor-back/pkg/repository"
)

type ImageService struct {
	repo repository.Image
}

func (i *ImageService) NewImage(ctx context.Context, image mementor_back.PostImage) (mementor_back.Image, error) {
	var resp mementor_back.Image

	cld, err := cloudinary.NewFromParams(viper.GetString("cloudinary.cloud"), viper.GetString("cloudinary.api"), viper.GetString("cloudinary.secret"))
	if err != nil {
		return mementor_back.Image{}, err
	}
	cld.Config.URL.Secure = true

	imagePublicId := uuid.New().String()

	bigImage, err := cld.Upload.Upload(ctx, image.Base64, uploader.UploadParams{
		AllowedFormats: api.CldAPIArray{"jpeg", "jpg", "png"},
		PublicID:       imagePublicId,
		Transformation: "h_512,w_512,b_white,c_fill",
	})
	if err != nil {
		return mementor_back.Image{}, err
	}

	imagePublicId = uuid.New().String()

	smallImage, err := cld.Upload.Upload(ctx, image.Base64, uploader.UploadParams{
		AllowedFormats: api.CldAPIArray{"jpeg", "jpg", "png"},
		PublicID:       imagePublicId,
		Transformation: "h_144,w_144,b_white,c_fill",
	})
	if err != nil {
		return mementor_back.Image{}, err
	}

	resp.BigImage = bigImage.SecureURL
	resp.SmallImage = smallImage.SecureURL

	image.Image = resp
	err = i.repo.NewImage(ctx, image)
	if err != nil {
		return mementor_back.Image{}, err
	}

	return resp, nil
}

func NewImageService(repo repository.Image) *ImageService {
	return &ImageService{repo: repo}
}
