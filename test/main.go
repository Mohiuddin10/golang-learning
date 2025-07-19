package main

import "fmt"

func main() {
	defer fmt.Println("x")
	defer fmt.Println("y")
	fmt.Println("Hello, World!")
}
