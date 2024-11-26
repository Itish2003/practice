# Use Go base image
FROM golang:1.22.5

# Set working directory
WORKDIR /build

# Copy all files into the container
COPY . .

# Download dependencies
RUN go mod download

# Build the application
RUN go build -o main .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./main"]
