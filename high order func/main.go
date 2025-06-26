package main

import "fmt"

func processOperation(a int, b int, opp func(x int, y int)) {
	opp(a, b)
}

func add(p int, q int) {
	x := p + q
	fmt.Println("Addition: ", x)
}

func main() {
	fmt.Println("Starting the program")

	// Passing function as an argument
	processOperation(10, 20, add)

	// Anonymous function
	processOperation(30, 40, func(x int, y int) {
		fmt.Println("Anonymous Addition: ", x-y)
	})

	fmt.Println("Ending the program")

	sum := add
	sum(50, 60)
}
