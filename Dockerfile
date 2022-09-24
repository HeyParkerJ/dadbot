# syntax=docker/dockerfile:1

FROM golang:1.19.1-alpine

WORKDIR /app
COPY go.mod ./
COPY go.sum ./

ARG DISCORD_TOKEN
ENV DISCORD_TOKEN=${ DISCORD_TOKEN }

RUN go mod download

COPY *.go ./

RUN go build -o ./bin/dadbot

EXPOSE 8080

ENTRYPOINT "./bin/dadbot" "-t" "$DISCORD_TOKEN"
# CMD ["-t", $DISCORD_TOKEN]
