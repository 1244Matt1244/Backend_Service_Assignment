# Build stage
FROM golang:1.20 AS builder
WORKDIR /app

# Install necessary tools
RUN apt-get update && apt-get install -y git

# Install swag for generating Swagger documentation
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy go.mod and go.sum first to leverage Docker cache and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project source code
COPY . .

# Generate Swagger documentation
RUN swag init --dir ./cmd --output ./docs

# Build the Go app targeting the correct Go file in the cmd directory
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

# Final stage: Use a minimal image to run the application
FROM alpine:3.16
WORKDIR /root/

# Add CA certificates to enable SSL connections from the application
RUN apk --no-cache add ca-certificates

# Ensure the docs directory exists before copying files into it
RUN mkdir -p ./docs

# Copy the compiled binary and Swagger documentation from the build stage
COPY --from=builder /app/main . 
COPY --from=builder /app/docs ./docs/

# Expose port 8080 for the application
EXPOSE 8080

# Run the compiled Go binary
CMD ["./main"]
