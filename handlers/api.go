package handlers

import (
	"fmt"
	"net/http"
)

func APIHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey the API works!")
}
