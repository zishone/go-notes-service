package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/zishone/go-notes-service/cmd/api/handlers"
)

// Init : Initializes the api routes
func Init() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/notes", handlers.GetNotes)
	return r
}
