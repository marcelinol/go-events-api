FROM golang:1.9-alpine

RUN apk add --no-cache git mercurial \
    && go get github.com/marcelinol/go-events-api \
    && apk del git mercurial

ENTRYPOINT cd /go/src/github.com/marcelinol/go-events-api \
    && go run main.go
