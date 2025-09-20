package main

import "fmt"

func devideByTwo(a int) int {
	const b = 2
	return a / b
}

func main() {

	sum := func(a, b int) int {
		return a + b
	}
	fmt.Println(sum(10, 20))
	minus := func(a, b int) int {
		return a - b
	}(30, 15)
	fmt.Println(minus)

	div := devideByTwo(300)
	fmt.Println(div)
}
