# Используем официальный образ Go 1.23
FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o server ./cmd/main.go

CMD ["./server"]
