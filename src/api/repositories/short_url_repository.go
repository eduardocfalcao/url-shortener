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

func (r short_url_repository) Create(u entities.ShortUrl) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("Error occured when trying to begin a new transaction. %w", err)
	}
	id := 0
	result := tx.QueryRow(`
			INSERT INTO short_urls (name, shorturl, url) 
			VALUES($1,$2,$3) 
			RETURNING id`,
		u.Name, u.ShortUrl, u.URL)

	if result.Err() != nil {
		return 0, fmt.Errorf("Error occured when trying to create a new short url. %w", result.Err())
	}

	if err = result.Scan(&id); err != nil {
		return 0, fmt.Errorf("Error occured when trying to fetch the id from the new short url. %w", result.Err())
	}

	tx.Commit()

	return id, nil
}

func (r short_url_repository) GetByShorturl(shorturl string) (entities.ShortUrl, error) {
	s := entities.ShortUrl{}
	result := r.db.QueryRow(`
			SELECT id, name, shorturl, url
			FROM short_urls
			WHERE shorturl = $1`,
		shorturl)

	if result.Err() != nil {
		return entities.ShortUrl{}, fmt.Errorf("Error occured when trying to get a short url from database. %w", result.Err())
	}

	if err := result.Scan(&s.ID, &s.Name, &s.ShortUrl, &s.URL); err != nil {
		return entities.ShortUrl{}, fmt.Errorf("Error occured when trying to fetch the short url from database. %w", result.Err())
	}

	return s, nil
}
