FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . . 

RUN CGO_ENABLED=0 GOOS=linux \
    go build -ldflags="-s -w" \
    -o toolhub-server \
    ./cmd/server


FROM alpine:3.22

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/toolhub-server .
COPY config/config.yaml ./config.yaml

EXPOSE 8080

USER nobody

ENTRYPOINT ["./toolhub-server"]
