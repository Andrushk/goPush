# Base environment (alias: base)
FROM amd64/golang:1.16.4 AS base

# Development environment
FROM base as dev
WORKDIR /root
RUN apt-get update && apt-get install -y fswatch
RUN go get github.com/go-delve/delve/cmd/dlv
WORKDIR /gopush


FROM base as builder
COPY . /gopush
WORKDIR /gopush
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o gopush

#Production environment
FROM scratch as prod
COPY --from=builder /gopush/config/config.yml /gopush/config/config.yml
COPY --from=builder /gopush/gopush /gopush/gopush
WORKDIR /gopush