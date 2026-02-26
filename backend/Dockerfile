# Build stage
FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o server ./cmd/server

# Run stage
FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/server /app/server
EXPOSE 18601
CMD ["/app/server"]
