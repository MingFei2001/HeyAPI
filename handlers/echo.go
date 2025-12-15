package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	// some http stuff
	w.Header().Set("Content-Type", "application/json")

	// Handle non-POST requests
	if r.Method != http.MethodPost {
		log.Printf("Invalid method: %s", r.Method)
		http.Error(w, "Only POST is allowed on this API endpoint.", http.StatusMethodNotAllowed)
		return
	}

	// configure JSON payload type as ANY cuz i am lazy
	var payload map[string]any
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Printf("Invalid JSON payload: %v", err)
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// if nothing goes wrong, echo back the payload
	json.NewEncoder(w).Encode(payload)
}
