package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
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

// function to serve a random number in JSON
func RandomHandler(w http.ResponseWriter, r *http.Request) {
	// wrap random number with JSON as response
	w.Header().Set("Content-Type", "application/json")
	resp := map[string]int{"random": rand.Intn(100)}
	json.NewEncoder(w).Encode(resp)
}

// main function to start the server
func main() {
	// Serve static files from /static/
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/api", APIHandler)
	http.HandleFunc("/api/random", RandomHandler)

	fmt.Println("Server is running at http://localhost" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
