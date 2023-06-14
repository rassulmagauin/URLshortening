# URL Shortening Service

This is a URL shortening service written in Go. It's a simple and effective tool to make long URLs more manageable. This can be particularly useful when sharing links, especially in situations where character count matters, like Twitter posts.

## Features

- **Shorten URLs:** Convert long URLs into short, manageable ones.
- **Redirect:** Use the shortened URL to redirect to the original URL.

## Getting Started

### Prerequisites

- Go (Check the version in `go.mod`)

### Installation

1. Clone the repo: `git clone https://github.com/rassulmagauin/URLshortening.git`
2. Navigate to the project directory: `cd URLshortening`
3. Install dependencies: `go mod download`
4. Run the application: `go run main.go`

## Project Structure

- `main.go`: The entry point of the application.
- `models/models.go`: Contains the data models used in the project.
- `utils/encodeutils.go`: Contains utility functions for encoding.
