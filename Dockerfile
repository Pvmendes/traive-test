# Use the official Go image as a parent image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download and install Go module dependencies
RUN go mod download

# Copy the entire project directory to the container
COPY . .

# Build the Go application
RUN go build -o my-transactions-api cmd/main.go

# Expose the port the application will run on
EXPOSE 8080

# Define the command to run the application
CMD ["./my-transactions-api"]
