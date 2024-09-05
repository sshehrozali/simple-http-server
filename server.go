package main

import (
	"fmt"
	"net/http"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Golang Server!")
}

func main() {

	http.HandleFunc("/", getHandler)

	fmt.Println("Starting server on port 8080...")

	err:= http.ListenAndServe(":8080", nil)
	if (err != nil) {
		fmt.Println("Error starting server on port 8080...")
	}

	fmt.Println("Hello Golang!")
}
