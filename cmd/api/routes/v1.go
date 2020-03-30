package routes

import (
	"github.com/go-chi/chi"
	"github.com/zishone/go-notes-service/cmd/api/handlers"
)

func v1(r chi.Router) {
	r.Route("/notes", func(r chi.Router) {
		r.Get("/", handlers.FetchNotes)
		r.Post("/", handlers.AddNote)
		// r.Put("/", handlers.UpdateNote)
		// r.Delete("/", handlers.DeleteNote)
	})
}
