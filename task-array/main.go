package main

import (
	"fmt"
	"time"
)

type Product struct {
	title     string
	id        string
	price     float64
	createdAt time.Time
}

func main() {
	var hobbies = [3]string{"Reading", "Traveling", "Gaming"}
	fmt.Println("Hobbies:", hobbies)

	fmt.Println("My First Hobby:", hobbies[0])
	fmt.Println("My others Hobbies are:", hobbies[1:])

	sliceHobby := hobbies[:2]
	sliceHobby1 := hobbies[0:2]
	fmt.Println("Hobbies slice of first and second:", sliceHobby)
	fmt.Println("Hobbies slice of first and second:", sliceHobby1)

	resliceHobby := sliceHobby[1:3]
	fmt.Println("Resliced Hobby:", resliceHobby)

	courseGoal := []string{"Learn Go", "Master Go", "Contribute to Go"}
	courseGoal[1] = "Become proficient in Go"
	courseGoal = append(courseGoal, "Build Go projects")
	fmt.Println("Course Goals:", courseGoal)

	products := []Product{
		{
			"mouse",
			"11",
			23.99,
			time.Now(),
		},
		{
			"keyboard",
			"12",
			29.99,
			time.Now(),
		},
	}

	products = append(products, Product{
		"headphone",
		"17",
		29.54,
		time.Now(),
	})

	fmt.Println(products)
}
