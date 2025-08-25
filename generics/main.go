package main

import "fmt"

func main() {
	result := add(1.12, 2.32)
	result = result + 10.0
	fmt.Println("result:", result)

	var numbers = [6]int{1, 2, 3, 4, 5, 6}
	num1 := numbers[1:3]
	fmt.Println("Numbers slice:", num1)
	fmt.Println("Numbers length:", len(num1))
	fmt.Println("Numbers capacity:", cap(num1))

	num1 = num1[:4]
	fmt.Println("Numbers slice:", num1)
	fmt.Println("Numbers length:", len(num1))
	fmt.Println("Numbers capacity:", cap(num1))
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
