# Используем официальный образ Go 1.23
FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o client ./cmd/main.go

RUN go test ./cmd/main.go -cover
