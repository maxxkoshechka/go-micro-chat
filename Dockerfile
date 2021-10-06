FROM golang:1.17.1-bullseye as builder

RUN apt-get update && apt-get install -y ca-certificates openssl tzdata
RUN update-ca-certificates

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

ADD . /app
WORKDIR /app/src
RUN go build -o main .

FROM ubuntu:latest as runner

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

RUN mkdir /app
WORKDIR /app/src
COPY --from=builder /app/src/main main

CMD ["./main"]
