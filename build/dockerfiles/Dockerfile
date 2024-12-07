FROM golang:1.23.2-bullseye AS builder

# Move to working directory (/build).
WORKDIR /build

# Copy dependency files first to leverage Docker caching.
COPY go.mod go.sum ./

# Download dependencies early to cache them if unchanged.
RUN go mod download

# Copy the rest of the application code.
COPY . .

# Set environment variables for Go build.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Build the Go application.
RUN go build -ldflags="-s -w" -o server .

# Use a minimal base image for production.
FROM alpine:latest

# Set working directory.
WORKDIR /app/

# Copy the binary from the builder stage.
COPY --from=builder /build/server /app/bin/server

# Expose the application port.
EXPOSE 8080
EXPOSE 50051

# Command to run the application.
CMD ["./bin/server", "serve"]
