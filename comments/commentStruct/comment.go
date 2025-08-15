package commentStruct

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"strings"
	"time"
)

type Comment struct {
	Author    string
	Comment   string
	CreatedAt time.Time
}

func New(author, comment string) (*Comment, error) {
	if author == "" || comment == "" {
		return nil, errors.New("author and comment cannot be empty")
	}
	return &Comment{
		Author:    author,
		Comment:   comment,
		CreatedAt: time.Now(),
	}, nil
}

func (c *Comment) DisplayComment() {
	fmt.Printf("Author Name: %v have the following comments: \n\n%v\n", c.Author, c.Comment)
}

func (c *Comment) SaveData() error {
	fileName := strings.ReplaceAll(c.Author, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
