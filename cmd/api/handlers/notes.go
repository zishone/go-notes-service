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

	data, err := notes.FetchNotes()
	if err != nil {
		errs = append(errs, err)
	}

	if len(errs) != 0 {
		helpers.Error(errs).Send(w)
	}
	helpers.Success(data).WithMeta("meta").Send(w)
}

// AddNote : Handles POST /notes call
func AddNote(w http.ResponseWriter, r *http.Request) {
	errs := []error{}
	n := notes.Note{}
	jsoniter.NewDecoder(r.Body).Decode(&n)

	data, err := notes.AddNote(n)
	if err != nil {
		errs = append(errs, err)
	}

	if len(errs) != 0 {
		helpers.Error(errs).Send(w)
	}
	helpers.Success(data).Send(w)
}
