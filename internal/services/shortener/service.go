package shortener

import (
	"github.com/phpCoder88/golang-test/internal/entities"
	"github.com/phpCoder88/golang-test/internal/repositories/shortener"
)

type Service struct {
	repo shortener.ShortURLRepository
}

func NewService(repo shortener.ShortURLRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) FindAll() ([]entities.ShortURL, error) {
	return s.repo.FindAll()
}
