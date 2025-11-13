# --- Stage 1: Build ---
FROM golang:1.25-bookworm AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /app/server -ldflags="-w -s" ./cmd/api

# --- Stage 2: Final Image ---
FROM gcr.io/distroless/static-debian11 AS final
# Copy configuration files (if you have them)
# COPY --from=builder /app/configs ./configs
COPY --from=builder /app/server /server
EXPOSE 8080
ENTRYPOINT ["/server"]