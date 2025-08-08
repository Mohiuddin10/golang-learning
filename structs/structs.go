package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!

	appUser := User{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}

	outPutUser(&appUser)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}

func outPutUser(u *User) {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}
