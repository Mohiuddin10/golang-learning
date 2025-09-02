package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from server")
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from server at about")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)

	mux.HandleFunc("/about", aboutHandler)

	fmt.Println("server is running at port: 3000")

	err := http.ListenAndServe(":3000", mux)

	if err != nil {
		fmt.Println("Error:", err)
	}
}
