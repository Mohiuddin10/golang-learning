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
func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from server cntact info")
}
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from server of home")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)

	mux.HandleFunc("/about", aboutHandler)

	mux.HandleFunc("/contact", contact)

	mux.HandleFunc("/home", home)

	fmt.Println("server is running at port: 3000")

	err := http.ListenAndServe(":3000", mux)

	if err != nil {
		fmt.Println("Error:", err)
	}
}
