# Build Stage
FROM golang:1.22-alpine3.19 AS builder

WORKDIR /app

COPY . .

RUN go build -o main ./cmd/api

# Final Stage
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

COPY wait-for.sh .

RUN chmod +x wait-for.sh

EXPOSE 8080
CMD ["/app/wait-for.sh", "db:5432", "--", "/app/main"]