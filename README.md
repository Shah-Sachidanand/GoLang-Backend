# GoLang Backend Boilerplate

This is a backend project built with Go, providing a robust and scalable server implementation.

## Prerequisites

- Go (version 1.16 or higher)
- Git
- MongoDB database 

## Project Structure

```
GoLang-Backend/
├── cmd/
│   └── main.go
├── internal/
│   ├── handlers/
│   ├── models/
│   └── database/
├── pkg/
├── configs/
└── README.md
```

## Getting Started

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd GoLang-Backend
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up environment variables:
   Copy `.env.example` to create a new `.env` file in the root directory:
   ```bash
   cp .env.example .env
   ```
   Then modify the values in `.env` according to your setup:
   ```env
   PORT=8080
   MONGO_URI=mongodb://localhost:27017
   DB_NAME=GolangDB
   ALLOWED_ORIGINS=http://localhost:3000
   ENV=development
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

## API Documentation

### Swagger Documentation

The API is documented using Swagger/OpenAPI. Once the server is running, you can access the Swagger UI at:

```
http://localhost:8080/swagger/index.html
```

To generate/update Swagger documentation, run:
```bash
swag init -g main.go
```

Make sure you have swaggo installed:
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```