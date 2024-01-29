# Use an official Golang runtime as a parent image
FROM golang:latest

WORKDIR /go/src/app

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o app

CMD ["./app"]
