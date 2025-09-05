package main

import "fmt"

func main() {
	userName := []string{}
	userName = append(userName, "Mohammad")
	userName = append(userName, "Mohiuddin")
	userName = append(userName, "Sujon")
	userName = append(userName, "Mainudddin")
	userName = append(userName, "Sumon")
	fmt.Println(userName)
	fmt.Println(len(userName))
	fmt.Println(cap(userName))
}
