package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTService struct {
	AccessSecret  string
	RefreshSecret string
}

func NewJWTService(accessSecret, refreshSecret string) *JWTService {
	return &JWTService{
		AccessSecret:  accessSecret,
		RefreshSecret: refreshSecret,
	}
}

func (j *JWTService) GenerateAccessToken(userId uuid.UUID) (string, error) {

	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.AccessSecret))

}

func (j *JWTService) GenerateRefreshToken(userId uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.RefreshSecret))
}
