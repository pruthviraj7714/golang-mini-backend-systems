package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateAccessToken(userId uuid.UUID) (string, error) {

	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))

}

func GenerateRefreshToken(userId uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("REFRESH_TOKEN_SECRET")))
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
	})
	return token, err
}
