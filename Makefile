.SILENT:

run:
	go run ./cmd/main.go

build:
	go mod download && go build -o ./bin/app ./cmd/main.go

