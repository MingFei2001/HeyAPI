package handlers

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"
)

// factory function that returns a handler function to serves version information retrieved from main.go
func VersionHandler(version string, startTime time.Time) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
