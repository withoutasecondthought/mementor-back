package parser

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	mementor_back "mementor-back"
)

func ParseToken(accessToken string, signingKey []byte) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &mementor_back.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*mementor_back.Claims); ok && token.Valid {
		return claims.UserID, nil
	}

	return "", errors.New("invalid Token Access")
}
