package main

import "fmt"

func main() {
	var website = map[int]string{
		1: "https://www.example.com",
		2: "https://www.example.org",
		3: "https://www.example.net",
	}
	fmt.Println(website)
	fmt.Println(website[2])
	// assign value
	website[4] = "https://www.example.edu"
	fmt.Println(website)

	// delete key/value
	delete(website, 2)
	fmt.Println(website)
}
