# Stage 1: Build the Go binary
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go mod init inventory-service && \
    CGO_ENABLED=0 GOOS=linux go build -o inventory-service main.go

# Stage 2: Run in a minimal Alpine container
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/inventory-service .

EXPOSE 8002
CMD ["./inventory-service"]