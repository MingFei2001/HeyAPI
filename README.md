# 🌐 HeyAPI

A simple web server built with Golang for serving static pages and RESTful APIs. This project is for learning the basics of Go web development.

## ⚙️ Features

*   Serves static HTML files.
*   Handles HTTP requests for API endpoints.
*   Basic routing and request handling.
*   Dynamic note-taking feature with form submission and table display.

## 🚀 Getting Started

*You need to have Go installed on your system.*

- On Linux, install Go using your distro's package manager:
- On Windows, install Go using the official installer from [golang.org](https://golang.org/dl/).

```bash
git clone https://github.com/MingFei2001/HeyAPI.git
cd HeyAPI
go run main.go
```

The server will by default start on `http://localhost:8080` (you can configure it in `main.go`).

### Accessing Web Pages

Once the server is running, you can open your web browser and navigate to:

*   `http://localhost:8080/` - To access the home page.
*   `http://localhost:8080/weather` - To access the weather page.
*   `http://localhost:8080/currency` - To access the currency converter page.
*   `http://localhost:8080/notes` - To access the notes page.

#### API Key Setup

The server integrates with external APIs for certain features.

**Weather API Key Setup (`weatherapi.com`)**:

The weather page (`/weather`) fetches current weather conditions from `weatherapi.com`. To make this work, you need an API key from `weatherapi.com` and must set it as an environment variable before running the server.

1.  **Obtain a Key from `weatherapi.com`**:
    *   Go to [https://www.weatherapi.com/](https://www.weatherapi.com/).
    *   Sign up for a free account or log in.
    *   Find and copy your API key from your dashboard.

2.  **Set the Environment Variable**:\
    Before running `go run main.go`, open your terminal and set the `WEATHERAPI_KEY` environment variable:

    ```bash
    export WEATHERAPI_KEY="your_weatherapi_key_here"
    ```
    (Replace `"your_weatherapi_key_here"` with your actual key. This variable is only set for the current terminal session.)

**Currency Converter API Key Setup (`api.fxratesapi.com`)**:

The currency converter page (`/currency`) fetches exchange rates from `api.fxratesapi.com`. You need an API key from this service and set it as an environment variable.

1.  **Obtain a Key from `api.fxratesapi.com`**:
    *   Go to [https://fxratesapi.com/](https://fxratesapi.com/).
    *   Sign up for a free account or log in.
    *   Find and copy your API key (or 'Access Token') from your dashboard.

2.  **Set the Environment Variable**:\
    Before running `go run main.go`, open your terminal and set the `EXCHANGERATE_API_KEY` environment variable:

    ```bash
    export EXCHANGERATE_API_KEY="your_fxratesapi_key_here"
    ```
    (Replace `"your_fxratesapi_key_here"` with your actual key. This variable is only set for the current terminal session.)


**Important:** During development, if CSS or other static changes don't appear, try a hard refresh (Ctrl+F5 or Cmd+Shift+R) to clear browser's cache.

### API Endpoints

The server also serves various API endpoints. You can test these using tools like `curl`, API clients like `Postman`, or FOSS alternatives like `Requestly`.

#### Basic Endpoints

*   **GET `/api`**: Returns a simple "Hey the API works!" message.
    ```bash
    curl http://localhost:8080/api
    ```

*   **GET `/random`**: Returns a JSON object containing a random integer between 0 and 99.
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

#### Notes Endpoints

*   **POST `/notes/create`**: Creates a new note. The note is submitted via a form or `curl`:
    ```bash
    curl -X POST -d "note=This is a test note" http://localhost:8080/notes/create
    ```
    After submission, the user is redirected to the `/notes` page.

*   **GET `/notes`**: Displays the notes page with a form for creating notes and a table for viewing existing notes.

*   **GET `/notes/get`**: Retrieves all notes in JSON format.

*   **GET `/notes/getbyid?id=<id>`**: Retrieves a specific note by its ID.

*   **DELETE `/notes/delete?id=<id>`**: Deletes a specific note by its ID.

*(Note: API endpoints are defined in `main.go` and handled by functions in `handlers/` directory.)*

## 📂 Project Structure

*   `main.go`: The entry point of the application, defining routes and handlers.
*   `templates/`: Directory for static HTML templates (e.g., `index.html`).
*   `handlers/`: Directory for API-specific handlers and logic (e.g., `random.go`, `note.go`).
*   `.gitignore`: Specifies intentionally untracked files that Git should ignore (e.g., `.env` for API keys).

## 🐳 Dockerization

To containerize and run this project using Docker, follow the [Dockerization Guide](dockerize.md).

## 📝 TODO

- [x] Split `main.go` file into handlers for better organization.
- [x] Implement a weather page using external API.
- [x] Implement a currency exchange rate page using external API.
- [x] Fix the form in `templates/currency.html`.
- [x] Dockerize the project.
- [ ] Use LXC to containerize the project.

## 📢 Credits
Shoutout to **The man, The myth, The legend: [ThePrimeagen](https://github.com/theprimeagen)** for introducing me to Go and helping me learn many things, *no joke* i learned half of my networking from him.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
