FROM golang:1.20.5 AS build
WORKDIR /app
COPY .. /app
COPY go.mod go.sum /modules/

RUN go mod download
RUN go env -w CGO_ENABLED=0

RUN go build -o main ./cmd/main.go

FROM ubuntu:20.04
RUN apt-get update && apt-get install -y curl
COPY --from=build /app/configs/app/prod.toml /app/main /usr/local/bin/