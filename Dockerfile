FROM golang:1.15.2-alpine3.12

COPY scripts/wait-for /usr/local/bin

RUN chmod +x /usr/local/bin/wait-for

RUN mkdir /app

RUN apk add gcc

WORKDIR /app

