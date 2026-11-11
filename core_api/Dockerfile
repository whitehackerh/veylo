FROM golang:1.27-alpine AS development

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]

FROM golang:1.27-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /usr/local/bin/core-api ./src

FROM alpine:latest AS production

RUN apk add --no-cache ca-certificates

COPY --from=builder /usr/local/bin/core-api /usr/local/bin/core-api

EXPOSE 8080

ENTRYPOINT ["core-api"]
