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

	data, fails, err := notes.FetchNotes()
	if err != nil {
		errs = append(errs, err)
	}

	if len(fails) != 0 {
		helpers.FailResponse(fails).Send(w)
	}
	if len(errs) != 0 {
		helpers.ErrorResponse(errs).Send(w)
	}
	helpers.SuccessResponse(data).WithMeta("meta").Send(w)
}

// AddNote : Handles POST /notes call
func AddNote(w http.ResponseWriter, r *http.Request) {
	errs := []error{}
	n := notes.Note{}
	jsoniter.NewDecoder(r.Body).Decode(&n)

	data, fails, err := notes.AddNote(n)
	if err != nil {
		errs = append(errs, err)
	}

	if len(fails) != 0 {
		helpers.FailResponse(fails).Send(w)
	}
	if len(errs) != 0 {
		helpers.ErrorResponse(errs).Send(w)
	}
	helpers.SuccessResponse(data).Send(w)
}
