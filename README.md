# 🌐 HeyAPI

A simple web server built with Golang for serving static pages and RESTful APIs. This project is for learning the basics of Go web development.

## ⚙️ Features

*   Serves static HTML files.
*   Handles HTTP requests for API endpoints.
*   Basic routing and request handling.

## 🚀 Getting Started

*You need to have Go installed on your system.*

- On Linux, install Go using your distro's package manager:
- On Windows, install Go using the official installer from [golang.org](https://golang.org/dl/).

```bash
git clone https://github.com/MingFei2001/HeyAPI.git
cd HeyAPI
go run main.go
```

The server will typically start on `http://localhost:8080` (or whatever port is configured in `main.go`).

### Accessing Web Pages

Once the server is running, you can open your web browser and navigate to:

*   `http://localhost:8080/` - To access the home page.
*   `http://localhost:8080/weather` - To access the weather page.

#### Weather API Key Setup

The weather page (`/weather`) fetches current weather conditions from `weatherapi.com`. To make this work, you need an API key from `weatherapi.com` and set it as environment variable to fetch the weather information.

1.  **Obtain a Key from `weatherapi.com`**:
*   Go to [weatherapi.com](https://www.weatherapi.com/), sign up or log in, and copy your API key from your dashboard.

2.  **Set the Environment Variable**:
    Before running `go run main.go`, open your terminal and set the `WEATHERAPI_KEY` environment variable:

    ```bash
    export WEATHERAPI_KEY="your_actual_weatherapi_key_here"
    ```
    (Replace `"your_actual_weatherapi_key_here"` with the key you copied. This variable is only set for the current terminal session.)

**Important:** During development, if CSS or other static changes don't appear, try a hard refresh (Ctrl+F5 or Cmd+Shift+R) to clear browser's cache.

### API Endpoints

The server also serves various API endpoints. You can test these using tools like `curl`, API clients like `Postman`, or FOSS alternatives like `Requestly`.

Example API endpoint:

*   **GET `/api`**: Returns a simple "Hey the API works!" message.
    ```bash
    curl http://localhost:8080/api
    ```
*   **GET `/api/random`**: Returns a JSON object containing a random integer between 0 and 99.
    ```bash
    curl http://localhost:8080/api/random
    ```
    Example response: `{"random": 42}`
*   **GET `/api/version`**: Returns JSON object with server version, Go runtime version, and uptime.
    ```bash
    curl http://localhost:8080/api/version
    ```
    Example response: `{"go_version": "go1.21.0","status": "ok","uptime": "2m3.456s","version": "0.0.3"}`
*   **POST `/api/echo`**: Accepts a JSON payload and echoes it back in the response.
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"message": "Hello", "value": 123}' http://localhost:8080/api/echo
    ```
    Example response (same as input payload): `{"message": "Hello", "value": 123}`
    *(Note: Only POST requests are allowed for this endpoint.)*

*(Note: API endpoints are defined in `main.go` and handled by functions in `handlers/` directory.)*

## 📂 Project Structure

*   `main.go`: The entry point of the application, defining routes and handlers.
*   `templates/`: Directory for static HTML templates (e.g., `index.html`).
*   `handlers/`: Directory for API-specific handlers and logic.
*   `.gitignore`: Specifies intentionally untracked files that Git should ignore (e.g., `.env` for API keys).

## 📝 TODO

- [x] Split `main.go` file into handlers for better organization.
- [x] Implement a weather page using external API.
- [ ] Add a currency rate page using external API.
- [ ] Add a database (sqlite) to store information.
- [ ] Implement additional API endpoints.
- [ ] Implement a testing mechanism.
- [ ] Dockerize the project.

## 📢 Credits
Shoutout to **The man, The myth, The legend: [ThePrimeagen](https://github.com/theprimeagen)** for introducing me to Go and helping me learn many things, *no joke* i learned half of my networking from him.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
