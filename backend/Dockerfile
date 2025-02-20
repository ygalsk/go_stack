FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN go build -o main

# Final container
FROM debian:bookworm
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
