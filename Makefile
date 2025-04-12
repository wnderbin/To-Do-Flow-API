go-compile: go-build go-build-run

go-run:
	CONFIG_PATH=./config/config.yaml go run ./cmd/todo-flow/main.go

go-build: 
	go build ./cmd/todo-flow/main.go

go-build-run:
	CONFIG_PATH=./config/config.yaml ./main

go-run-workflow:
	CONFIG_PATH=./config/config.yaml go run ./cmd/todo-flow/main.go 1

go-build-run-workflow: go-build
	CONFIG_PATH=./config/config.yaml ./main 1