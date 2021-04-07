package cmd

import (
	"log"
	"net/http"
	"time"

	"github.com/eduardocfalcao/url-shortener/src/api/routes"
)

func StartHttpServer(address string) {
	handler := routes.RegisterRoutes()
	server := &http.Server{
		Addr:         address,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Starting the http server on port %s", address)

	log.Fatal(server.ListenAndServe())
}
