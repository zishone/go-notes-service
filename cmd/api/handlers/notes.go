package handlers

import (
	"net/http"

	"github.com/zishone/go-notes-service/internal/platform/helpers"

	"github.com/zishone/go-notes-service/internal/notes"
)

// FetchNotes : Handles GET /notes call
func FetchNotes(w http.ResponseWriter, r *http.Request) {
	n, err := notes.FetchNotes()
	if err != nil {
		helpers.Error([]string{"\"Cant get data\""}).Send(w)
	}
	helpers.Success(n).WithMeta("{\"sample\":\"meta\"}").Send(w)
}
