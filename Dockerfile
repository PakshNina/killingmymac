# Build stage
ARG GO_VERSION=1.24.1
FROM golang:${GO_VERSION} AS builder

WORKDIR /app

# Copy source code and dependencies
COPY go.mod ./
COPY withstat ./withstat

# Build the application (statically linked)
RUN CGO_ENABLED=0 GOOS=linux go build \
    -a -installsuffix cgo \
    -o app \
    ./withstat

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]

