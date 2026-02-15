# ---------- Build Stage ----------
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN  go build -o app .

# ---------- Runtime Stage ----------
FROM golang:1.22-alpine
WORKDIR /app
COPY --from=builder /app/app .

ENTRYPOINT ["/app/app"]
