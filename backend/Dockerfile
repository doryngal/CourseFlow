# Используем базовый образ Golang
FROM golang:1.18-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем все файлы в контейнер
COPY docker .

# Устанавливаем зависимости
RUN go mod download

# Собираем приложение
RUN go build -o main .

# Открываем порт 8080 для приложения
EXPOSE 8080

# Команда для запуска
CMD ["./main"]