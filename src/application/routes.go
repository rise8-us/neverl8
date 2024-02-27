package application

import (
	"github.com/drewfugate/neverl8/cli"
	"github.com/drewfugate/neverl8/handler"
	hello "github.com/drewfugate/neverl8/repository"
	"github.com/drewfugate/neverl8/service"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (a *App) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	hello := &handler.Hello{
		Repo: &hello.PostgresRepo{
			DB: a.db,
		},
	}

	router.Get("/helloworld", hello.HelloWorldHandler)

	// Initialize your repository, service, and CLI
	meetingRepo := meetingRepo.NewMeetingRepository(db) // Assuming you have initialized your DB
	meetingService := service.NewMeetingService(meetingRepo)
	cli := cli.NewCLI(meetingService)

	// Call the CreateMeetingFromCLI method
	cli.CreateMeetingFromCLI()
	a.router = router
}
