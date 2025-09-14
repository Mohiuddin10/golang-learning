package main

import "fmt"

func makeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func transformNumber(x *[]int, f func(int) int) int {
	if (*x)[0] > 10 {
		doubleNumber := []int{}
		for _, v := range *x {
			doubleNumber = append(doubleNumber, f(v))
		}
		return doubleNumber
	}
	return nil
}

func double(x int) int {
	return x * 2
}

func main() {

	numbers := []int{11, 14, 17}
	dNumbers := transformNumber(&numbers, double)
	fmt.Println(dNumbers)
	add10 := makeAdder(20)
	fmt.Println(add10(10))

	add20 := makeAdder(30)
	fmt.Println(add20(20))

	add50 := makeAdder(50)
	fmt.Println(add50(50))

}
