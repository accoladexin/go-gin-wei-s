FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o order-service ./cmd

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/order-service .

EXPOSE 8002
EXPOSE 9002
CMD ["./order-service"]

