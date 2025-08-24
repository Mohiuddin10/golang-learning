package main

import "fmt"

func main() {
	result := add(1.12, 2.32)
	result = result + 10.0
	fmt.Println("result:", result)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
