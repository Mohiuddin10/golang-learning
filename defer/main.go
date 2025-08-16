package main

import "fmt"

func calcNamded() (result int) {
	fmt.Println("calcNamded 1:", result)

	sum := func() {
		result = result + 10
		fmt.Println("calcNamed defer sum:", result)
	}

	defer sum()

	result++
	fmt.Println("calcNamed res last: ", result)

	defer fmt.Println("This is last defer")
	return
}

func calc() {
	var result int
	fmt.Println("calc 1:", result)

	sum := func() {
		result = result + 10
		fmt.Println("calc sum:", result)
	}

	defer sum()

	result++
	fmt.Println("calc res last: ", result)
}

func main() {
	calcNamded()
	calc()
}
