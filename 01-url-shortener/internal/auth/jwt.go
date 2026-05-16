package auth

import (
	"time"
	"url-shortener/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(config.LoadConfig().JWTSecret)

func GenerateToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil

}
