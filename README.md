# GoLang Backend Boilerplate

This is a backend project built with Go, providing a robust and scalable server implementation.

## Prerequisites

- Go (version 1.16 or higher)
- Git
- MongoDB database
- Docker (optional, for containerized deployment)

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
├── Dockerfile
├── docker-compose.yml
└── README.md
```

## Getting Started

### Local Development

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
   go run cmd/main.go
   ```

### Using Docker

1. **Build the Docker Image**:
   ```bash
   docker build -t golang-backend .
   ```

2. **Run the Docker Container**:
   ```bash
   docker run -d --name golang-backend -p 8080:8080 --env-file .env golang-backend
   ```

3. **Using Docker Compose**:
   If you have a `docker-compose.yml` file, you can start both the backend and MongoDB services:
   ```bash
   docker-compose up -d
   ```

   Example `docker-compose.yml`:
   ```yaml
   version: '3.8'
   services:
     app:
       build:
         context: .
       ports:
         - "8080:8080"
       env_file:
         - .env
       depends_on:
         - mongo
     mongo:
       image: mongo:latest
       ports:
         - "27017:27017"
       volumes:
         - mongo-data:/data/db
   volumes:
     mongo-data:
   ```

4. **Verify the Application**:
   After running the container, visit:
   ```
   http://localhost:8080
   ```

## API Documentation

### Swagger Documentation

The API is documented using Swagger/OpenAPI. Once the server is running, you can access the Swagger UI at:

```
http://localhost:8080/swagger/index.html
```

To generate/update Swagger documentation, run:
```bash
swag init -g cmd/main.go
```

Make sure you have swaggo installed:
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

## Environment Variables

- **PORT**: The port the application runs on.
- **MONGO_URI**: The URI of the MongoDB server.
- **DB_NAME**: The database name to use.
- **ALLOWED_ORIGINS**: CORS allowed origins.
- **ENV**: The environment (e.g., development, production).