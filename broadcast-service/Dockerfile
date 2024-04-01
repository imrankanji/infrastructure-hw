FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod init broadcast-service
RUN go get github.com/redis/go-redis/v9@latest

RUN go mod tidy

RUN go build -o main .

FROM scratch

WORKDIR /app

COPY --from=builder /app/main /app/main

CMD ["./main"]
