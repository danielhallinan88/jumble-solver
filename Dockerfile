# Build Stage
FROM golang:1.23 AS builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o jumble-solver .

# Final Runtime Stage
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/jumble-solver .
COPY --from=builder /app/words_alpha.txt .

EXPOSE 8082
CMD ["./jumble-solver"] 
