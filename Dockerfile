# Stage 1: Build the Go binary
FROM golang:1.23 AS builder

# Set working directory
WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary (static for smaller final image)
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/

# Stage 2: Run the Go app
FROM alpine:latest

WORKDIR /root/

# Install CA certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Copy compiled binary
COPY --from=builder /app/main .

# Expose app port (change if different)
EXPOSE 8080

# Run the binary
CMD ["./main"]