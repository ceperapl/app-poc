# build the server binary
FROM golang:1.16-alpine AS base-builder
WORKDIR /app-poc
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/server -tags release ./cmd/server

# copy the server binary from builder stage; run the server binary
FROM alpine:latest
WORKDIR /bin
COPY --from=base-builder /app-poc/bin/server .
# COPY --from=base-builder /app-poc/db/migrations /db/migrations/
ENTRYPOINT ["server"]
