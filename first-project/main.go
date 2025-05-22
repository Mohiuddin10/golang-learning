package main

import "fmt"

// singleoutput with return
func add(a int, b int) int {
	sum := a + b
	return sum
}

// multiple output
func getNumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	mul := num1 * num2
	return sum, mul
}

func printAnything() {
	fmt.Println("Hello there!")
}

func printGreetings(name string) {
	fmt.Println("Hello there, welcome to Go programming", name)
}

func main() {
	num1 := 5
	num2 := 10
	sum := add(num1, num2)
	fmt.Println(sum)

	p, q := getNumbers(num1, num2)
	fmt.Println(`p:`, p)
	fmt.Println(`q:`, q)
	sum2 := add(20, 20)
	fmt.Println(sum2)
	a, b := getNumbers(10, num2)
	fmt.Println(`a:`, a)
	fmt.Println(`b:`, b)
	printAnything()

	printGreetings("Mohammd Mohiuddin")
}
