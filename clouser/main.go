package main

import "fmt"

const a = 100

var p = 10

func outer() func() {
	money := 100
	age := 30
	fmt.Println("Age:", age)
	show := func() {
		money = money + a + p
		fmt.Println("Money:", money)
	}
	return show
}

func call() {
	incr1 := outer()
	incr1()
	incr1()

	incr2 := outer()
	incr2()
	incr2()
}

func init() {
	fmt.Println("====Bank======")
}

func main() {
	call()

	counter := createCounter() // counter is a closure
	fmt.Println(counter())     // 1
	fmt.Println(counter())     // 2
	fmt.Println(counter())     // 3
}

func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// A closure is a function that remembers the variables from where it was created — even after that scope is gone.

// In short:
// ➡️ A function inside a function that can access outer function’s variables.
