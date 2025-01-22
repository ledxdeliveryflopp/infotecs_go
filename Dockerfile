FROM golang:1.23.3-alpine3.20 as builder

LABEL meintainer="LedxDeliveryFlopp"

WORKDIR /build

COPY . .

RUN go mod tidy && go mod download && go build -o /api main.go

FROM alpine:3.20.3

WORKDIR /app

COPY --from=builder api/ /app/api

ENTRYPOINT ["/app/api"]