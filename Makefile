.SILENT:

run:
	go run ./cmd/main.go

build:
	go mod download && go mod tidy && go mod verify && go build ./cmd/main.go

