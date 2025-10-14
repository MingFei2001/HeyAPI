# HeyAPI

A simple web server built with Golang, designed to serve both static web pages and RESTful API endpoints. This project is a learning endeavor to understand the fundamentals of building web applications in Go.

## Features

*   Serves static HTML files.
*   Handles HTTP requests for API endpoints.
*   Basic routing and request handling.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

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

### API Endpoints

The server also exposes various API endpoints. You can test these using tools like `curl`, Postman, or your browser's developer tools.

Example API endpoint:

*   **GET `/api`**: Returns a simple "Hey the API works!" message.
    ```bash
    curl http://localhost:8080/api
    ```

*(Note: Specific API endpoints will be defined within the `main.go` or related handler files.)*

## Project Structure (Expected)

*   `main.go`: The entry point of the application, defining routes and handlers.
*   `api/`: (Optional) Directory for API-specific handlers and logic.

## Contributing

This is a personal learning project, but feel free to explore.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
