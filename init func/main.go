package main

import "fmt"

var a = 10

func main() {
	fmt.Println("a = ", a)
	// Anonymous function IIFE
	func(a int, b int) {
		c := a + b
		fmt.Println("c = ", c)
	}(10, 20)
}

func init() {
	// it will execute first and before main
	fmt.Println("first print from init func", a)
	a = 50

}
