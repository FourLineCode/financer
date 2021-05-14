dev:
	go run ./cmd/api/*.go

build:
	go build -o ./cmd/bin/main ./cmd/api/*.go

run:
	go run ./cmd/api/*.go

start:
	./cmd/bin/main