package handlers

import (
	"github.com/eduardocfalcao/url-shortener/src/api/config"
	"github.com/eduardocfalcao/url-shortener/src/api/database"
	"github.com/eduardocfalcao/url-shortener/src/api/repositories"
)

type HandlersContainer struct {
	HealthcheckHandler
	ShortUrlHandler
}

func NewHandlersContainer(appConfig config.AppConfig) (*HandlersContainer, error) {
	db, err := database.NewConnection(appConfig.ConnectionString)
	shorturlRepository := repositories.NewShortUrlRepository(db)

	if err != nil {
		return nil, err
	}

	return &HandlersContainer{
		HealthcheckHandler{},
		ShortUrlHandler{shorturlRepository},
	}, err
}
