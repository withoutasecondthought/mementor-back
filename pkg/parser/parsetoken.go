package parser

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"mementor-back/pkg/service"
)

func ParseToken(accessToken string, signingKey []byte) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &service.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*service.Claims); ok && token.Valid {
		return claims.UserId, nil
	}

	return "", errors.New("invalid Token Access")
}
