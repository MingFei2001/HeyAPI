package handlers

import (
	"encoding/json"
	"net/http"
)

func EchoHandler(w http.ResponseWriter, r *http.Request) {
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
