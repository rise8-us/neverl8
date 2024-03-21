package application

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/middleware"
)

func (a *App) loadRoutes() {
	a.router.Use(middleware.Logger)
	a.router.Get("/helloworld", func(w http.ResponseWriter, _ *http.Request) {
		_, writeErr := w.Write([]byte("Hello, World!"))
		if writeErr != nil {
			log.Printf("Error writing response: %v", writeErr)
		}
	})

	a.router.Get("/api/meetings", a.meetingController.GetAllMeetings)
	a.router.Get("/api/meeting", a.meetingController.GetMeetingByID)
	a.router.Get("/api/meeting/time-slots", a.meetingController.GetAvailableTimeBlocks)
	a.router.Post("/api/meeting/schedule", a.meetingController.UpdateMeetingTime)

	// Handle frontend route
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "frontend"))
	a.router.Handle("/", http.FileServer(filesDir))                         // Serve index.html at root
	a.router.Handle("/*", http.StripPrefix("/", http.FileServer(filesDir))) // Serve static files
}
