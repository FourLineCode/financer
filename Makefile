dev:
	go build -o ./cmd/ src/*.go && ./cmd/main.exe

build:
	go build -o ./cmd/ src/*.go

run:
	go run src/*.go

start:
	./cmd/main.exe