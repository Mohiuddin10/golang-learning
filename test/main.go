package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	go add(1, 3)
	time.Sleep(2 * time.Second)
	fmt.Println("after 2 sec")
	time.Sleep(6 * time.Second)
	fmt.Println("after 6 sec")

}

func add(x int, y int) {
	time.Sleep(1 * time.Second)
	fmt.Println("The result is:", x+y)
}
