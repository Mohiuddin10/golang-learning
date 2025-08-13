package main

import (
	"fmt"

	"example.com/note/note"
)

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	note.DisplayNote(userNote)
}

func getUserData(prompt string) string {
	var value string
	fmt.Println(prompt)
	fmt.Scanln(&value)

	return value
}

func getNoteData() (string, string) {
	title := getUserData("Note title:")
	content := getUserData("Note content:")
	return title, content
}
