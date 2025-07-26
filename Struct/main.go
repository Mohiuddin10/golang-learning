package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// normal func
// func printUserDetails(usr User) {
// 	fmt.Println("Name:", usr.Name)
// 	fmt.Println("Age:", usr.Age)
// }

// receiver func

func (usr User) printUser() {
	fmt.Println("Name:", usr.Name)
	fmt.Println("Age:", usr.Age)
}

func (usr User) call(a int) {
	fmt.Println("Name:", usr.Name)
	fmt.Println("Age:", usr.Age)
	fmt.Println("Call:", a)
}

func main() {
	user1 := User{Name: "Sujon", Age: 34}
	user2 := User{Name: "Sumon", Age: 32}
	user3 := User{Name: "Honey", Age: 28}

	// call of normal func
	// printUserDetails(user1)
	// printUserDetails(user2)
	// printUserDetails(user3)

	// call reciver func
	user1.printUser()
	user2.printUser()
	user3.printUser()

	// call func with parameter
	user1.call(420)
	user2.call(210)
	user3.call(840)
}
