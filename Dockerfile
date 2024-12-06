# Используем официальный образ Golang как базовый
FROM golang:1.23.3-alpine

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы go.mod и go.sum (если есть)
COPY go.mod ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код приложения
COPY . .

# Собираем приложение
RUN go build -o main .

# Указываем порт, который будет прослушивать приложение
EXPOSE 8080

# Запускаем приложение
CMD ["./main"]