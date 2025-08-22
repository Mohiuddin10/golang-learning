package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/todo/todoStruct"
)

func main() {
	todo, err := getUserData("Add your todo")
	if err != nil {
		fmt.Println(err)
		return
	}

	newTodo, err := todoStruct.New(todo)
	if err != nil {
		fmt.Println(err)
		return
	}
	newTodo.DisplayTodo()
}

func getUserData(prompt string) (string, error) {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text, nil

}
