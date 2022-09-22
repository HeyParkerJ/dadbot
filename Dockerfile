FROM golang:1.19.1

ENV TOKEN ""

WORKDIR /app
COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download

COPY ./main.go ./

RUN go build -o ./dadbot

ENTRYPOINT ["/dadbot"]
