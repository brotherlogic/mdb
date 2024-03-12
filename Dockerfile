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

RUN apt update && apt install -y nmap

##
## Deploy
##
FROM busybox:1.35.0-uclibc as busybox
FROM gcr.io/distroless/base-debian11

WORKDIR /

# Now copy the static shell into base image.
COPY --from=busybox /bin/sh /bin/sh
COPY --from=busybox /usr/bin/apt /usr/bin/apt

COPY --from=build /mdb /mdb

EXPOSE 8080
EXPOSE 8081

USER root:root

ENTRYPOINT ["/mdb"]