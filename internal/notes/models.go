package notes

// Note : Represents a note
type Note struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// Notes : Represents list of notes
type Notes []Note
