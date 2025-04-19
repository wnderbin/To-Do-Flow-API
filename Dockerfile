FROM golang:1.22.2-alpine

WORKDIR /app

COPY ["go.mod", "go.sum", "./"]

RUN go mod download

COPY . .

ENV CONFIG_PATH=./config/config.yaml

CMD ["go", "run", "./cmd/todo-flow/main.go"]