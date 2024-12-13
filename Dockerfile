# syntax = docker/dockerfile:1-experimental
FROM golang:1.20-alpine AS build

# Install dependencies
RUN apk update && \
    apk upgrade && \
    apk add --no-cache bash git openssh make build-base

WORKDIR /build
COPY . /build/twitter-bee

RUN git clone https://github.com/xueqianLu/twitter-scraper /build/twitter-scraper
RUN cd /build/twitter-bee && make

FROM alpine

WORKDIR /root


COPY  --from=build /build/twitter-bee/build/bin/tbee /usr/bin/tbee
RUN chmod u+x /usr/bin/tbee

ENTRYPOINT [ "tbee" ]