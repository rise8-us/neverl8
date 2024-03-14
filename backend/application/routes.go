package application

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/middleware"
)

func (a *App) loadRoutes() {
	a.router.Use(middleware.Logger)
	a.router.Get("/helloworld", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	a.router.Get("/api/meetings", a.meetingController.GetAllMeetings)
	a.router.Get("/api/meeting", a.meetingController.GetMeetingByID)

	// Handle my /frontend route
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "./frontend"))
	a.router.Handle("/", http.FileServer(filesDir))                         // Serve index.html at root
	a.router.Handle("/*", http.StripPrefix("/", http.FileServer(filesDir))) // Serve static files

}
