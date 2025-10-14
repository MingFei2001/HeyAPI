package main

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func APIHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey the API works!")
}

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/api", APIHandler)

	fmt.Println("Server is running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
