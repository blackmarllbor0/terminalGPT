PROJECT_NAME = terminalGPT
PROJECT_PATH = cmd/$(PROJECT_NAME).go

.PHONY:run
run:
	go run $(PROJECT_PATH)

.PHONY:build
build:
	go build -o ./$(PROGRAM_NAME) $(PROJECT_PATH)

.PHONY:test
test:
	go test ./...

.PHONY:lint
lint:
	golangci-lint run

.PHONY: exec
exec:
	make build && sudo mv $(PROJECT_NAME) /usr/local/bin/${PROJECT_NAME}