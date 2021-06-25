package shortener

import "github.com/phpCoder88/golang-test/internal/entities"

type ShortURLRepository interface {
	FindAll() ([]entities.ShortURL, error)
	FindByID(int64) (entities.ShortURL, error)
	Add(entities.ShortURL) (bool, error)
}
