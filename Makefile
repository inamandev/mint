BINARY=mint
MODULE=github.com/yourusername/mint

.PHONY: build run test lint tidy install clean

build:
	go build -o bin/$(BINARY) .

run:
	go run . $(ARGS)

install:
	go install .

test:
	go test ./...

lint:
	golangci-lint run ./...

tidy:
	go mod tidy

clean:
	rm -rf bin/

test-create:
	go run . create testapp
