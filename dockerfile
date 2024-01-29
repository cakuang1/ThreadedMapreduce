# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the go.mod and go.sum files to download dependencies
COPY go.mod .

# Download dependencies
RUN go mod download

# Copy the entire project to the container
COPY . .

# Build the Go application
RUN go build -o app

# Run the Go application when the container starts
CMD ["./app"]
