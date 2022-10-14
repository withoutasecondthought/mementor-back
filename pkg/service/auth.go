package service

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	mementor_back "mementor-back"
	"mementor-back/pkg/repository"
	"time"
)

type AuthService struct {
	repos repository.Authorization
}

type Claims struct {
	jwt.StandardClaims
	UserId string `json:"userId" bson:"userId"`
}

func (a *AuthService) SignIn(ctx context.Context, user mementor_back.Auth) (string, error) {
	user.Password = hashPassword(user.Password)

	id, err := a.repos.GetUser(ctx, user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New("invalid login or password")
		}
		return "", err
	}

	return generateToken(id)
}

func (a *AuthService) SignUp(ctx context.Context, user mementor_back.Auth) (string, error) {
	user.ValidProfile = false
	user.Password = hashPassword(user.Password)
	id, err := a.repos.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}
	return generateToken(id)
}

func hashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(viper.GetString("password_salt"))))
}

func generateToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		jwt.StandardClaims{
			IssuedAt: time.Now().Unix(),
		},
		id,
	})

	return token.SignedString([]byte(viper.GetString("signing_key")))
}

func NewAuthService(repos repository.Authorization) *AuthService {
	return &AuthService{repos: repos}
}
