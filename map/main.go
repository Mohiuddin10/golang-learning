package main

import "fmt"

type floatMap map[string]float64

func (f floatMap) output() {
	fmt.Println(f)
	fmt.Println(len(f))
}

func main() {
	// userName := []string{}
	userName := make([]string, 2, 4)
	userName = append(userName, "Mohammad")
	userName = append(userName, "Mohiuddin")
	userName = append(userName, "Sujon")
	userName = append(userName, "Mainudddin")
	userName = append(userName, "Sumon")
	fmt.Println(userName)
	fmt.Println(len(userName))
	fmt.Println(cap(userName))

	// courseRating := map[string]float64{}
	courseRating := make(floatMap, 3)

	courseRating["Angular"] = 4.5
	courseRating["React"] = 4.7
	courseRating["Vue"] = 4.6
	courseRating["Node"] = 4.8
	courseRating.output()

	for index, value := range userName {
		fmt.Println(index, value)
	}

	for key, value := range courseRating {
		fmt.Println("key: ", key)
		fmt.Println("value:", value)
	}
}
