package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Define a model
type User struct {
	gorm.Model
	Name string
}

func main() {
	// Connect to the database
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=drewfugate dbname=mydatabase password=password sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Automigrate the User model
	db.AutoMigrate(&User{})

	// Initialize chi router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)

	// Routes
	r.Get("/helloworld", HelloWorldHandler)

	// Start server
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	_, err := w.Write([]byte("Hello, World!"))
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
