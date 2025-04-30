FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./main.go

FROM ubuntu:22.04
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/config/migrations ./config/migrations
EXPOSE 8080
CMD ["./main"]
