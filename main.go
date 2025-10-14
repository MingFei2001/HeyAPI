package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"time"
)

// define constant
const port = ":8080"
const version = "0.0.3"

// define variable
var startTime = time.Now()

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

// function to serve version information as JSON
func VersionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := map[string]string{
		"status":     "ok",
		"version":    version,
		"go_version": runtime.Version(),
		"uptime":     time.Since(startTime).String(),
	}
	json.NewEncoder(w).Encode(resp)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Handle non-POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST is allowed on this API endpoint.", http.StatusMethodNotAllowed)
		return
	}

	// configure JSON payload type as ANY
	var payload map[string]any
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(payload)
}

// main function to start the server
func main() {
	// Serve static files from /static/
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Define API endpoints to serve
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/api", APIHandler)
	http.HandleFunc("/api/random", RandomHandler)
	http.HandleFunc("/api/version", VersionHandler)
	http.HandleFunc("/api/echo", echoHandler)

	// Stdout Message
	fmt.Println("Server is running at http://localhost" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
