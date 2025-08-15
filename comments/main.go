package main

import (
	"fmt"

	"example.com/comment/commentStruct"
	"example.com/comment/getComment"
)

func main() {
	author, err := getComment.GetUserData("Author Name:")
	if err != nil {
		fmt.Println(err)
		return
	}
	comment, err := getComment.GetUserData("Comment:")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Anthou: %v\n Comment: %v\n", author, comment)
	newComment, err := commentStruct.New(author, comment)
	if err != nil {
		fmt.Println(err)
		return
	}
	newComment.DisplayComment()
	newComment.SaveData()
}

// Objectives ==>
// 1. Write a struct for comment
// 2. get input from user // ok
// 3. display comment //ok
// 4. allow user to add more comments
// 5. allow user to see previous comments
