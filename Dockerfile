FROM golang:1.21-alpine AS builder

WORKDIR /app

# Копируем только модули сначала (для кэширования)
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальной код
COPY . .

# Собираем бинарник под Linux amd64
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

# Финальный контейнер — минимальный
FROM alpine:latest

WORKDIR /app

# Копируем только бинарник из билдера
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
COPY --from=builder /app/migrations ./migrations

EXPOSE 8080

CMD ["./main"]