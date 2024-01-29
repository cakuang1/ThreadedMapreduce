# Use the official Go image as the base image
FROM golang:latest


WORKDIR /app

COPY . .

RUN go build -o MultiThreadedMR



CMD ["./MultiThreadedMR"]
