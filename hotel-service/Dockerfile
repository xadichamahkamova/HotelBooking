FROM golang:1.22.3-alpine3.20 AS builder

RUN mkdir app

COPY . /app

WORKDIR /app

RUN go build -o hotel_service ./cmd/main.go

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app .

CMD ["/app/hotel_service"]