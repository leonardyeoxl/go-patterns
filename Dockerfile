FROM golang:1.15.6-alpine3.12
LABEL maintainer="Leonard Yeo"

WORKDIR /app/
ENV GO111MODULE=auto
RUN apk --no-cache add --virtual build-deps
RUN apk update && apk add --no-cache bash~=5.0.17-r0 git~=2.26.3-r0 && rm /var/cache/apk/*

# ENV PATH=$PATH:$GOPATH/bin:/opt/protoc/bin