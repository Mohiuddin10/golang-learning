package main

import "fmt"

var b = 10

func call() {
	// b = 30
	fmt.Println(b)
}

func main() {
	b = 20
	call()
	defer fmt.Println("x")
	defer fmt.Println("y")
	fmt.Println("Hello, World!")
}
