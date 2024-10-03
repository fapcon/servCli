FROM golang:1.22-alpine as builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main ./cmd \
    && go build -o migrate ./migrations

FROM alpine:latest

COPY --from=builder /app /app

EXPOSE 8080

CMD ["/main"]