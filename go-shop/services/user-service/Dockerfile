FROM golang:1.24-alpine AS builder
WORKDIR /app
# --- 关键：添加下面这一行 ---
ENV GOPROXY=https://goproxy.cn,direct
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o user-service ./cmd

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/user-service .

EXPOSE 9000
CMD ["./user-service"]
