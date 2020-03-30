package routes

import (
	"github.com/go-chi/chi"
	"github.com/zishone/go-notes-service/cmd/api/handlers/hv1"
)

func rv1(r chi.Router) {
	r.Route("/notes", func(r chi.Router) {
		r.Get("/", hv1.FetchNotes)
		r.Post("/", hv1.AddNote)
		r.Get("/{title}", hv1.FetchNote)
		r.Put("/{title}", hv1.UpdateNote)
		r.Delete("/{title}", hv1.DeleteNote)
	})
}
