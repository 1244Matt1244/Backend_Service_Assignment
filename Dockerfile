# Build stage
FROM golang:1.20-bullseye AS builder
WORKDIR /app

# Install tools (combined layers for efficiency)
RUN apt-get update && \
    apt-get install -y --no-install-recommends git && \
    rm -rf /var/lib/apt/lists/*

# Install swag with version pinning
RUN go install github.com/swaggo/swag/cmd/swag@v1.8.12

# Copy dependency files first
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Generate Swagger docs (ensure source is copied first)
RUN swag init --dir ./cmd --output ./docs --parseDependency --parseInternal

# Build application with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s -X main.version=1.0.0" \
    -o /app/main ./cmd/main.go

# Final stage (using scratch for minimal size)
FROM scratch
WORKDIR /

# Import certificates and timezone data
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Copy application and docs
COPY --from=builder --chown=1000:1000 /app/main .
COPY --from=builder /app/docs ./docs/

# Metadata
EXPOSE 8080
ENV TZ=UTC \
    GODEBUG=netdns=go
USER 1000:1000

# Entrypoint
ENTRYPOINT ["./main"]
