package handlers

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/zishone/go-notes-service/internal/platform/helpers"

	"github.com/zishone/go-notes-service/internal/notes"
)

// FetchNotes : Handles GET /notes call
func FetchNotes(w http.ResponseWriter, r *http.Request) {
	errs := []error{}

	notes, fails, err := notes.FetchNotes()
	if err != nil {
		errs = append(errs, err)
	}

	if len(fails) != 0 {
		helpers.FailResponse(fails).Send(w)
	}
	if len(errs) != 0 {
		helpers.ErrorResponse(errs).Send(w)
	}
	helpers.SuccessResponse(notes).WithMeta(len(notes)).Send(w)
}

// AddNote : Handles POST /notes call
func AddNote(w http.ResponseWriter, r *http.Request) {
	errs := []error{}
	note := notes.Note{}
	jsoniter.NewDecoder(r.Body).Decode(&note)

	note, fails, err := notes.AddNote(note)
	if err != nil {
		errs = append(errs, err)
	}

	if len(fails) != 0 {
		helpers.FailResponse(fails).Send(w)
	}
	if len(errs) != 0 {
		helpers.ErrorResponse(errs).Send(w)
	}
	helpers.SuccessResponse(note).Send(w)
}
