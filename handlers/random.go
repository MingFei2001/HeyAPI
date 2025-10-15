package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

// function to serve a random number in JSON
func RandomHandler(w http.ResponseWriter, r *http.Request) {
	// wrap random number with JSON as response
	w.Header().Set("Content-Type", "application/json")
	resp := map[string]int{"random": rand.Intn(100)}
	json.NewEncoder(w).Encode(resp)
}
