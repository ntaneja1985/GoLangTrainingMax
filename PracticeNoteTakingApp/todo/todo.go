package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	filename := "todo.json"
	//Convert data to json
	//This package only converts to json the publicly available data
	jsonText, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, jsonText, 0644)
}

func New(content string) (Todo, error) {

	//Do validation
	if content == "" {
		return Todo{}, errors.New("Empty todo")
	}

	return Todo{
		Text: content,
	}, nil
}
