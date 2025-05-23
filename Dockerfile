# Use the official Golang image as a base image
FROM golang:1.24.3-alpine3.21

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker layer caching
COPY go.* ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port the application runs on
EXPOSE 6380

# Command to run the executable
CMD ["./main"]
