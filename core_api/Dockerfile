FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /usr/local/bin/core-api ./src

FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=builder /usr/local/bin/core-api /usr/local/bin/core-api

EXPOSE 8080

ENTRYPOINT ["core-api"]
