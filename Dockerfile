FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

COPY .env .

RUN apk --no-cache add ca-certificates

EXPOSE 8080

CMD ["/app/main"]