# Use official Golang image as builder
FROM golang:1.24.3 AS builder

# Force static binary
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

# Copy go.mod and go.sum first to leverage caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the binary
RUN go build -o server ./cmd/server/main.go

# ---- Final Image ----
FROM debian:bullseye-slim

# Add a non-root user for security
RUN useradd -m appuser

# Copy the compiled binary from the builder
COPY --from=builder /app/server /server

# Set permissions
RUN chown appuser:appuser /server

USER appuser

EXPOSE 3600

ENTRYPOINT ["/server"]
