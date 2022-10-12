.SILENT:

run:
	go run ./cmd/main.go

build:
	go mod download && go build ./cmd/main.go

