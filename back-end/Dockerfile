# Use official Golang image
FROM golang:1.20-alpine

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy the rest of the application
COPY . .

# Expose port 8080
EXPOSE 8080

# Run the Go app in development mode
CMD ["go", "run", "./cmd/main.go", "http-server"]
