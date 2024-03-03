build:
	go build -o bin/main server/main.go

run:
	go run server/main.go

all: build run