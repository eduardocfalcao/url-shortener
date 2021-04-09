package routes

import (
	"net/http"

	"github.com/eduardocfalcao/url-shortener/src/api/handlers"
	"github.com/eduardocfalcao/url-shortener/src/api/middlewares"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterRoutes(container *handlers.HandlersContainer) http.Handler {

	r := mux.NewRouter()

	r.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/healthcheck", container.HealthcheckHandler.Healthcheck)

	r.HandleFunc("/shorturl", container.ShortUrlHandler.Create).Methods("POST")

	m := middlewares.NewApiRateLimiter(1, 1)
	r.HandleFunc("/short/{shorturl}", m.Limit(container.ShortUrlHandler.Redirect))

	return r
}
