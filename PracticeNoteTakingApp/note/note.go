package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display() {
	fmt.Printf("Title: %s, Content: %s, CreatedAt: %s\n", note.Title, note.Content, note.CreatedAt.Format(time.RFC3339))
}

func (note Note) Save() error {
	filename := strings.ReplaceAll(note.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"
	//Convert data to json
	//This package only converts to json the publicly available data
	jsonText, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, jsonText, 0644)
}

func New(title, content string) (Note, error) {

	//Do validation
	if title == "" || content == "" {
		return Note{}, errors.New("Empty note")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
