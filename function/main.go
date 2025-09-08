// use function as value
// anonymous function
// recursion function
// variadic function

package main

import "fmt"

type multyplayByO func(int, int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	double := doubleNumbers(&numbers)
	fmt.Println(double)
	fmt.Println(numbers)

	addTen := plusTen(&numbers)
	fmt.Println(addTen)
	fmt.Println(numbers)

	// get func as value
	multyplay := transformNumber(&numbers, multiplyBy)
	fmt.Println("multy", multyplay)

	// take a number that multyplay by the element of {}

	multyplayBy := multiplayOp(&numbers, multiPlayOperation, 5)
	fmt.Println("multyplayBy", multyplayBy)

}

func doubleNumbers(numbers *[]int) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, value*2)
	}
	return dNumbers
}

func plusTen(numbers *[]int) []int {
	plusTenNumbers := []int{}
	for _, value := range *numbers {
		plusTenNumbers = append(plusTenNumbers, value+10)
	}
	return plusTenNumbers
}

func transformNumber(numbers *[]int, f func(int) int) []int {
	newNumbers := []int{}
	for _, v := range *numbers {
		newNumbers = append(newNumbers, f(v))
	}
	return newNumbers
}

func multiplyBy(n int) int {
	return n * 3
}

func multiplayOp(numbers *[]int, multyplayByO multyplayByO, mNum int) []int {
	newNumbers := []int{}
	for _, v := range *numbers {
		newNumbers = append(newNumbers, multyplayByO(v, mNum))
	}
	return newNumbers
}

func multiPlayOperation(n, m int) int {
	return n * m
}
