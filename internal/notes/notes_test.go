package notes_test

import (
	"testing"

	"github.com/zishone/go-notes-service/internal/notes"
)

func TestGetNotes(t *testing.T) {
	_, err := notes.GetNotes()
	if err != nil {
		t.Errorf("Expected error to be not nil, but got %v", err)
	}
}
