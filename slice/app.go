package main

import "fmt"

func main() {
	arr := [6]string{"This", "is", "a", "slice", "example", "!"}
	s := arr[1:4]
	fmt.Println(s)

	s1 := s[1:2]
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	s2 := make([]int, 7)
	s2[3] = 7
	s3 := s2[2:5]
	fmt.Println(s3)
	fmt.Println(len(s3))
	fmt.Println(cap(s3))

	s4 := make([]int, 3, 5) // len 3 cap 5
	s4[1] = 3
	fmt.Println(s4)
	fmt.Println(len(s4))
	fmt.Println(cap(s4))
}

// 1. slice from existing array
// 2. slice from a slice
// 3. slice literal
// 4. make function with len
// 5. make function with len and cap
