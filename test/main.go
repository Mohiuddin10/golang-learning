package main

import "fmt"

type Operation func(int, int) int
type plusTwo func(int) int

// a func that return a func that adds 2 to all numbers of the array

func addTwo() plusTwo {
	return func(n int) int {
		return n + 2
	}
}

func applyToAll(nums *[]int, plusTwo plusTwo) {
	plusNumbers := []int{}
	for _, value := range *nums {
		plusNumbers = append(plusNumbers, plusTwo(value))
	}
	fmt.Println(plusNumbers)
}

func add(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func multiplay(a, b int) int {
	return a * b
}
func main() {
	numbers := []int{10, 2, 3, 4}
	applyToAll(&numbers, addTwo())

	f := Operation(add)

	fmt.Println("10 + 5 =", f(10, 5))

	g := Operation(minus)

	fmt.Println("10 - 5 =", g(10, 5))
	h := Operation(multiplay)
	fmt.Println("10 * 5 =", h(10, 5))

}
