FROM golang:1.22.3

ENV GOPATH=/

COPY ./ ./
RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

RUN go mod download && go build -o app cmd/app/main.go

CMD ["./app"]