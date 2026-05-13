package auth

import (
	"time"
	"url-shortener/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId string) (string, error) {

	var jwtKey = []byte(config.LoadConfig().JWTSecret)

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}
