# Stage 1: Build the Go binary
FROM golang:1.21 AS builder

WORKDIR /app
COPY . .

RUN go build -o server cmd/web/main.go

# Stage 2: Create the final image
FROM ubuntu:22.04

WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/config.json .  # Ensure the config file is copied


EXPOSE 8080

CMD ["./server"]
