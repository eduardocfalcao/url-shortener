package interfaces

import (
	"github.com/eduardocfalcao/url-shortener/src/api/entities"
)

type Cache interface {
	HSet(key string, field string, value interface{}) error
	HGet(key string, field string, returnValue interface{}) error
}

type ShortUrlRepository interface {
	Create(entities.ShortUrl) (int, error)
	GetByShorturl(shorturl string) (entities.ShortUrl, error)
	// Update(int, entities.ShortUrl) error
	// Delete(int) error
	// GetById(int) (entities.ShortUrl, error)
}
