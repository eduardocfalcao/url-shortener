package repositories

import (
	"database/sql"
	"fmt"

	"github.com/eduardocfalcao/url-shortener/src/api/entities"
	"github.com/eduardocfalcao/url-shortener/src/api/interfaces"
)

type short_url_repository struct {
	db *sql.DB
}

func NewShortUrlRepository(db *sql.DB) interfaces.ShortUrlRepository {
	return &short_url_repository{db}
}

func (r short_url_repository) Create(u entities.ShortUrl) (int64, error) {
	tx, _ := r.db.Begin()
	result, err := tx.Exec("INSERT INTO short_urls (name, shorturl, url) VALUES(?,?,?)", u.Name, u.ShortUrl, u.URL)
	if err != nil {
		return 0, fmt.Errorf("Error occured when trying to create a new short url. %w", err)
	}
	tx.Commit()

	return result.LastInsertId()
}
