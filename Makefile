dev:
	go run ./cmd/api/*.go

build:
	go build -o ./bin/main ./cmd/api/*.go

run:
	go run ./cmd/api/*.go