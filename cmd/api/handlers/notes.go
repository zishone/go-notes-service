package handlers

import (
	"net/http"

	"github.com/zishone/go-notes-service/internal/platform/helpers"

	"github.com/zishone/go-notes-service/internal/notes"
)

// GetNotes : Handles GET /notes call
func GetNotes(w http.ResponseWriter, r *http.Request) {
	n, err := notes.GetNotes()
	if err != nil {
		helpers.Error("\"Cant get data\"").Send(w)

	}
	helpers.Success(n).WithMeta("{\"sample\":\"meta\"}").Send(w)
}
