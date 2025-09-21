package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	double := createTransformar(2)
	triple := createTransformar(3)

	doubled := transformNumber(&numbers, double)
	tripled := transformNumber(&numbers, triple)
	fmt.Println("Doubled:", doubled)
	fmt.Println("Tripled:", tripled)

}

func transformNumber(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}
	for _, number := range *numbers {
		dNumbers = append(dNumbers, transform(number))
	}
	return dNumbers
}

func createTransformar(factor int) func(int) int {
	return func(i int) int {
		return i * factor
	}
}
