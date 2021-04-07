package routes

import (
	"net/http"

	"github.com/eduardocfalcao/url-shortener/src/api/handlers"
)

func RegisterRoutes() http.Handler {
	hcHandler := handlers.HealthcheckHandler{}
	mux := http.NewServeMux()

	mux.HandleFunc("/healthcheck", hcHandler.Healthcheck)

	return mux
}
