FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o quotes .

FROM scratch

WORKDIR /app
COPY --from=builder /app/migrations /app/migrations
COPY --from=builder /app/quotes /app/quotes
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["./quotes"]
