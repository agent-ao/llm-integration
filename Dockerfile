# syntax=docker/dockerfile:1.4
FROM golang:1.24-alpine AS base

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# === Build Stage (for production) ===
FROM base AS build
COPY . .
RUN go build -o /app/bin/server ./cmd/server

# === Dev Stage with Air ===
FROM base AS dev
RUN go install github.com/air-verse/air@latest
COPY . .
COPY .air.toml .air.toml
CMD ["air"]

# === Final Production Stage ===
FROM alpine:3.19 AS prod
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=build /app/bin/server .
CMD ["./server"]
