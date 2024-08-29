FROM golang:1.22.3-alpine3.20 AS builder

RUN mkdir app

COPY . /app

WORKDIR /app

RUN go build -o api-gateway ./cmd/main.go

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app .

CMD ["/app/api-gateway"]