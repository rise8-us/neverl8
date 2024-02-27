package application

import (
	"github.com/drewfugate/neverl8/handler"
	hello "github.com/drewfugate/neverl8/repository"

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

	a.router = router
}
