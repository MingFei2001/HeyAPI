package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func APIHandler(w http.ResponseWriter, r *http.Request) {

	// Handles non GET requests, return 405
	if r.Method != http.MethodGet {
		log.Printf("Invalid method: %s", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// If nothing goes wrong
	fmt.Fprintf(w, "Hey the API works!")
}
