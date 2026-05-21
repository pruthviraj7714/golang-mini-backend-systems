package repository

import (
	"banking-system/internal/models"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) Register(email, password string) error {

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	err = r.DB.Create(&models.User{
		Email:    email,
		Password: string(hashed),
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Login(email, password string) (string, string, error) {

	var user models.User

	err := r.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		return "", "", errors.New("User not found")
	}

	fmt.Println(user.Password)
	fmt.Println(password)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "", "", errors.New("Incorrect password")
	}

	accessToken, _ := GenerateAccessToken(user.ID)
	refreshToken, _ := GenerateRefreshToken(user.ID)

	return accessToken, refreshToken, nil

}
