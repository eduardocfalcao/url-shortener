package interfaces

import (
	"github.com/eduardocfalcao/url-shortener/src/api/entities"
)

type ShortUrlRepository interface {
	Create(entities.ShortUrl) (int64, error)
	// Update(int, entities.ShortUrl) error
	// Delete(int) error
	// GetById(int) (entities.ShortUrl, error)
}
