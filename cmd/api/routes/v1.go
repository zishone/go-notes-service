package routes

import (
	"github.com/go-chi/chi"
	"github.com/zishone/go-notes-service/cmd/api/handlers"
)

func v1(r chi.Router) {
	r.Route("/notes", func(r chi.Router) {
		r.Get("/", handlers.FetchNotesV1)
		r.Post("/", handlers.AddNoteV1)
		// r.Put("/", handlers.UpdateNoteV1)
		// r.Delete("/", handlers.DeleteNoteV1)
	})
}
