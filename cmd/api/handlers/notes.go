package handlers

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
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

// AddNote : Handles POST /notes call
func AddNote(w http.ResponseWriter, r *http.Request) {
	n := notes.Note{}
	jsoniter.NewDecoder(r.Body).Decode(&n)
	err := notes.AddNote(n)
	if err != nil {
		helpers.Error([]string{"\"Cant add data\""}).Send(w)
	}
	helpers.Success(n).WithMeta(sample{Sample: "meta"}).Send(w)
}
