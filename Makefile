.PHONY: run dev build docker-up docker-down migrate-up migrate-down test lint

run:
	go run cmd/server/main.go

dev:
	air # Use air for hot reloading if installed

build:
	go build -o bin/server cmd/server/main.go

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

test:
	go test -v ./...

lint:
	golangci-lint run
