package main

import "fmt"

var a = 10
var b = 20

func add(x int, y int) {
	z := x + y
	fmt.Println("Inside add function, z =", z)
}

func main() {
	var c = 30
	var d = 40
	add(a, b)
	add(c, d)
	add(a, d)
}
