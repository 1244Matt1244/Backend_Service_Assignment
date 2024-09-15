# Step 1: Build the Go app using the official Go image
FROM golang:1.20 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy only go.mod and go.sum to leverage Docker cache and download dependencies
COPY go.mod go.sum ./

# Download dependencies early for better caching
RUN go mod download

# Copy the rest of the project files
COPY . .

# Build the Go app with disabled CGO for a static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Step 2: Prepare the runtime environment using a minimal Alpine image
FROM alpine:latest  

# Install CA certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Set the working directory for the final container
WORKDIR /root/

# Copy the statically built Go binary from the builder stage
COPY --from=builder /app/main .

# Expose the port the app will run on (optional, but good practice)
EXPOSE 8080

# Run the Go binary
CMD ["./main"]
