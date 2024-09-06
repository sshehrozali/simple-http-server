FROM golang:1.20-alpine as builder

WORKDIR /app

COPY . .

RUN go mod init go-docker-app && go build -o app

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]
