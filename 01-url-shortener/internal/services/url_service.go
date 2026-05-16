package services

import (
	"url-shortener/internal/repository"

	"github.com/google/uuid"
)

type UrlService struct {
	UrlRepository *repository.UrlRepository
}

func (s *UrlService) CreateUrl(userId uuid.UUID, url string) (string, error) {
	return s.UrlRepository.CreateUrl(userId, url)
}

func (s *UrlService) RedirectToShortUrl(url string) (string, error) {
	return s.UrlRepository.RedirectShortUrl(url)
}
