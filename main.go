package main

import (
	"fmt"
	"net/http"
	"time"

	"heyapi/handlers"
)

// define constant
const port = ":8080"
const version = "0.0.5"

// define variable
var startTime = time.Now()

// main function to start the server and route the requests
func main() {
	// Serve static files from /static/
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Define API endpoints to route
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/api", handlers.APIHandler)
	http.HandleFunc("/api/random", handlers.RandomHandler)
	http.HandleFunc("/api/version", handlers.VersionHandler(version, startTime))
	http.HandleFunc("/api/echo", handlers.EchoHandler)

	// Stdout Message
	fmt.Println("Server is running at http://localhost" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
