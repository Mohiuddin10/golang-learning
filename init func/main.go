package main

import "fmt"

var a = 10

func main() {
	fmt.Println("a = ", a)
}

func init() {
	fmt.Println("first print from init func", a)
	a = 50

}
