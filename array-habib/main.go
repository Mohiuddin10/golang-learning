package main

import "fmt"

var arr2 = []string{"I", "Love", "You"}

var arr3 = []string{"I", "Love", "myself"}

func main() {
	// var arr [2]int

	// arr[1] = 6

	arr := [2]int{6, 3}
	fmt.Println(arr)
	fmt.Println(arr2[2])
	arr2 = append(arr2, arr3...)
	fmt.Println("marge array:", arr2)
}
