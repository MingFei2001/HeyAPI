# Use official go image as base image
FROM golang:alpine AS builder

# Set the working directory if not already set
WORKDIR /app

# Copy the go mod and download dependencies
COPY go.mod ./
RUN go mod download

# Copy necessary assets and source code
COPY main.go ./
COPY handlers ./handlers
COPY static ./static
COPY templates ./templates

# Build the binary
RUN go build -o heyapi .

################################################

# Use alpine as final container image
FROM alpine:latest

# Set the working directory if not already set
WORKDIR /app

# Copy the binary and assets from builder stage
COPY --from=builder /app/heyapi .
COPY --from=builder /app/static ./static
COPY --from=builder /app/templates ./templates

# Expose port 8080
EXPOSE 8080

# Run the binary
CMD ["./heyapi"]
