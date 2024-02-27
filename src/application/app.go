package application

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jinzhu/gorm"
)

type App struct {
	router *chi.Mux
}

func New() *App {
	app := &App{
		router: chi.NewRouter(),
	}

	app.loadRoutes()

	return app
}

func (a *App) Start(ctx context.Context) error {
	// Connect to the database
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=drewfugate dbname=mydatabase password=password sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Automigrate the User model
	db.AutoMigrate(&User{})

	// Initialize chi router
	// r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)

	// Routes
	r.Get("/helloworld", HelloWorldHandler)

	// Start server
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
