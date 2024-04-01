FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go mod init browser-service

RUN go get github.com/gorilla/websocket@latest
RUN go get github.com/redis/go-redis/v9@latest

RUN go mod tidy

RUN go build -o main .

CMD ["./main"]