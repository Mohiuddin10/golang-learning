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

	var product1 = Product{
		title:     "A4 Tech mouse",
		id:        "P001",
		price:     29.99,
		createdAt: time.Now(),
	}

	fmt.Println(product1)
}
