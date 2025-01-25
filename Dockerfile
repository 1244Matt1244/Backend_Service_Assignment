# Stage 1: Go Application Build Stage
FROM golang:1.20 AS go-builder

# Set working directory
WORKDIR /app

# Install required tools, including Git and Swag for Swagger documentation generation
RUN apt-get update && \
    apt-get install -y --no-install-recommends git && \
    rm -rf /var/lib/apt/lists/*

# Install Swag CLI tool with version pinning to generate Swagger docs
RUN go install github.com/swaggo/swag/cmd/swag@v1.8.12

# Copy Go dependency files first to leverage caching
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

# Stage 2: Java Application Build Stage
FROM maven:3.8.6-eclipse-temurin-17 AS java-builder

# Set working directory for the Java application
WORKDIR /app

# Copy Maven dependency files and install dependencies offline
COPY pom.xml .
RUN mvn dependency:go-offline

# Copy the Java source code
COPY src ./src
RUN mvn package -DskipTests

# Stage 3: Final Runtime Image for Go and Java
FROM eclipse-temurin:17-jre-jammy AS runtime

# Set working directory for the runtime
WORKDIR /app

# Copy the Go application and Swagger docs from the Go build stage
COPY --from=go-builder --chown=1000:1000 /app/main .
COPY --from=go-builder /app/docs ./docs/

# Copy the Java application JAR and necessary files from the Java build stage
COPY --from=java-builder /app/target/*.jar app.jar
COPY --from=java-builder /app/src/main/resources/opentelemetry-javaagent.jar .

# Expose ports for both Go and Java applications
EXPOSE 8080

# Set environment variables (including time zone)
ENV TZ=UTC \
    GODEBUG=netdns=go

# Set non-root user for security
USER 1000:1000

# Define entrypoint to run both applications as needed
ENTRYPOINT ["sh", "-c", "java -javaagent:opentelemetry-javaagent.jar -jar app.jar & ./main"]
