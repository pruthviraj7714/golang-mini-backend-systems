package repository

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"url-shortener/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UrlRepository struct {
	DB *gorm.DB
}

func generateShortUrl(url string) string {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	urlHash := sha256.New()

	urlHash.Write([]byte(url))

	shortUrlCode := ""

	sum := urlHash.Sum(nil)

	for i := 0; i < 6; i++ {
		shortUrlCode += string(characters[sum[i]%byte(len(characters))])
	}

	return shortUrlCode
}

func (r *UrlRepository) CreateUrl(userId uuid.UUID, url string) (string, error) {
	shortUrl := generateShortUrl(url)

	var alreadyExistedUrl models.Url

	resp := r.DB.Where("long_url = ?", url).Find(&alreadyExistedUrl)
	if resp.Error != nil {
		if resp.Error == gorm.ErrRecordNotFound {
			//continue
		} else {
			return "", resp.Error
		}
	}

	if resp.RowsAffected > 0 {
		return "http://localhost:8080/" + alreadyExistedUrl.ShortURL, nil
	}

	res := r.DB.Create(&models.Url{
		ShortURL: shortUrl,
		LongURL:  url,
		UserID:   userId,
	})

	if res.Error != nil {
		return "", res.Error
	}

	return "http://localhost:8080/" + shortUrl, nil
}

func (r *UrlRepository) RedirectShortUrl(shortUrlCode string) (string, error) {
	var reqUrl models.Url

	res := r.DB.First(&reqUrl, "short_url = ?", shortUrlCode)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return "", fmt.Errorf("url not found")
	}

	return reqUrl.LongURL, nil
}
