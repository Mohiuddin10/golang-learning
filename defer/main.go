package main

import (
	"fmt"
)

func first() int {
	c := 3 + 5
	sum := func() int {
		newResult := 2 + c
		return newResult
	}
	return sum()
}

func calcNamded() (result int) {
	fmt.Println("calcNamded 1:", result)

	sum := func() {
		result = result + 10
		fmt.Println("calcNamed defer sum:", result)
	}

	defer sum()

	result = 5
	p := func(a int) {
		fmt.Println("ami", a)
	}
	defer p(result)
	defer fmt.Println(result)
	fmt.Println("calcNamed res last: ", result)

	defer fmt.Println("This is last defer")
	defer fmt.Println(5)
	return
}

// 0,5,5,this is, 5, ami 5, 15, 15

// ===> defer list pointer

func calc() int {
	var result int
	fmt.Println("calc 1:", result)

	sum := func() {
		result = result + 10
		fmt.Println("calc sum:", result)
	}

	defer sum()

	result = 5
	fmt.Println("calc res last: ", result)
	return result
}

func main() {
	a := calcNamded()
	fmt.Println("named return end: ", a)
	// b := calc()
	// fmt.Println("b =", b)

	// fmt.Println(add1)
}
