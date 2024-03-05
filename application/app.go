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
	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/rise8-us/neverl8/cli"
	"github.com/rise8-us/neverl8/repository"
	"github.com/rise8-us/neverl8/service"
	"gorm.io/driver/postgres"
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

	app.loadRoutes()

	return app
}

const requestTimeout = 5 * time.Second

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:              ":8080",
		Handler:           a.router,
		ReadHeaderTimeout: requestTimeout,
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	// Connect to the database
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbUser, dbPassword, dbName, dbSSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}
	a.db = db

	// Migrate db
	m, err := migrate.New(
		"file://db/migrations",
		fmt.Sprintf(("postgres://%s:%s@%s:%s/%s?sslmode=%s"), dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode))
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		if err.Error() == "no change" {
			log.Println("No migration to run")
		} else {
			log.Fatal(err)
		}
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
		// Initialize repository, service, and CLI
		meetingRepo := repository.NewMeetingRepository(db)
		meetingService := service.NewMeetingService(meetingRepo)
		cliInstance := cli.NewCLI(meetingService)

		// Create a meeting
		cliInstance.CreateMeetingFromCLI()

		// Retrieve all meetings
		cliInstance.GetAllMeetingsFromCLI()
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
