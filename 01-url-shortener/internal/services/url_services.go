package services

import "url-shortener/internal/repository"

type UrlServices struct {
	UrlRepository *repository.UrlRepository
}

func (s *UrlServices) CreateUrl(userId int, url string) (string, error) {
	return s.UrlRepository.CreateUrl(userId, url)
}
