FROM golang:1.19.1

WORKDIR /app
COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download

COPY ./main.go ./

RUN go build -o ./dadbot

EXPOSE 8080

ENTRYPOINT ["/dadbot", "-t", $DISCORD_TOKEN]
