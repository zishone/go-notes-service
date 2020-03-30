package notes_test

import (
	"testing"

	"github.com/zishone/go-notes-service/internal/notes"
)

func TestFetchNotes(t *testing.T) {
	_, fails, err := notes.FetchNotes()
	if err != nil {
		t.Errorf("Expected error to be not nil, but got %v", err)
	}
	if len(fails) != 0 {
		t.Errorf("Expected fails length to be 0, but got %v", len(fails))
	}
}

func TestAddNote(t *testing.T) {
	note := notes.Note{}
	_, fails, err := notes.AddNote(note)
	if err != nil {
		t.Errorf("Expected error to be not nil, but got %v", err)
	}
	if len(fails) != 0 {
		t.Errorf("Expected fails length to be 0, but got %v", len(fails))
	}
}

func TestFetchNote(t *testing.T) {
	title := "First"
	_, fails, err := notes.FetchNote(title)
	if err != nil {
		t.Errorf("Expected error to be not nil, but got %v", err)
	}
	if len(fails) != 0 {
		t.Errorf("Expected fails length to be 0, but got %v", len(fails))
	}
}

func TestUpdateNote(t *testing.T) {
	title := "First"
	note := notes.Note{}
	_, fails, err := notes.UpdateNote(title, note)
	if err != nil {
		t.Errorf("Expected error to be not nil, but got %v", err)
	}
	if len(fails) != 0 {
		t.Errorf("Expected fails length to be 0, but got %v", len(fails))
	}
}
