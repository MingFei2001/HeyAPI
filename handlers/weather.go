package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// This go function fetches API from weatherapi.com
// and display it on the html page through template

// define type for decoded JSON response from weatherapi.com
type WeatherAPIResponse struct {
	// Nested struct for location details
	Location struct {
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
	// might return an "error" field even for 200 status code
	Error *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// define type for holding final data to render in HTML template
type WeatherPageData struct {
	City        string
	Temperature float64
	Humidity    int
	Description string
	Error       string
}

// Main logic
func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the API key from environment variable
	apiKey := os.Getenv("WEATHERAPI_KEY")
	if apiKey == "" {
		log.Println("Error: WEATHERAPI_KEY environment variable not set.")
		http.Error(w, "Weather API key is not configured.", http.StatusInternalServerError)
		return
	}

	// Parse the HTML template file and prepares template for rendering
	tmpl, err := template.ParseFiles("templates/weather.html")
	if err != nil {
		log.Printf("Template parsing error: %v", err)
		http.Error(w, "There was an error rendering the page.", http.StatusInternalServerError)
		return
	}

	// Extract the 'city' query parameter from the URL
	city := r.URL.Query().Get("city")

	// If no city is provided, render the template with no data,
	if city == "" {
		tmpl.Execute(w, nil)
		return
	}

	// Construct URL to query weatherapi.com current weather
	weatherAPIURL := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, city)

	// Perform HTTP GET request to the weather API.
	resp, err := http.Get(weatherAPIURL)
	if err != nil {
		log.Printf("Error fetching weather data from WeatherAPI.com: %v", err)
		tmpl.Execute(w, WeatherPageData{Error: "Could not connect to the weather service."})
		return
	}
	// Ensure the response body is closed to prevent resource leaks
	defer resp.Body.Close()

	// Decode API JSON response into `WeatherAPIResponse` struct.
	var weatherData WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		log.Printf("Error decoding WeatherAPI.com JSON for city '%s': %v", city, err)
		tmpl.Execute(w, WeatherPageData{Error: "Could not parse weather data from the service."})
		return
	}

	// If error in API response, render template with error message.
	if weatherData.Error != nil {
		log.Printf("WeatherAPI.com returned an error for '%s': Code %d - %s", city, weatherData.Error.Code, weatherData.Error.Message)
		tmpl.Execute(w, WeatherPageData{Error: fmt.Sprintf("Weather API Error: %s", weatherData.Error.Message)})
		return
	}

	// If location name in response is empty, it means the city was not found thus error
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
