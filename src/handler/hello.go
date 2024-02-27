package handler

import (
	"net/http"

	hello "github.com/drewfugate/neverl8/repository"
)

type Hello struct {
	Repo *hello.PostgresRepo
}

func (h *Hello) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	_, err := w.Write([]byte("Hello, World!"))
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
