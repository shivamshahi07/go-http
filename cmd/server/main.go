package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shivam/http-server/internal/config"
	"github.com/shivam/http-server/internal/handlers"
	"github.com/shivam/http-server/internal/middleware"
)

func main() {
	// Load configuration
	cfg := config.New()

	// Create a new mux router
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/health", handlers.HealthCheckHandler)
	mux.HandleFunc("/api/v1/hello", handlers.HelloHandler)
	mux.HandleFunc("/api/v1/echo", handlers.EchoHandler)

	// Create the server with middleware
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: middleware.RecoveryMiddleware(middleware.LoggingMiddleware(mux)),
	}

	// Channel to listen for errors coming from the server
	serverErrors := make(chan error, 1)

	// Start the server
	go func() {
		log.Printf("Server is starting on port %d", cfg.Port)
		serverErrors <- server.ListenAndServe()
	}()

	// Channel to listen for an interrupt or terminate signal from the OS
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Blocking select waiting for either a server error or a shutdown signal
	select {
	case err := <-serverErrors:
		log.Printf("Error starting server: %v", err)

	case sig := <-shutdown:
		log.Printf("Got signal: %v", sig)
		log.Println("Server is shutting down...")

		// Give outstanding requests a deadline for completion
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.ShutdownTimeout)*time.Second)
		defer cancel()

		// Asking listener to shut down and shed load
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Graceful shutdown did not complete in %v: %v", cfg.ShutdownTimeout, err)
			if err := server.Close(); err != nil {
				log.Printf("Could not stop server: %v", err)
			}
		}
	}
} 