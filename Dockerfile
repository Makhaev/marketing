# Stage 1: build
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o price-app ./cmd/app

# Stage 2: run
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/price-app .
COPY .env .

EXPOSE 8080

CMD ["./price-app"]