# Build stage
FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY .env .

RUN CGO_ENABLED=0 GOOS=linux go build -o backend

# Run stage
FROM alpine:latest
RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=builder /app/backend .

EXPOSE 8080

CMD ["/app/backend"]
