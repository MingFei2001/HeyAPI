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

// defines structure for api.fxratesapi.com endpoint response
type CurrencyAPIResponse struct {
	Success   bool               `json:"success"` // Indicates if the API call was successful
	Timestamp int64              `json:"timestamp"`
	Base      string             `json:"base"`  // The base currency used for the rates
	Date      string             `json:"date"`  // The date for which rates are given
	Rates     map[string]float64 `json:"rates"` // A map of currency codes to their exchange rates

	// capture errors if success is false
	Error *struct {
		Code int    `json:"code"`
		Type string `json:"type"`
		Info string `json:"info"`
	} `json:"error"`
}

// holds data to render in HTML template
type CurrencyPageData struct {
	Amount         float64
	FromCurrency   string
	ToCurrency     string
	ExchangeRate   float64
	ConvertedValue float64
	Error          string // To display user-friendly error messages on the page
}

// fetches currency conversion rate and renders it to HTML page using templates
func CurrencyHandler(w http.ResponseWriter, r *http.Request) {
	// Get API Key from environment variable
	apiKey := os.Getenv("EXCHANGERATE_API_KEY")
	if apiKey == "" {
		log.Println("Error: EXCHANGERATE_API_KEY environment variable not set.")
		http.Error(w, "Currency Exchange API key is not configured.", http.StatusInternalServerError)
		return
	}

	// Parse the HTML template file
	tmpl, err := template.ParseFiles("templates/currency.html")
	if err != nil {
		log.Printf("Template parsing error: %v", err)
		http.Error(w, "There was an error rendering the page.", http.StatusInternalServerError)
		return
	}

	// Extract user inputs from query parameters
	amountStr := r.URL.Query().Get("amount")
	fromCurrency := r.URL.Query().Get("from")
	toCurrency := r.URL.Query().Get("to")

	// Handle initial page load or incomplete input as empty form
	if amountStr == "" || fromCurrency == "" || toCurrency == "" {
		tmpl.Execute(w, nil)
		return
	}

	// check and clean amountStr by converting it to float64
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

	// Construct API URL for "latest" rates endpoint
	currencyAPIURL := fmt.Sprintf(
		"https://api.fxratesapi.com/latest?api_key=%s&base=%s&currencies=%s&format=json",
		apiKey, fromCurrency, toCurrency,
	)

	// Call the currency API using GET method
	resp, err := http.Get(currencyAPIURL)
	if err != nil {
		log.Printf("Error fetching currency data from api.fxratesapi.com: %v", err)
		tmpl.Execute(w, CurrencyPageData{Error: "Could not connect to the currency exchange service."})
		return
	}
	defer resp.Body.Close()

	// Decode JSON response into variable
	var currencyData CurrencyAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&currencyData); err != nil {
		log.Printf("Error decoding api.fxratesapi.com JSON: %v", err)
		tmpl.Execute(w, CurrencyPageData{Error: "Could not parse currency data from the service."})
		return
	}

	// Check for API-specific errors by "success: false" and "error" object in API errors
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

	// Ensure the base currency in the response matches what we requested
	if currencyData.Base != fromCurrency {
		log.Printf("API returned base currency '%s', but requested '%s'.", currencyData.Base, fromCurrency)
		tmpl.Execute(w, CurrencyPageData{Error: fmt.Sprintf("API returned unexpected base currency: %s", currencyData.Base)})
		return
	}

	// Extract the exchange rate and perform calculation
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

	// Prepare data for HTML template
	pageData := CurrencyPageData{
		Amount:         amount,
		FromCurrency:   fromCurrency,
		ToCurrency:     toCurrency,
		ExchangeRate:   exchangeRate,
		ConvertedValue: convertedValue,
		Error:          "", // No error
	}

	// Render the template if all goes well
	tmpl.Execute(w, pageData)
}
