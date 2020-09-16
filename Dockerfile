FROM golang:1.15-alpine3.12 as build

WORKDIR /usr/local/src

COPY ["./hello-world", "."]
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o hello-world

FROM alpine:latest

ARG port=$HTTP_PORT

WORKDIR /usr/bin/
COPY --from=build ["/usr/local/src/hello-world", "."]

EXPOSE $port

ENTRYPOINT ["./hello-world"]
