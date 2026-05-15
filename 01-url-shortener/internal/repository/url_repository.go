package repository

import (
	"crypto/sha256"
	"encoding/hex"
	"url-shortener/internal/models"

	"gorm.io/gorm"
)

type UrlRepository struct {
	DB *gorm.DB
}

func generateShortUrl(url string) string {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	urlHash := sha256.New()

	urlHash.Write([]byte(url))

	hash := hex.EncodeToString(urlHash.Sum(nil))

	shortUrl := ""

	for i := 0; i < 6; i++ {
		shortUrl += string(characters[hash[i]%byte(len(characters))])
	}

	return shortUrl
}

func (r *UrlRepository) CreateUrl(userId int, url string) (string, error) {

	shortUrl := generateShortUrl(url)

	res := r.DB.Create(&models.Url{
		ShortURL: shortUrl,
		LongURL:  url,
		UserID:   userId,
	})

	if res.Error != nil {
		return "", res.Error
	}

	return shortUrl, nil
}
