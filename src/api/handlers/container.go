package handlers

import (
	"github.com/eduardocfalcao/url-shortener/src/api/cache"
	"github.com/eduardocfalcao/url-shortener/src/api/config"
	"github.com/eduardocfalcao/url-shortener/src/api/database"
	"github.com/eduardocfalcao/url-shortener/src/api/repositories"
	"github.com/go-redis/redis"
)

type HandlersContainer struct {
	HealthcheckHandler
	ShortUrlHandler
}

func NewHandlersContainer(appConfig config.AppConfig) (*HandlersContainer, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // host:port of the redis server
		Password: "",               // no password set
		DB:       0,                // use default DB
	})
	c := cache.New(redisClient)
	db, err := database.NewConnection(appConfig.ConnectionString)
	shorturlRepository := repositories.NewShortUrlRepository(db)
	cachedShorturlRepository := repositories.NewCachedShorturlRepository(shorturlRepository, c)

	if err != nil {
		return nil, err
	}

	return &HandlersContainer{
		HealthcheckHandler{},
		ShortUrlHandler{cachedShorturlRepository},
	}, err
}
