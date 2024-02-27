package application

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (a *App) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", a.loadOrderRoutes)

	a.router = router
}

func (a *App) loadOrderRoutes(router chi.Router) {
	helloWorldHandler := &handler.Hello{
		Repo: &handler.RedisRepo{
			Client: a.rdb,
		},
	}

	// router.Post("/", orderHandler.Create)
	router.Get("/helloworld", helloWorldHandler)
}
