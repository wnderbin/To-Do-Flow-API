FROM golang:1.22.2-alpine

# Установите необходимые пакеты для работы cgo и SQLite
RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app

COPY ["go.mod", "go.sum", "./"]

# Установите зависимости
RUN go mod download

COPY . .

# Устанавливаем переменную окружения
ENV CONFIG_PATH=./config/config.yaml

# Сборка приложения с включенным cgo
# Используем CGO_ENABLED=1 для сборки
CMD ["sh", "-c", "CGO_ENABLED=1 go run ./cmd/todo-flow/main.go"]