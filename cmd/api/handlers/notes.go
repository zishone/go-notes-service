package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	jsoniter "github.com/json-iterator/go"
	"github.com/zishone/go-notes-service/internal/notes"
	"github.com/zishone/go-notes-service/internal/platform/helpers"
)

// FetchNotesV1 : Handles GET /v1/notes call
func FetchNotesV1(w http.ResponseWriter, r *http.Request) {
	errs := []error{}

	notes, fails, err := notes.FetchNotes()
	if err != nil {
		errs = append(errs, err)
	}

	helpers.NewResponse(notes, fails, errs).WithMeta(len(notes)).Send(w)
}

// FetchNoteV1 : Handles GET /v1/notes/{title} call
func FetchNoteV1(w http.ResponseWriter, r *http.Request) {
	errs := []error{}
	title := chi.URLParam(r, "title")

	note, fails, err := notes.FetchNote(title)
	if err != nil {
		errs = append(errs, err)
	}

	helpers.NewResponse(note, fails, errs).Send(w)
}

// AddNoteV1 : Handles POST /v1/notes call
func AddNoteV1(w http.ResponseWriter, r *http.Request) {
	errs := []error{}
	note := notes.Note{}
	jsoniter.NewDecoder(r.Body).Decode(&note)

	note, fails, err := notes.AddNote(note)
	if err != nil {
		errs = append(errs, err)
	}

	helpers.NewResponse(note, fails, errs).Send(w)
}
