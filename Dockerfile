FROM golang:1.23.0-alpine AS builder
LABEL authors="osami"

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main cmd/server/main.go

FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]