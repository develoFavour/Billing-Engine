# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server/main.go

# Final stage
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .
# Copy migrations so the app can run them if needed (or run manually)
COPY --from=builder /app/migrations ./migrations

# Expose the port
EXPOSE 8080

# Command to run the application
CMD ["./main"]
