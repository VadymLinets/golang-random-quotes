FROM golang:1.22 as builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o quotes .

FROM scratch

WORKDIR /app
COPY --from=builder /app/migrations /app/migrations
COPY --from=builder /app/quotes /app/quotes
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["./quotes"]
