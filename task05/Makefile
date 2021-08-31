.PHONY: build start up down clean cleanall

build:
	GOOS=linux GOARCH=amd64 go build -o build/app/task05 cmd/*.go
	docker-compose build

start: build
	docker-compose up

up:
	docker-compose up

down:
	docker-compose down

clean:
	docker rm u-shylianok-task05-app u-shylianok-task05-pg
	rm build/app/task05

cleanall: clean
	docker rmi u-shylianok-task05-app u-shylianok-task05-pg