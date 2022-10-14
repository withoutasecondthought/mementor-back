FROM golang:bullseye

RUN mkdir "/app"

ADD . /app/

WORKDIR /app

RUN make build

CMD ["./main"]



