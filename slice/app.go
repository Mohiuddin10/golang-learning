package main

import "fmt"

func changeSlice(a []int) []int {
	a[0] = 10         // 10,6,7
	a = append(a, 11) // 10, 6, 7,, 11
	return a
}

func main() {
	// arr := [6]string{"This", "is", "a", "slice", "example", "!"}
	// s := arr[1:4]
	// fmt.Println(s)

	// s1 := s[1:2]
	// fmt.Println(s1)
	// fmt.Println(len(s1))
	// fmt.Println(cap(s1))

	// s2 := make([]int, 7)
	// s2[3] = 7
	// s3 := s2[2:5]
	// fmt.Println(s3)
	// fmt.Println(len(s3))
	// fmt.Println(cap(s3))

	// s4 := make([]int, 3, 5) // len 3 cap 5
	// s4[1] = 3
	// fmt.Println(s4)
	// fmt.Println(len(s4))
	// fmt.Println(cap(s4))

	// var s5 []int
	// s5 = append(s5, 1, 2, 3, 4, 5)
	// fmt.Println(s5)
	// fmt.Println(len(s5))
	// fmt.Println(cap(s5))

	// var x []int
	// x = append(x, 1)
	// x = append(x, 2)
	// x = append(x, 3)
	// x = append(x, 4)
	// fmt.Println("x:", x) //1 2 3
	// fmt.Println(len(x))
	// fmt.Println(cap(x))

	x1 := []int{1, 2, 3, 4, 5} //1,2,3,4,5
	x1 = append(x1, 6)         //123456
	x1 = append(x1, 7)         // 1234567

	ab := x1[4:] // 567

	y1 := changeSlice(ab)

	y1 = append(y1, 13)
	y1 = append(y1, 14)
	y1 = append(y1, 15)
	y1 = append(y1, 16)

	fmt.Println("x1:", x1)          // 1234 10 6 7 11
	fmt.Println("y1:", y1)          //10 6 7 11
	fmt.Println("x1 len:", len(x1)) // 1234 10 6 7 11
	fmt.Println("y1 len:", len(y1)) //10 6 7 11
	fmt.Println("x1 cap:", cap(x1)) // 1234 10 6 7 11
	fmt.Println("y1 cap:", cap(y1)) //10 6 7 11

	// y := x

	// x = append(x, 4)
	// y = append(y, 5)

	// x[0] = 10
	// fmt.Println("x:", x) //10 2 3 4
	// fmt.Println("y:", y) //1 2 3 5
	// fmt.Println(len(y))
	// fmt.Println(cap(y))
}

// 1. slice from existing array
// 2. slice from a slice
// 3. slice literal
// 4. make function with len
// 5. make function with len and cap
// 6. empty slice or nil slice
// 7. append function
// 8. slice underlying array rule => untill 1024 its 100% ==> After 1024 its 25%
