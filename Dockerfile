FROM golang:1.15.5-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app .
CMD ["/app/main"]
