# Use official Go image as base
FROM golang:1.21 AS builder

WORKDIR /app

# Copy go files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code and build
COPY . .
RUN go build -o calculator-service

# Use a minimal image for production
FROM debian:bullseye-slim

WORKDIR /app
COPY --from=builder /app/calculator-service .

EXPOSE 8080
CMD ["./calculator-service"]
