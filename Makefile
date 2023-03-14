.SILENT:

run:
	go run ./cmd/main.go

build:
	go mod download && go mod tidy && go mod verify && go build -o mementor ./cmd/main.go

swagger:
	swag init -g cmd/main.go

docker:
		docker build -t mementor .

test:
	go test ./...