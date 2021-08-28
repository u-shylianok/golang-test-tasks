.PHONY: build start run clean

build:
	go build -o bin/task02 cmd/*.go

start:
	make build
	./bin/task02

run:
	go run cmd/*.go

clean:
	rm bin/task02