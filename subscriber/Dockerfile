FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /nats-subscribe

ENV PORT 8080

EXPOSE 8080

CMD ["/nats-subscribe"]