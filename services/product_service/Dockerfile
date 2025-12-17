# Dockerfile
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Dependencies
COPY go.mod go.sum ./
RUN go mod download

# Source
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o product-service .

# -------------------------
# Final stage
# -------------------------
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/product-service .

EXPOSE 8082

CMD ["./product-service"]

