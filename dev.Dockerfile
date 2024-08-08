FROM golang:1.22.6-alpine as builder
RUN apk add make 

WORKDIR /app
COPY go.mod Makefile ./

RUN go install github.com/air-verse/air@latest

COPY . .

CMD ["air", "-c", ".air.toml"]

