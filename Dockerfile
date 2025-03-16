# Use the official Golang image as the builder
FROM golang:1.23

# Set the working directory
WORKDIR /app

ENV GOEXPERIMENT=arenas

# Copy go.mod and go.sum to leverage caching
COPY go.mod go.sum ./

# Download dependencies first (leveraging caching)
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["./main"]