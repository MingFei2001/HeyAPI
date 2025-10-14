# HeyAPI

A simple web server built with Golang for serving static pages and RESTful APIs. This project is for learning the basics of Go web development.

## Features

*   Serves static HTML files.
*   Handles HTTP requests for API endpoints.
*   Basic routing and request handling.

## Getting Started

### Prerequisites

*You need to have Go installed on your system.*

- On Linux, install Go using your distro's package manager:
- On Windows, install Go using the official installer from [golang.org](https://golang.org/dl/).

### Cloning the Repository

```bash
git clone https://github.com/MingFei2001/HeyAPI.git
cd HeyAPI
```

### Running the Server

To run the web server, execute the following command in the project root directory:

```bash
go run main.go
```

The server will typically start on `http://localhost:8080` (or whatever port is configured in `main.go`).

### Accessing Web Pages

Once the server is running, you can open your web browser and navigate to:

*   `http://localhost:8080/` - To access the home page, which serves `templates/index.html`.

*(Note: Additional page paths depend on the server's routing configuration.)*

**Important:** During development, if CSS or other static changes don't appear, try a hard refresh (Ctrl+F5 or Cmd+Shift+R) to clear your browser's cache.

### API Endpoints

The server also exposes various API endpoints. You can test these using command-line tools like `curl`, API clients like Postman, or FOSS alternatives such as Requestly.

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

*(Note: Specific API endpoints will be defined within the `main.go` or related handler files.)*

## Project Structure

*   `main.go`: The entry point of the application, defining routes and handlers.
*   `templates/`: Directory for static HTML templates (e.g., `index.html`).
*   `api/`: (Optional) Directory for API-specific handlers and logic.

## TODO

*   Split `main.go` file into handlers for better organization.
*   Implement additional API endpoints.
*   Implement a testing mechanism.
*   Implement error handling and logging.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
