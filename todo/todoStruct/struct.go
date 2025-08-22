package todoStruct

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Todo struct {
	Title     string
	CreatedAt time.Time
}

func New(title string) (*Todo, error) {
	if title == "" {
		return nil, errors.New("title should not be empty")
	}
	return &Todo{
		Title:     title,
		CreatedAt: time.Now(),
	}, nil
}

func (t *Todo) DisplayTodo() {
	fmt.Printf("you have %v task in your list time: %v \n", t.Title, t.CreatedAt)
}

func (t *Todo) SaveTodo() error {
	fileName := strings.ReplaceAll(t.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
