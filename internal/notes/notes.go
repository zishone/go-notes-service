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

// FetchNote : Return a note given title
func FetchNote(title string) (Note, helpers.Fails, error) {
	fails := helpers.Fails{}
	for _, n := range notes {
		if n.Title == title {
			return n, fails, nil
		}
	}
	fails = append(fails, helpers.Fail{
		Code:    "NOT_FOUND",
		Message: "Note was not found.",
	})
	return Note{}, fails, nil
}

// UpdateNote : Updates a note given title
func UpdateNote(title string, note Note) (Note, helpers.Fails, error) {
	fails := helpers.Fails{}
	for i, n := range notes {
		if n.Title == title {
			notes[i] = note
			return notes[i], fails, nil
		}
	}
	fails = append(fails, helpers.Fail{
		Code:    "NOT_FOUND",
		Message: "Note was not found.",
	})
	return Note{}, fails, nil
}

// DeleteNote : Deletes a note given title
func DeleteNote(title string) (Note, helpers.Fails, error) {
	fails := helpers.Fails{}
	for i, n := range notes {
		if n.Title == title {
			copy(notes[i:], notes[i+1:])
			notes = notes[:len(notes)-1]
			return n, fails, nil
		}
	}
	fails = append(fails, helpers.Fail{
		Code:    "NOT_FOUND",
		Message: "Note was not found.",
	})
	return Note{}, fails, nil
}
