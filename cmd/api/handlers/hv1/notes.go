package hv1

import (
	"net/http"

	"github.com/go-chi/chi"
	jsoniter "github.com/json-iterator/go"
	"github.com/zishone/go-notes-service/internal/notes"
	"github.com/zishone/go-notes-service/internal/platform/helpers"
)

// FetchNotes : Handles GET /v1/notes call
func FetchNotes(w http.ResponseWriter, r *http.Request) {
	errs := []error{}

	notes, fails, err := notes.FetchNotes()
	if err != nil {
		errs = append(errs, err)
	}

	helpers.NewResponse(notes, fails, errs).WithResponseMeta(len(notes)).Send(w)
}

// AddNote : Handles POST /v1/notes call
func AddNote(w http.ResponseWriter, r *http.Request) {
	errs := []error{}
	note := notes.Note{}
	jsoniter.NewDecoder(r.Body).Decode(&note)

	note, fails, err := notes.AddNote(note)
	if err != nil {
		errs = append(errs, err)
	}

	helpers.NewResponse(note, fails, errs).Send(w)
}

// FetchNote : Handles GET /v1/notes/{title} call
func FetchNote(w http.ResponseWriter, r *http.Request) {
	errs := []error{}
	title := chi.URLParam(r, "title")

	note, fails, err := notes.FetchNote(title)
	if err != nil {
		errs = append(errs, err)
	}

	helpers.NewResponse(note, fails, errs).Send(w)
}

// UpdateNote : Handles PUT /v1/notes/{title} call
func UpdateNote(w http.ResponseWriter, r *http.Request) {
	errs := []error{}
	title := chi.URLParam(r, "title")
	note := notes.Note{}
	jsoniter.NewDecoder(r.Body).Decode(&note)

	note, fails, err := notes.UpdateNote(title, note)
	if err != nil {
		errs = append(errs, err)
	}

	helpers.NewResponse(note, fails, errs).Send(w)
}

// DeleteNote : Handles DELETE /v1/notes/{title} call
func DeleteNote(w http.ResponseWriter, r *http.Request) {
	errs := []error{}
	title := chi.URLParam(r, "title")

	note, fails, err := notes.DeleteNote(title)
	if err != nil {
		errs = append(errs, err)
	}

	helpers.NewResponse(note, fails, errs).Send(w)
}
