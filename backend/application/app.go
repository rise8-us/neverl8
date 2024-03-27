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
	controller "github.com/rise8-us/neverl8/controller"
	"github.com/rise8-us/neverl8/repository"
	hostSvc "github.com/rise8-us/neverl8/service/host"
	meetingSvc "github.com/rise8-us/neverl8/service/meeting"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	router            *chi.Mux
	db                *gorm.DB
	meetingService    *meetingSvc.MeetingService
	meetingController *controller.MeetingController
	hostService       *hostSvc.HostService
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

	// Load environment variables
	envMap, err := LoadEnvVariables()
	if err != nil {
		return fmt.Errorf("failed to load environment variables: %w", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
		envMap["DB_HOST"], envMap["DB_USER"], envMap["DB_PASSWORD"], envMap["DB_NAME"], envMap["DB_SSLMODE"])

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}
	a.db = db

	// Migrate db
	m, err := migrate.New(
		"file://db/migrations",
		fmt.Sprintf(("postgres://%s:%s@%s:%s/%s?sslmode=%s"),
			envMap["DB_USER"], envMap["DB_PASSWORD"], envMap["DB_HOST"], envMap["DB_PORT"], envMap["DB_NAME"], envMap["DB_SSLMODE"]))
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

	meetingRepo := repository.NewMeetingRepository(a.db)
	a.meetingService = meetingSvc.NewMeetingService(meetingRepo, nil)
	a.meetingController = controller.NewMeetingController(a.meetingService)
	hostRepo := repository.NewHostRepository(a.db)
	a.hostService = hostSvc.NewHostService(hostRepo)

	a.loadRoutes()

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
		meetingService := meetingSvc.NewMeetingService(meetingRepo, nil)
		hostRepo := repository.NewHostRepository(db)
		hostService := hostSvc.NewHostService(hostRepo)
		cliInstance := cli.NewCLI(meetingService, hostService)

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

func LoadEnvVariables() (map[string]string, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	envMap := make(map[string]string)
	envMap["DB_HOST"] = os.Getenv("DB_HOST")
	envMap["DB_PORT"] = os.Getenv("DB_PORT")
	envMap["DB_USER"] = os.Getenv("DB_USER")
	envMap["DB_PASSWORD"] = os.Getenv("DB_PASSWORD")
	envMap["DB_NAME"] = os.Getenv("DB_NAME")
	envMap["DB_SSLMODE"] = os.Getenv("DB_SSLMODE")

	return envMap, nil
}
