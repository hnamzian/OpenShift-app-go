# Use the official Golang image as the build environment
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the application code into the container
COPY . .

# Build the application
RUN go build -o my-go-app .

# Use a minimal image for the final container
FROM alpine:latest

# Copy the compiled binary from the builder stage
COPY --from=builder /app/my-go-app /my-go-app

# Expose port 8080
EXPOSE 8080

# Set the entry point to run the application
CMD ["/my-go-app"]

