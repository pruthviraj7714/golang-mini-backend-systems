package repository

import (
	"errors"
	"fmt"
	"url-shortener/internal/auth"
	"url-shortener/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) RegisterUser(email, password string) error {

	var existing models.User

	res := r.DB.Where("email = ?", email).First(&existing)

	if res.Error == nil {
		return errors.New("user already exists")
	}

	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return errors.New("database error")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return errors.New("error hashing password")
	}

	res = r.DB.Create(&models.User{
		Email:    email,
		Password: string(hashedPassword),
	})

	fmt.Println(res)

	if res.Error != nil {
		return errors.New("error creating user")
	}

	return nil
}

func (r *UserRepository) LoginUser(email, password string) (string, error) {
	var user models.User

	res := r.DB.Where("email = ?", email).First(&user)

	if res.Error != nil {
		return "", errors.New("user not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	token, err := auth.GenerateToken(user.Id.String())

	if err != nil {
		return "", errors.New("error generating token")
	}

	return token, nil
}
