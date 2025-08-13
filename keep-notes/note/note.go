package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("title and content cannot be empty")
	}
	return &Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil

}

func DisplayNote(note *Note) {
	fmt.Println("Title:", note.title)
	fmt.Println("Content:", note.content)
	fmt.Println("Created At:", note.createdAt)
}
