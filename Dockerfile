FROM golang:1.20.5 AS build
WORKDIR /app
COPY . /app
COPY go.mod go.sum /modules/

RUN go mod download
RUN go env -w CGO_ENABLED=0
RUN go env -w GOOS=linux
RUN go env -w GOARCH=amd64

RUN go build -o main ./cmd/main.go

FROM ubuntu:20.04
COPY --from=build /app/config/prod.toml /app/main /usr/local/bin/