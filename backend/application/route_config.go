package application

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/middleware"
	"github.com/rise8-us/neverl8/meeting"
)

func (a *App) configureRoutes() {
	meetingController := meeting.NewMeetingController(meeting.NewMeetingService(meeting.NewMeetingRepository(a.db), nil))

	a.router.Use(middleware.Logger)
	a.router.Mount("/api", meetingController.RegisterRoutes())

	// Handle frontend route
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "frontend"))
	a.router.Handle("/", http.FileServer(filesDir))                         // Serve index.html at root
	a.router.Handle("/*", http.StripPrefix("/", http.FileServer(filesDir))) // Serve static files

	a.router.Get("/helloworld", func(w http.ResponseWriter, _ *http.Request) {
		_, writeErr := w.Write([]byte("Hello, World!"))
		if writeErr != nil {
			log.Printf("Error writing response: %v", writeErr)
		}
	})
}
