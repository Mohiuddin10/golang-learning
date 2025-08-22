package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"strings"
)

type Todo struct {
	Todo string `json:"todo"`
}

func New(todo string) (*Todo, error) {
	if todo == "" {
		return nil, errors.New("todo cannot be empty")
	}
	return &Todo{
		Todo: todo,
	}, nil
}

func (c Todo) DisplayTodo() {
	fmt.Println(c.Todo)
}

func (c *Todo) SaveData() error {
	fileName := strings.ReplaceAll(c.Todo, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
