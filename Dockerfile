# Build stage for Go application
FROM golang:1.20 AS builder

# Set working directory
WORKDIR /app

# Install required tools, including Git and Swag for Swagger documentation generation
RUN apt-get update && \
    apt-get install -y --no-install-recommends git && \
    rm -rf /var/lib/apt/lists/*

# Install Swag CLI tool with version pinning to generate Swagger docs
RUN go install github.com/swaggo/swag/cmd/swag@v1.8.12

# Copy dependency files first to leverage caching
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Generate Swagger docs
RUN swag init --dir ./cmd --output ./docs --parseDependency --parseInternal

# Build the Go application with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s -X main.version=1.0.0" \
    -o /app/main ./cmd/main.go

# Final runtime image
FROM scratch

# Set working directory for the runtime
WORKDIR /

# Import certificates and timezone data for a secure runtime environment
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Copy the built application and Swagger docs to the runtime image
COPY --from=builder --chown=1000:1000 /app/main . 
COPY --from=builder /app/docs ./docs/

# Expose the application's port
EXPOSE 8080

# Set environment variables (including time zone)
ENV TZ=UTC \
    GODEBUG=netdns=go

# Set non-root user for security
USER 1000:1000

# Define entrypoint for the Go application
ENTRYPOINT ["./main"]
