package routes

import (
	"net/http"

	"github.com/eduardocfalcao/url-shortener/src/api/handlers"
	"github.com/eduardocfalcao/url-shortener/src/api/middlewares"
	"github.com/gorilla/mux"
)

func RegisterRoutes(container *handlers.HandlersContainer) http.Handler {

	r := mux.NewRouter()

	r.HandleFunc("/healthcheck", container.HealthcheckHandler.Healthcheck)
	r.HandleFunc("/shorturl", middlewares.EnsurePost(container.ShortUrlHandler.Create))
	r.HandleFunc("/short/{shorturl}", container.ShortUrlHandler.Redirect)

	return r
}
