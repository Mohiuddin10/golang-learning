package main

import "fmt"

type Operation func(int, int) int

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

	f := Operation(add)

	fmt.Println("10 + 5 =", f(10, 5))

	g := Operation(minus)

	fmt.Println("10 - 5 =", g(10, 5))
	h := Operation(multiplay)
	fmt.Println("10 * 5 =", h(10, 5))

}
