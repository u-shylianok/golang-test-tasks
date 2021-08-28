.PHONY: build start run clean

build:
	go build -o bin/task01 cmd/*.go

start:
	make build
	./bin/task01

run:
	go run cmd/*.go

clean:
	rm bin/task01