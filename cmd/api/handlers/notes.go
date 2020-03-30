package handlers

import (
	"net/http"

	"github.com/zishone/go-notes-service/internal/platform/helpers"

	"github.com/zishone/go-notes-service/internal/notes"
)

type sample struct {
	Sample string `json:"sample"`
}

// FetchNotes : Handles GET /notes call
func FetchNotes(w http.ResponseWriter, r *http.Request) {
	n, err := notes.FetchNotes()
	if err != nil {
		helpers.Error([]string{"\"Cant get data\""}).Send(w)
	}
	helpers.Success(n).WithMeta(sample{Sample: "meta"}).Send(w)
}
