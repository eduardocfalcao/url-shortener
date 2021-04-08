package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/eduardocfalcao/url-shortener/src/api/config"
	"github.com/eduardocfalcao/url-shortener/src/api/handlers"
	"github.com/eduardocfalcao/url-shortener/src/api/routes"
)

func StartHttpServer(address string, config config.AppConfig) {
	container, err := handlers.NewHandlersContainer(config)
	if err != nil {
		log.Fatalf("Error when trying to create the handlers container: %s", err.Error())
	}

	handler := routes.RegisterRoutes(container)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.AppPort),
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
			log.Print(err)
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
