package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

var b = 10

func call() {
	// b = 30
	fmt.Println(b)
}

func add(a, b float64) float64 {
	return a + b
}

func main() {
	b = 20
	call()
	defer fmt.Println("x")
	defer fmt.Println("y")
	fmt.Println("Hello, World!")
	num1 := add(15, 20)
	fmt.Println("Sum:", num1)

	fmt.Println("Random Name:", randomdata.FullName(randomdata.Female))
}
