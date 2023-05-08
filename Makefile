include .env # load vars in .env file

PROJECT_NAME = terminalGPT
PROJECT_PATH = cmd/$(PROJECT_NAME).go

CONTAINER_DB_NAME 		 = chat_history
CONTAINER_DB_SERVICE 	 = mongodb
CONTAINER_WEB_WB_NAME 	 = mongo-web
CONTAINER_WEB_WB_SERVICE = webDB

MONGO_AUTH_STRING = mongodb://172.22.0.2:27017/$(MONGO_INITDB_DATABASE) -u $(MONGO_USER) -p $(MONGO_PASSWORD)

.PHONY:run 			## used to run in dev mod.
.PHONY:build 		## used to build an app into a bin file.
.PHONY:test 		## used to run all tests in the app.
.PHONY:lint 		## used to run the linter.
.PHONY: exec 		## places the bin file in /usr/local/bin to make it executable.
.PHONY: up_db 		## lifts the docker container from the database.
.PHONY: down_db 	## delete container from database.
.PHONY: mongosh 	## launched a database management utility from the terminal.
.PHONY: get_db_host	## returns a string with the host from the database.
.PHONY: restore_db	## recreates the container with the database.

.DEFAULT_GOAL:run
run:
	go run $(PROJECT_PATH)

build:
	go build -o ./$(PROGRAM_NAME) $(PROJECT_PATH)

test:
	go test ./...

lint:
	golangci-lint run

exec: build
	sudo mv $(PROJECT_NAME) /usr/local/bin/${PROJECT_NAME}

up_db:
	docker-compose up ${CONTAINER_DB_SERVICE} -d

down_db:
	docker stop $(CONTAINER_DB_NAME) && docker rm $(CONTAINER_DB_NAME)

restore_db: down_db up_db

mongosh:
	docker exec -it $(CONTAINER_DB_NAME) mongosh ${MONGO_AUTH_STRING}

get_db_host:
	docker inspect $(CONTAINER_DB_NAME) | grep IPAddress