package handlers

import "net/http"

// main logic
func CurrencyHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/currency.html")
}
