# Start with the official Go image
FROM golang:1.20 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the Go source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o app

# Start a new, smaller image
FROM alpine:latest

# Install Ghostscript
RUN apk add --update --no-cache ghostscript

# Set the working directory inside the container
WORKDIR /app

# Copy the built Go binary from the previous stage
COPY --from=builder /app/app .

# Copy the 'script' directory and 'shrinkpdf.sh' file into the container
COPY script ./script

# Expose the port the Go server will listen on
EXPOSE 8080

# Start the Go server
CMD ["./app"]
