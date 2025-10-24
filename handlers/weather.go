package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Hold decoded JSON response from weatherapi.com
type WeatherAPIResponse struct {
	// Nested struct for location details
	Location struct {
		// City name
		Name string `json:"name"`
	} `json:"location"`
	// Nested struct for current weather conditions
	Current struct {
		TempC    float64 `json:"temp_c"`
		Humidity int     `json:"humidity"`
		// Nested struct for weather condition description
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	// might return an "error" field, use HTTP status and core data to check
	Error *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// holds data to pass to HTML template
type WeatherPageData struct {
	City        string
	Temperature float64
	Humidity    int
	Description string
	Error       string
}

// fetches weather data with API and renders it to HTML template
func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the API key from environment variable
	apiKey := os.Getenv("WEATHERAPI_KEY")
	if apiKey == "" {
		log.Println("Error: WEATHERAPI_KEY environment variable not set.")
		http.Error(w, "Weather API key is not configured.", http.StatusInternalServerError)
		return
	}

	// Parse the HTML template file. This prepares the template for rendering.
	tmpl, err := template.ParseFiles("templates/weather.html")
	if err != nil {
		log.Printf("Template parsing error: %v", err)
		http.Error(w, "There was an error rendering the page.", http.StatusInternalServerError)
		return
	}

	// Extract the 'city' query parameter from the URL (e.g., /weather?city=London).
	city := r.URL.Query().Get("city")

	// If no city is provided, render the template with no data,
	if city == "" {
		tmpl.Execute(w, nil)
		return
	}

	// --- Step 1: Fetch Current Weather from weatherapi.com ---
	// Construct the URL for weatherapi.com current weather for the specified city in Celsius.
	weatherAPIURL := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, city)

	// Make an HTTP GET request to the weather API.
	resp, err := http.Get(weatherAPIURL)
	if err != nil {
		log.Printf("Error fetching weather data from WeatherAPI.com: %v", err)
		tmpl.Execute(w, WeatherPageData{Error: "Could not connect to the weather service."})
		return
	}
	defer resp.Body.Close() // Ensure the response body is closed to prevent resource leaks.

	// Decode the JSON response into our `WeatherAPIResponse` struct.
	var weatherData WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		log.Printf("Error decoding WeatherAPI.com JSON for city '%s': %v", city, err)
		tmpl.Execute(w, WeatherPageData{Error: "Could not parse weather data from the service."})
		return
	}

	// Check for API-specific errors embedded in the JSON response
	if weatherData.Error != nil {
		log.Printf("WeatherAPI.com returned an error for '%s': Code %d - %s", city, weatherData.Error.Code, weatherData.Error.Message)
		tmpl.Execute(w, WeatherPageData{Error: fmt.Sprintf("Weather API Error: %s", weatherData.Error.Message)})
		return
	}

	// Basic validation: If the location name is empty, it usually means the city was not found.
	if weatherData.Location.Name == "" {
		log.Printf("WeatherAPI.com: No valid location data or city not found for '%s'.", city)
		tmpl.Execute(w, WeatherPageData{Error: fmt.Sprintf("Could not find weather for '%s'. Please try again with a valid city name.", city)})
		return
	}

	// Prepare the data to be passed to the HTML template.
	pageData := WeatherPageData{
		City:        weatherData.Location.Name,
		Temperature: weatherData.Current.TempC,
		Humidity:    weatherData.Current.Humidity,
		Description: weatherData.Current.Condition.Text,
	}

	// Execute the template, injecting our `pageData` into the HTML placeholders.
	tmpl.Execute(w, pageData)
}
