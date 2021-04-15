dev:
	go build -o ./cmd/bin/ ./cmd/api/*.go && ./cmd/bin/main.exe

build:
	go build -o ./cmd/bin/ ./cmd/api/*.go

run:
	go run ./cmd/api/*.go

start:
	./cmd/bin/main.exe