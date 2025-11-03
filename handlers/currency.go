package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"      // Needed for os.Getenv
	"strconv" // Needed for converting string to float64
	"strings" // Needed for strings.ToUpper
)

// CurrencyAPIResponse defines the structure for api.fxratesapi.com's latest rates endpoint response
type CurrencyAPIResponse struct {
	Success   bool               `json:"success"` // Indicates if the API call was successful
	Timestamp int64              `json:"timestamp"`
	Base      string             `json:"base"`  // The base currency used for the rates
	Date      string             `json:"date"`  // The date for which rates are given
	Rates     map[string]float64 `json:"rates"` // A map of currency codes to their exchange rates

	// This field will capture API-specific errors if success is false
	Error *struct {
		Code int    `json:"code"`
		Type string `json:"type"`
		Info string `json:"info"`
	} `json:"error"`
}

// CurrencyPageData holds the specific data to render in the HTML template
type CurrencyPageData struct {
	Amount         float64
	FromCurrency   string
	ToCurrency     string
	ExchangeRate   float64
	ConvertedValue float64
	Error          string // To display user-friendly error messages on the page
}

// CurrencyHandler fetches currency conversion data using api.fxratesapi.com and renders it to an HTML page.
func CurrencyHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Get API Key from environment variable
	apiKey := os.Getenv("EXCHANGERATE_API_KEY") // Using the specified environment variable name
	if apiKey == "" {
		log.Println("Error: EXCHANGERATE_API_KEY environment variable not set.")
		http.Error(w, "Currency Exchange API key is not configured.", http.StatusInternalServerError)
		return
	}

	// 2. Parse the HTML template file
	tmpl, err := template.ParseFiles("templates/currency.html")
	if err != nil {
		log.Printf("Template parsing error: %v", err)
		http.Error(w, "There was an error rendering the page.", http.StatusInternalServerError)
		return
	}

	// 3. Extract user inputs from query parameters
	amountStr := r.URL.Query().Get("amount")
	fromCurrency := r.URL.Query().Get("from")
	toCurrency := r.URL.Query().Get("to")

	// 4. Handle initial page load or incomplete input
	// If any of the required parameters are missing, render the empty form.
	if amountStr == "" || fromCurrency == "" || toCurrency == "" {
		tmpl.Execute(w, nil)
		return
	}

	// Convert amount string to float64
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		log.Printf("Error parsing amount '%s': %v", amountStr, err)
		tmpl.Execute(w, CurrencyPageData{Error: "Invalid amount provided."})
		return
	}
	if amount <= 0 {
		tmpl.Execute(w, CurrencyPageData{Error: "Amount must be a positive number."})
		return
	}

	// Convert currency codes to uppercase for consistent API requests
	fromCurrency = strings.ToUpper(fromCurrency)
	toCurrency = strings.ToUpper(toCurrency)

	// 5. Construct API URL for the "latest" rates endpoint
	// Request only the target currency to keep the response small and efficient.
	// Endpoint: https://api.fxratesapi.com/latest?api_key=YOUR_ACCESS_TOKEN&base=USD&currencies=EUR&format=json
	currencyAPIURL := fmt.Sprintf(
		"https://api.fxratesapi.com/latest?api_key=%s&base=%s&currencies=%s&format=json",
		apiKey, fromCurrency, toCurrency,
	)

	// 6. Make HTTP GET request to the currency API
	resp, err := http.Get(currencyAPIURL)
	if err != nil {
		log.Printf("Error fetching currency data from api.fxratesapi.com: %v", err)
		tmpl.Execute(w, CurrencyPageData{Error: "Could not connect to the currency exchange service."})
		return
	}
	defer resp.Body.Close()

	// 7. Decode JSON response
	var currencyData CurrencyAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&currencyData); err != nil {
		log.Printf("Error decoding api.fxratesapi.com JSON: %v", err)
		tmpl.Execute(w, CurrencyPageData{Error: "Could not parse currency data from the service."})
		return
	}

	// 8. Check for API-specific errors
	// api.fxratesapi.com uses a "success: false" field and an "error" object for API errors.
	if !currencyData.Success {
		if currencyData.Error != nil {
			log.Printf("FXRatesAPI returned error for %s to %s: Code %d - %s", fromCurrency, toCurrency, currencyData.Error.Code, currencyData.Error.Info)
			tmpl.Execute(w, CurrencyPageData{Error: fmt.Sprintf("Currency API Error: %s", currencyData.Error.Info)})
		} else {
			log.Println("FXRatesAPI returned unsuccessful response without specific error details.")
			tmpl.Execute(w, CurrencyPageData{Error: "Currency API returned an unknown error."})
		}
		return
	}

	// 9. Extract the exchange rate and perform calculation
	// Ensure the base currency in the response matches what we requested
	if currencyData.Base != fromCurrency {
		log.Printf("API returned base currency '%s', but requested '%s'.", currencyData.Base, fromCurrency)
		tmpl.Execute(w, CurrencyPageData{Error: fmt.Sprintf("API returned unexpected base currency: %s", currencyData.Base)})
		return
	}

	exchangeRate, ok := currencyData.Rates[toCurrency]
	if !ok {
		log.Printf("Target currency '%s' not found in API response rates for base '%s'.", toCurrency, fromCurrency)
		tmpl.Execute(w, CurrencyPageData{Error: fmt.Sprintf("Exchange rate for '%s' not available for base '%s'.", toCurrency, fromCurrency)})
		return
	}
	if exchangeRate <= 0 { // Sanity check for rate
		tmpl.Execute(w, CurrencyPageData{Error: "Invalid exchange rate returned by API."})
		return
	}

	convertedValue := amount * exchangeRate

	// 10. Prepare data for HTML template (CurrencyPageData)
	pageData := CurrencyPageData{
		Amount:         amount,
		FromCurrency:   fromCurrency,
		ToCurrency:     toCurrency,
		ExchangeRate:   exchangeRate,
		ConvertedValue: convertedValue,
		Error:          "", // No error
	}

	// 11. Execute the template
	tmpl.Execute(w, pageData)
}
