package main

import "fmt"

func printByValue(numbers [5]int) {
	fmt.Println("Inside printByValue function:", numbers)
}

func printByReference(numbers *[5]int) {
	fmt.Println("Inside printByReference function:", *numbers)
}

func main() {
	x := 20
	addressX := &x
	fmt.Println("x :", x)
	fmt.Println("address of x :", addressX)
	fmt.Println("value at address of x :", *addressX)

	// Changing the value at the address
	fmt.Println("Changing the value at the address...")
	*addressX = 40
	fmt.Println("x :", x)
	fmt.Println("address of x :", addressX)
	fmt.Println("value at address of x :", *addressX)

	// pass by value
	// pass by referance

	var arr = [5]int{1, 2, 3, 4, 5}
	// pass by value
	printByValue(arr)
	// pass by reference
	printByReference(&arr)
}
