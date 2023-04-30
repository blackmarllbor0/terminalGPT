include .env # load vars in .env file

PROJECT_NAME = terminalGPT
PROJECT_PATH = cmd/$(PROJECT_NAME).go

CONTAINER_DB_NAME = chat_history

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
	make build && \
	sudo mv $(PROJECT_NAME) /usr/local/bin/${PROJECT_NAME}

.PHONY: up_db
up_db:
	docker-compose up -d

.PHONY: down_db
down_db:
	docker stop $(CONTAINER_DB_NAME) && \
	docker rm $(CONTAINER_DB_NAME)

.PHONY: mongosh
mongosh:
	docker exec -it $(CONTAINER_DB_NAME) mongosh \
	mongodb://172.22.0.2:27017/$(MONGO_INITDB_DATABASE) \
	-u $(MONGO_USER) \
	-p $(MONGO_PASSWORD)


.PHONY: get_db_host
get_db_host:
	docker inspect $(CONTAINER_DB_NAME) | grep IPAddress

.PHONY: restore_db
restore_db:
	make down_db && make up_db

.PHONY: container_logs
container_logs:
	docker logs ${CONTAINER_DB_NAME}