package routes

import (
	"github.com/go-chi/chi"
	"github.com/zishone/go-notes-service/cmd/api/handlers"
)

func v1(r chi.Router) {
	r.Route("/notes", func(r chi.Router) {
		r.Get("/", handlers.FetchNotesV1)
		r.Post("/", handlers.AddNoteV1)
		r.Get("/{title}", handlers.FetchNoteV1)
		// r.Put("/{title}", handlers.UpdateNoteV1)
		// r.Delete("/{title}", handlers.DeleteNoteV1)
	})
}
