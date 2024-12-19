# Use a lightweight Go image as the base
FROM golang:1.23.4-alpine

# Set the working directory
WORKDIR /medoc-kms

# Copy the project source code
COPY . .

# Download dependencies
RUN go mod download
# Build the Go application
RUN go build -o build/medoc-kms

# Expose the port your application listens on (adjust if needed)
EXPOSE 8080

# Command to run the application when the container starts
CMD ["./build/medoc-kms"]