package main

import (
	"fmt"
	"net/http"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Golang Server!")
}

func main() {

	// Don't block the main thread, start HTTP server in a new goroutine
	go func() {
		http.HandleFunc("/", getHandler)

		fmt.Println("Starting server on port 8080...")
	
		err:= http.ListenAndServe(":8080", nil)
		if (err != nil) {
			fmt.Println("Error starting server on port 8080...")
		}
	}()

	println("Server is running in the background..and you can perform tasks asynchronously")

	// Keep running the main thread (blocks it forever)
	select{}

}
