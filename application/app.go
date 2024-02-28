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
	"github.com/jinzhu/gorm"
	"github.com/rise8-us/neverl8/cli"
	"github.com/rise8-us/neverl8/model"
	"github.com/rise8-us/neverl8/repository"
	"github.com/rise8-us/neverl8/service"
)

type App struct {
	router *chi.Mux
	db     *gorm.DB
}

func New() *App {
	app := &App{
		router: chi.NewRouter(),
	}

	app.loadRoutes()

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: a.router,
	}

	// Connect to the database
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=mydatabase password=password sslmode=disable")
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %v", err)
	}
	defer db.Close()

	a.db = db

	// Automatically create or update database tables based on struct definitions
	if err := db.AutoMigrate(&model.Meeting{}).Error; err != nil {
		return fmt.Errorf("failed to auto migrate database: %v", err)
	}

	// Channel to signal server startup
	serverStarted := make(chan struct{})

	// Start server
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	// Wait for the server to start listening
	go func() {
		<-serverStarted
		// Initialize your repository, service, and CLI
		meetingRepo := repository.NewMeetingRepository(a.db)
		meetingService := service.NewMeetingService(meetingRepo)
		cli := cli.NewCLI(meetingService)

		// Call the CreateMeetingFromCLI method
		cli.CreateMeetingFromCLI()
	}()

	fmt.Println("Server is running on port 8080")

	// Notify server startup
	close(serverStarted)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown error: %v", err)
	}

	fmt.Println("Server stopped gracefully")

	return nil
}
