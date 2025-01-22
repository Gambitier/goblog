//go:generate swag init -o ./_apidocs

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gambitier/goblog/startup/context"
	httpserver "github.com/gambitier/goblog/startup/server/http"
)

// @title Go Blog API
// @version 1.0
// @description This is a server for Go Blog API.
// @host localhost:3000
// @BasePath /
func main() {
	appCtx, err := context.NewAppContext()
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}

	httpServer := httpserver.NewHttpServer(appCtx)
	httpServer.Configure()
	httpServer.RegisterRoutes()

	// Channel to listen for errors
	errChan := make(chan error, 1)

	// Start Fiber server in a goroutine
	go func() {
		if err := httpServer.RunServer(3000); err != nil {
			errChan <- fmt.Errorf("failed to start Fiber server: %w", err)
		}
	}()

	// Channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errChan:
		log.Fatalf("server error: %v", err)
	case <-quit:
		log.Println("shutting down servers...")

		// Gracefully shut down the Fiber server
		if err := httpServer.Shutdown(); err != nil {
			log.Fatalf("Fiber server shutdown error: %v", err)
		}

		log.Println("servers shut down gracefully")
	}
}
