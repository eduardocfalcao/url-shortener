package cache

import (
	"errors"

	"github.com/eduardocfalcao/url-shortener/src/api/interfaces"
	"github.com/go-redis/redis"
)

var NotFoundErr = errors.New("Not in cache")

type cache_impl struct {
	redis *redis.Client
}

func New(r *redis.Client) interfaces.Cache {
	return &cache_impl{r}
}

func (c *cache_impl) HSet(key string, field string, value interface{}) error {
	_, err := c.redis.HSet(key, field, value).Result()
	return err
}

func (c *cache_impl) HGet(key string, field string, returnValue interface{}) error {
	err := c.redis.HGet(key, field).Scan(returnValue)
	if err == redis.Nil {
		return NotFoundErr
	}
	return err
}
