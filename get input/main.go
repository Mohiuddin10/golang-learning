package main

import "fmt"

func main() {
	// Print Welcome Message
	fmt.Println("Welcome to Go programming")

	// get name as input
	var name string
	fmt.Println("Enter your name:")
	fmt.Scanln(&name)

	// Print name in greeting message
	fmt.Println("Hello there, welcome to Go programming", name)

	// get year of birth as input
	var yearOfBirth int
	fmt.Println("Enter your year of birth:")
	fmt.Scanln(&yearOfBirth)
	// get current year
	var currentYear int
	fmt.Println("Enter current year:")
	fmt.Scanln(&currentYear)
	// calculate age
	age := currentYear - yearOfBirth
	// Print age
	fmt.Println("Hello ", name, " your age is", age)
}
