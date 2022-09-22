FROM golang:1.19.1

ENV DISCORD_TOKEN=$DISCORD_TOKEN

WORKDIR /app
COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download

COPY ./main.go ./

RUN go build -o ./dadbot

EXPOSE 8080

ENTRYPOINT ["/dadbot"]
