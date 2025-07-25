package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	user1 := User{Name: "Sujon", Age: 34}
	user2 := User{Name: "Sumon", Age: 32}
	user3 := User{Name: "Honey", Age: 28}

	fmt.Println("User 1 Name:", user1.Name)
	fmt.Println("User 1 Age:", user1.Age)
	fmt.Println("User 2 Name:", user2.Name)
	fmt.Println("User 2 Age:", user2.Age)
	fmt.Println("User 3 Name:", user3.Name)
	fmt.Println("User 3 Age:", user3.Age)
}
