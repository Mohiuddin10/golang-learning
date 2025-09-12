package main

import "fmt"

func makeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func main() {

	add10 := makeAdder(20)
	fmt.Println(add10(10))

	add20 := makeAdder(30)
	fmt.Println(add20(20))

}
