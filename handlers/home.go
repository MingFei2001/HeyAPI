package handlers

import "net/http"

// function to serve the html page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}
