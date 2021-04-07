package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/eduardocfalcao/url-shortener/src/api/config"
	"github.com/eduardocfalcao/url-shortener/src/api/routes"
)

func StartHttpServer(address string, config config.AppConfig) {
	handler := routes.RegisterRoutes()
	server := &http.Server{
		Addr:         address,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		log.Printf("Starting server on the http server on port %s", address)
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
			cancel()
		}
	}()

	select {
	case <-c:
		log.Print("Stopping server...")
		server.Shutdown(ctx)
		os.Exit(0)
	case <-ctx.Done():
	}
}
