package main

import (
	"fmt"
	"net/http"
)

func APIHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey the API works!")
}

func main() {
	http.HandleFunc("/api", APIHandler)

	fmt.Println("Server is running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
