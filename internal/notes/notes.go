package notes

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
func FetchNotes() (Notes, error) {
	return notes, nil
}

// AddNote : Adds a note
func AddNote(note Note) (Note, error) {
	notes = append(notes, note)
	return note, nil
}
