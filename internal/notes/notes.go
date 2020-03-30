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
	return notes, nil, nil
}

// AddNote : Adds a note
func AddNote(note Note) (Note, helpers.Fails, error) {
	notes = append(notes, note)
	return note, nil, nil
}
