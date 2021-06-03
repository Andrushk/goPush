# Base environment (alias: base)
FROM amd64/golang:1.16.4 AS base

# Development environment
FROM base as dev
WORKDIR /root
RUN apt-get update && apt-get install -y fswatch
RUN go get github.com/go-delve/delve/cmd/dlv
WORKDIR /goPush
