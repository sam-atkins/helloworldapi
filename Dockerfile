FROM golang:1.15.5-alpine

WORKDIR /app
COPY . .

RUN go mod download

RUN go build -o main .

CMD ["/app/main"]
