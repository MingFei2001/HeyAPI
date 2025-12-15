package handlers

import (
	"log"
	"net/http"
	"os"
)

// function to serve the html page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	filePath := "templates/index.html"

	_, err := os.Stat(filePath)
	// If the file does not exist, log the error and return 404
	if os.IsNotExist(err) {
		log.Printf("File not found: %s", filePath)
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	} else if err != nil {
		// Handles all other errors
		log.Printf("Error checking file: %s", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Only execute if the file exists and no other errors occurred
	http.ServeFile(w, r, "templates/index.html")
	log.Printf("Served file: %s, %d", filePath, http.StatusOK)
}
