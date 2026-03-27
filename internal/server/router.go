package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/neptune2k21/chemin-d-or/internal/tasks"
)

func NewRouter(taskHandler *tasks.Handler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status": "ok",
		})
	})

	r.Route("/tasks", func(r chi.Router) {
		r.Post("/", taskHandler.Create)
		r.Get("/", taskHandler.List)
		r.Get("/{id}", taskHandler.GetByID)
		r.Patch("/{id}", taskHandler.Update)
		r.Delete("/{id}", taskHandler.Delete)
	})

	return r
}
