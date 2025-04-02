# Stage 1: Build the Go app
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main ./cmd/main.go

EXPOSE 8990
CMD ["./main"]
