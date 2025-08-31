package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from server")
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from server")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)

	mux.HandleFunc("/about", aboutHandler)
}
