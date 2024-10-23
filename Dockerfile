FROM golang:1.22.3

ENV GOPATH=/
WORKDIR /app

COPY ./ ./
RUN apt-get update && \ 
    go mod download && \ 
    go build -o app cmd/app/main.go

CMD ["./app"]