go-compile: go-build go-build-run

go-build: 
	go build ./cmd/todo-flow/main.go

go-build-run:
	CONFIG_PATH=./config/config.yaml ./main
