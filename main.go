package main

import (
	"fmt"
	"net/http"
)

// define port number used
const port = ":8080"

// function to serve the html page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

// function to serve the API endpoint
func APIHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey the API works!")
}

// main function to start the server
func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/api", APIHandler)

	fmt.Println("Server is running at http://localhost" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
