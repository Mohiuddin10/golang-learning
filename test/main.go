package main

import "fmt"

func main() {
	number1 := 32
	fmt.Println(number1)
	add(&number1, 10)
	fmt.Println(number1)
}

func add(x *int, y int) {
	*x = *x + y
	fmt.Println("The result is:", *x)
}
