# Stage 1: Build the Go binary
FROM golang:1.21 AS builder

WORKDIR /app

# Copy go files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code and build
COPY . .
RUN go build -o calculator-service

# Stage 2: Use a modern base image with GLIBC ≥ 2.34
FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/calculator-service .

EXPOSE 8080
CMD ["./calculator-service"]
