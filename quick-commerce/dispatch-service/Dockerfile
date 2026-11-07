# Stage 1: Build the Go binary
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go mod init dispatch-service && \
    CGO_ENABLED=0 GOOS=linux go build -o dispatch-service main.go

# Stage 2: Run in a minimal Alpine container
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/dispatch-service .

EXPOSE 8003
CMD ["./dispatch-service"]