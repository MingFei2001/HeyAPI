package handlers

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

// function to serve a random number in JSON
func RandomHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the method is GET only
	if r.Method != http.MethodGet {
		log.Printf("Invalid method: %s. Only GET is allowed.", r.Method)
		http.Error(w, "Method not allowed. Only GET is allowed.", http.StatusMethodNotAllowed)
		return
	}

	// wrap random number with JSON as response
	w.Header().Set("Content-Type", "application/json")
	resp := map[string]int{"random": rand.Intn(100)}
	json.NewEncoder(w).Encode(resp)
}
