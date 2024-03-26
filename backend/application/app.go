package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

type App struct {
	router *chi.Mux
	db     *gorm.DB
}

func New() *App {
	app := &App{
		router: chi.NewRouter(),
	}

	return app
}

const requestTimeout = 5 * time.Second

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:              ":8080",
		Handler:           a.router,
		ReadHeaderTimeout: requestTimeout,
	}

	a.loadEnv()
	a.configureDB()
	a.migrateDBUp()
	a.configureRoutes()

	// Channel to signal server startup
	serverStarted := make(chan struct{})

	// Start server
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	fmt.Println("Server is running on port 8080")

	// Notify server startup
	close(serverStarted)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown error: %w", err)
	}

	fmt.Println("Server stopped gracefully")

	return nil
}
