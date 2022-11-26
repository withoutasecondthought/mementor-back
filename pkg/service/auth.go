package service

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	mementor_back "mementor-back"
	"mementor-back/pkg/repository"
	"time"
)

type AuthService struct {
	repos repository.Authorization
}

type Claims struct {
	jwt.StandardClaims
	UserID string `json:"userId" bson:"userId"`
}

func (a *AuthService) SignIn(ctx context.Context, user mementor_back.Auth) (string, error) {
	pass := user.Password + viper.GetString("password_salt")
	hash, err := hashPassword(user.Password)
	if err != nil {
		return "", err
	}
	user.Password = hash

	resp, err := a.repos.GetUser(ctx, user)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(resp.Password), []byte(pass))
	if err != nil {
		return "", err
	}

	return generateToken(resp.ID.Hex())
}

func (a *AuthService) SignUp(ctx context.Context, user mementor_back.Auth) (string, error) {
	user.ValidProfile = false
	hash, err := hashPassword(user.Password)
	if err != nil {
		return "", err
	}

	user.Password = hash
	id, err := a.repos.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}
	return generateToken(id)
}

func hashPassword(password string) (string, error) {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password+viper.GetString("password_salt")), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(fromPassword), nil
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
