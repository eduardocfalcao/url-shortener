package routes

import (
	"net/http"

	"github.com/eduardocfalcao/url-shortener/src/api/handlers"
	"github.com/eduardocfalcao/url-shortener/src/api/middlewares"
)

func RegisterRoutes(container *handlers.HandlersContainer) http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/healthcheck", container.HealthcheckHandler.Healthcheck)
	mux.HandleFunc("/shorturl", middlewares.EnsurePost(container.ShortUrlHandler.Create))

	return mux
}
