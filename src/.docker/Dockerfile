
FROM golang:1.12.5-alpine3.9

WORKDIR /go/src

RUN apk update && \
    apk add --update git && \
    apk add openssh && \
    rm -rf /var/cache/apk/*

ENV GO111MODULE=on 
ENV CGO_ENABLED=0
ENV TZ = Asia/Tokyo
