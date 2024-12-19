FROM golang:1.23-alpine

WORKDIR /app

# Copy the dependency files and download modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Rename .env.example to .env
RUN cp .env.example .env

# Build the Go application
RUN go build -o server main.go

# Expose the application port
EXPOSE 8080

# Set the command to run the built server
CMD ["./server"]
