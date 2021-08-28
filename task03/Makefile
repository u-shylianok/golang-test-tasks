.PHONY: build start run clean

build:
	go build -o bin/task03 cmd/*.go

start:
	make build
	./bin/task03

run:
	go run cmd/*.go

clean:
	rm bin/task03