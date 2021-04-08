package repositories

import (
	"github.com/eduardocfalcao/url-shortener/src/api/cache"
	"github.com/eduardocfalcao/url-shortener/src/api/entities"
	"github.com/eduardocfalcao/url-shortener/src/api/interfaces"
)

type cachedShorturlRepository struct {
	inner interfaces.ShortUrlRepository
	cache interfaces.Cache
}

func NewCachedShorturlRepository(inner interfaces.ShortUrlRepository, cache interfaces.Cache) interfaces.ShortUrlRepository {
	return &cachedShorturlRepository{
		inner,
		cache,
	}
}

func (r *cachedShorturlRepository) Create(s entities.ShortUrl) (int, error) {
	id, err := r.inner.Create(s)
	if err != nil {
		return id, err
	}

	return id, r.putInCache(id, s)

}

func (r *cachedShorturlRepository) putInCache(id int, s entities.ShortUrl) error {
	return actOrErr(
		func() error { return r.cache.HSet(s.ShortUrl, "id", s.ID) },
		func() error { return r.cache.HSet(s.ShortUrl, "shorturl", s.ShortUrl) },
		func() error { return r.cache.HSet(s.ShortUrl, "name", s.Name) },
		func() error { return r.cache.HSet(s.ShortUrl, "url", s.URL) },
	)
}

func (r *cachedShorturlRepository) GetByShorturl(shorturl string) (entities.ShortUrl, error) {
	s := entities.ShortUrl{}

	err := actOrErr(
		func() error { return r.cache.HGet(shorturl, "id", &s.ID) },
		func() error { return r.cache.HGet(shorturl, "shorturl", &s.ShortUrl) },
		func() error { return r.cache.HGet(shorturl, "name", &s.Name) },
		func() error { return r.cache.HGet(shorturl, "url", &s.URL) })

	if err == cache.NotFoundErr {
		s, err = r.inner.GetByShorturl(shorturl)
		return s, r.putInCache(s.ID, s)
	}

	return s, err
}

func actOrErr(fn ...func() error) error {
	for _, f := range fn {
		err := f()
		if err != nil {
			return err
		}
	}
	return nil
}
