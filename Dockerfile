# syntax=docker/dockerfile:1

FROM golang:1.23 AS build

WORKDIR $GOPATH/src/github.com/brotherlogic/mdb

COPY go.mod ./
COPY go.sum ./

RUN mkdir proto
COPY proto/*.go ./proto/

RUN mkdir server
RUN mkdir lookup
COPY server/*.go ./server/
COPY lookup/*.go ./lookup

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 go build -o /mdb

##
## Deploy
##
FROM ubuntu:22.04
USER root:root

WORKDIR /
COPY --from=build /mdb /mdb

RUN apt update && apt upgrade -y && apt install -y nmap net-tools

EXPOSE 8080
EXPOSE 8081


ENTRYPOINT ["/mdb"]