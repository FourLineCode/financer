dev:
	go build -o ./cmd/bin/ ./cmd/api/*.go && ./cmd/bin/main

build:
	go build -o ./cmd/bin/ ./cmd/api/*.go

run:
	go run ./cmd/api/*.go

start:
	./cmd/bin/main