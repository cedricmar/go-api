FROM golang:alpine

ENV GOPATH /go 
ENV PATH $PATH:/go/bin

RUN apk update && \
    apk upgrade && \
    apk add git

RUN go get github.com/pilu/fresh

ADD . /go/src/apps/api
RUN cd /go/src/apps/api && \
    go get ./...
