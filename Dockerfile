# syntax=docker/dockerfile:1

FROM golang:1.20 AS build

WORKDIR $GOPATH/src/github.com/brotherlogic/mdb

COPY go.mod ./
COPY go.sum ./

RUN mkdir proto
COPY proto/*.go ./proto/

RUN mkdir server
COPY server/*.go ./server/

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 go build -o /mdb

##
## Deploy
##
FROM ubuntu:22.10

WORKDIR /
COPY --from=build /mdb /mdb

RUN apt update && apt install -y nmap

EXPOSE 8080
EXPOSE 8081

USER root:root

ENTRYPOINT ["/mdb"]