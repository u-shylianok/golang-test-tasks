.PHONY: build start run clean

build:
	go build -o bin/task04 cmd/*.go

start:
	make build
	./bin/task04

run:
	go run cmd/*.go

clean:
	rm bin/task04