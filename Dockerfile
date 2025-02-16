# Используем базовый образ с Go
FROM golang:1.22.3 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы проекта
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Сборка бинарного файла
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/main.go

# Минимальный образ для запуска
FROM alpine:latest

WORKDIR /root/

# Устанавливаем зависимости
RUN apk --no-cache add ca-certificates

# Копируем бинарник из builder
COPY --from=builder /app/app .

# Запускаем приложение
CMD ["./app"]
