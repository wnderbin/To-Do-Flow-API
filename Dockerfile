FROM golang:1.22.2-alpine


RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app

COPY ["go.mod", "go.sum", "./"]

RUN go mod download

COPY . .

ENV CONFIG_PATH=./config/config.yaml

CMD ["sh", "-c", "CGO_ENABLED=1 go run ./cmd/todo-flow/main.go"]