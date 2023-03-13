FROM golang:bullseye as builder

WORKDIR /mementor
COPY . .

RUN CGO_ENABLED=0  GOOS=linux  GOARCH=amd64 go build -v -o ./application ./cmd/*.go

FROM alpine:3.15.4
WORKDIR /mementor

COPY --from=builder /mementor/application /mementor/application
COPY ./config/config.yml ./config/config.yml
CMD ["/mementor/application"]

