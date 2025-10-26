package handlers

import "net/http"

// define type for decoding JSON response from exchange rate API
type CurrencyAPIResponse struct {
	// placeholder
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type CurrencyPageData struct {
	Result float64
}

// main logic
func CurrencyHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/currency.html")
}
