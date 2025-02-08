# Step 1: Build the app using a Go image
FROM golang:1.20-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy Go modules files to install dependencies
COPY src/go.mod ./
RUN go mod download

# Copy the entire source code
COPY ./src /app

# Build the Go app inside the container
RUN go build -o main .

# Step 2: Create a smaller final image to run the app
FROM alpine:latest

WORKDIR /root/

# Copy the built binary from the builder image
COPY --from=builder /app/main .

# Expose the necessary port (8080)
EXPOSE 8080

# Run the Go application
CMD ["./main"]
