package notes

import "github.com/zishone/go-notes-service/internal/platform/helpers"

var notes = Notes{
	Note{
		Title: "First",
		Body:  "Hello, World!",
	},
	Note{
		Title: "Second",
		Body:  "Sup, Pinas!",
	},
}

// FetchNotes : Returns list of notes
func FetchNotes() (Notes, helpers.Fails, error) {
	fails := helpers.Fails{}
	return notes, fails, nil
}

// AddNote : Adds a note
func AddNote(note Note) (Note, helpers.Fails, error) {
	fails := helpers.Fails{}
	notes = append(notes, note)
	return note, fails, nil
}
