package handlers

import (
	"net/http"

	"github.com/zishone/go-notes-service/internal/notes"
)

// GetNotes : Handles GET /notes call
func GetNotes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(notes.GetNotes()))
}
