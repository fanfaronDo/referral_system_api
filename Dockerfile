FROM golang:1.22.3

WORKDIR=/app
ENV GOPATH=/app

COPY ./ ./
RUN apt-get update

RUN go mod download
RUN go build -o app ./cmd/app/main
