package main

import "fmt"

type str string

func (text str) print() {
	fmt.Println(text)
}
func main() {
	var name str = "Hello world"
	name.print()

	var a int8 = -12
}
