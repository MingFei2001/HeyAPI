package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime"
	"time"
)

// factory function that returns a handler function to serves version information retrieved from main.go
func VersionHandler(version string, startTime time.Time) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ensure the method is GET only
		if r.Method != http.MethodGet {
			log.Printf("Invalid method: %s. Only GET is allowed.", r.Method)
			http.Error(w, "Method not allowed. Only GET is allowed.", http.StatusMethodNotAllowed)
			return
		}

		// some http stuff
		w.Header().Set("Content-Type", "application/json")
		resp := map[string]string{
			"status":     "ok",
			"version":    version,
			"go_version": runtime.Version(),
			"uptime":     time.Since(startTime).String(),
		}
		json.NewEncoder(w).Encode(resp)
	}
}
